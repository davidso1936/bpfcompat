#!/usr/bin/env bash
set -euo pipefail

# Install the fetch-only Production Runtime Agent Alpha systemd units.
# This script does not configure secrets and does not enable host loading.

PREFIX="${PREFIX:-/usr/local}"
BIN_SRC="${BIN_SRC:-./bin/bpfcompat}"
VALIDATOR_SRC="${VALIDATOR_SRC:-validator/c-libbpf/bin/bpfcompat-validator}"
VALIDATOR_DEST="${VALIDATOR_DEST:-$PREFIX/libexec/bpfcompat/bpfcompat-validator}"
AGENT_USER="${AGENT_USER:-bpfcompat-agent}"
AGENT_GROUP="${AGENT_GROUP:-bpfcompat-agent}"
ENV_PATH="${ENV_PATH:-/etc/bpfcompat/agent.env}"
LOAD_ENV_PATH="${LOAD_ENV_PATH:-/etc/bpfcompat/agent-load.env}"
LOAD_POLICY_EXAMPLE_PATH="${LOAD_POLICY_EXAMPLE_PATH:-/etc/bpfcompat/agent-load-policy.example.yaml}"

if [[ "${EUID:-$(id -u)}" -ne 0 ]]; then
  echo "run as root, for example: sudo $0" >&2
  exit 1
fi

if [[ ! -x "$BIN_SRC" ]]; then
  echo "missing executable $BIN_SRC; run make build or set BIN_SRC=/path/to/bpfcompat" >&2
  exit 1
fi

install -d -m 0755 "$PREFIX/bin"
install -m 0755 "$BIN_SRC" "$PREFIX/bin/bpfcompat"
if [[ -x "$VALIDATOR_SRC" ]]; then
  install -d -m 0755 "$(dirname "$VALIDATOR_DEST")"
  install -m 0755 "$VALIDATOR_SRC" "$VALIDATOR_DEST"
else
  echo "warning: validator binary not installed; reviewed host-load preflight will fail until VALIDATOR_SRC points to bpfcompat-validator" >&2
fi

if ! getent group "$AGENT_GROUP" >/dev/null; then
  groupadd --system "$AGENT_GROUP"
fi
if ! id -u "$AGENT_USER" >/dev/null 2>&1; then
  useradd --system --gid "$AGENT_GROUP" --home-dir /var/lib/bpfcompat-agent --shell /usr/sbin/nologin "$AGENT_USER"
fi

install -d -m 0750 -o "$AGENT_USER" -g "$AGENT_GROUP" /var/lib/bpfcompat-agent
install -d -m 0750 -o "$AGENT_USER" -g "$AGENT_GROUP" /var/log/bpfcompat-agent
install -d -m 0750 /etc/bpfcompat

if [[ ! -f "$ENV_PATH" ]]; then
  install -m 0600 packaging/systemd/bpfcompat-agent.env.example "$ENV_PATH"
  echo "created $ENV_PATH; edit it before enabling the timer" >&2
fi
if [[ ! -f "$LOAD_ENV_PATH" ]]; then
  install -m 0600 packaging/systemd/bpfcompat-agent-load.env.example "$LOAD_ENV_PATH"
  echo "created $LOAD_ENV_PATH; edit it only for reviewed host loads" >&2
fi
if [[ ! -f "$LOAD_POLICY_EXAMPLE_PATH" ]]; then
  install -m 0600 packaging/systemd/bpfcompat-agent-load-policy.example.yaml "$LOAD_POLICY_EXAMPLE_PATH"
  echo "created $LOAD_POLICY_EXAMPLE_PATH; copy/edit to /etc/bpfcompat/agent-load-policy.yaml before host loading" >&2
fi

install -m 0644 packaging/systemd/bpfcompat-agent.service /etc/systemd/system/bpfcompat-agent.service
install -m 0644 packaging/systemd/bpfcompat-agent.timer /etc/systemd/system/bpfcompat-agent.timer
install -m 0644 packaging/systemd/bpfcompat-agent-load.service /etc/systemd/system/bpfcompat-agent-load.service

systemctl daemon-reload

cat <<MSG
Installed bpfcompat agent alpha units.

Next:
  1. Edit $ENV_PATH
  2. Run preflight:
       sudo -u $AGENT_USER $PREFIX/bin/bpfcompat agent preflight --workdir /var/lib/bpfcompat-agent --out-dir /var/lib/bpfcompat-agent/selected --check-host-probe=false
  3. Run a single fetch-only check:
       systemctl start bpfcompat-agent.service
       journalctl -u bpfcompat-agent.service -n 80 --no-pager
  4. Enable scheduled checks:
       systemctl enable --now bpfcompat-agent.timer

Reviewed host loading is separate and disabled by default:
  1. Copy/edit $LOAD_POLICY_EXAMPLE_PATH to /etc/bpfcompat/agent-load-policy.yaml
  2. Edit $LOAD_ENV_PATH
  3. Start one reviewed load:
       systemctl start bpfcompat-agent-load.service
MSG
