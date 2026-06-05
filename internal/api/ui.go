package api

const uiHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>bpfcompat UI</title>
  <style nonce="__CSP_NONCE__">
    :root { color-scheme: light dark; }
    body {
      margin: 0;
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, sans-serif;
      background: #0f1115;
      color: #e8ebf1;
    }
    .preview-banner {
      border-bottom: 1px solid #2f3748;
      background: #1f2530;
      color: #ffdc8a;
      padding: 10px 14px;
      font-size: 12px;
      letter-spacing: 0;
    }
    .layout {
      display: grid;
      grid-template-columns: 420px 1fr;
      gap: 14px;
      height: calc(100vh - 44px);
      padding: 14px;
      box-sizing: border-box;
    }
    .panel {
      border: 1px solid #2c3340;
      border-radius: 8px;
      background: #151a22;
      overflow: auto;
      min-width: 0;
    }
    .panel h2 {
      margin: 0;
      padding: 12px 14px;
      font-size: 14px;
      border-bottom: 1px solid #2c3340;
    }
    .section { padding: 12px 14px; border-bottom: 1px solid #2c3340; }
    .section:last-child { border-bottom: 0; }
    .advanced-settings {
      border-bottom: 1px solid #2c3340;
    }
    .advanced-settings > summary {
      cursor: pointer;
      padding: 10px 14px;
      color: #cbd6ea;
      font-size: 12px;
      font-weight: 600;
      list-style: none;
    }
    .advanced-settings[open] > summary {
      border-bottom: 1px solid #2c3340;
    }
    .advanced-settings .section {
      border-bottom: 0;
    }
    .workflow-strip {
      display: grid;
      gap: 8px;
    }
    .workflow-steps {
      display: grid;
      grid-template-columns: repeat(5, minmax(0, 1fr));
      gap: 6px;
    }
    .workflow-step {
      border: 1px solid #334057;
      border-radius: 6px;
      background: #101723;
      padding: 7px 8px;
      font-size: 12px;
      color: #d8e2f4;
      text-align: center;
    }
    .step-title {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 8px;
      margin-bottom: 8px;
    }
    .step-title strong {
      font-size: 13px;
      color: #edf1f8;
    }
    .step-title span {
      color: #95a3bc;
      font-size: 11px;
    }
    label { display: block; font-size: 12px; color: #b5bfd1; margin: 8px 0 6px; }
    input, select, textarea, button {
      width: 100%;
      box-sizing: border-box;
      border-radius: 6px;
      border: 1px solid #3b4557;
      background: #0f131b;
      color: #edf1f8;
      padding: 8px 10px;
      font-size: 13px;
      letter-spacing: 0;
    }
    textarea { min-height: 120px; resize: vertical; font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; }
    button {
      background: #1d5bd8;
      border-color: #2d66d8;
      cursor: pointer;
      font-weight: 600;
    }
    button.secondary { background: #1c2431; border-color: #3b4557; }
    button.secondary.active {
      background: #17335f;
      border-color: #4d78bd;
      color: #e8f0ff;
    }
    button:disabled {
      cursor: not-allowed;
      opacity: 0.55;
    }
    .row { display: grid; grid-template-columns: 1fr 1fr; gap: 8px; }
    .segmented {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 8px;
      margin-bottom: 8px;
    }
    .target-presets {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 8px;
      margin-bottom: 8px;
    }
    .quad-actions {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 8px;
      margin-top: 8px;
    }
    .profiles { max-height: 220px; overflow: auto; border: 1px solid #2f3748; border-radius: 6px; padding: 8px; }
    .profile-header {
      display: grid;
      grid-template-columns: 18px 1fr auto;
      gap: 8px;
      align-items: center;
      color: #95a3bc;
      font-size: 11px;
      margin: 8px 0 4px;
      padding: 0 8px;
    }
    .profile {
      display: grid;
      grid-template-columns: 18px 1fr auto;
      align-items: center;
      gap: 8px;
      margin-bottom: 6px;
      font-size: 12px;
    }
    .profile .meta { color: #95a3bc; }
    .profile input[type="checkbox"] { width: 14px; height: 14px; margin: 0; }
    .mono { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; }
    .status {
      font-size: 13px;
      margin-bottom: 8px;
      color: #9fd5a7;
    }
    .status.error { color: #f2a3a3; }
    .verdict-bar {
      border: 1px solid #334057;
      border-radius: 8px;
      background: #111722;
      padding: 12px;
      display: grid;
      gap: 4px;
    }
    .verdict-bar.pass { border-color: #2f7b58; background: #102119; }
    .verdict-bar.fail, .verdict-bar.error { border-color: #8b4a4a; background: #241516; }
    .verdict-bar.running { border-color: #3863ab; background: #111d31; }
    .verdict-title {
      font-size: 16px;
      font-weight: 700;
      color: #edf1f8;
    }
    .verdict-meta {
      font-size: 12px;
      color: #b8c5da;
      line-height: 1.4;
    }
    .hint {
      margin-top: 8px;
      font-size: 12px;
      color: #c5d0e6;
      line-height: 1.35;
      white-space: pre-line;
    }
    .hint.error {
      color: #f2a3a3;
    }
    .results {
      padding: 12px 14px;
      display: grid;
      grid-template-rows: auto auto auto 1fr auto;
      gap: 10px;
      height: calc(100% - 52px);
      box-sizing: border-box;
    }
    .progress-wrap {
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #0d121a;
      padding: 8px;
      display: grid;
      gap: 6px;
    }
    .progress-track {
      height: 10px;
      border-radius: 999px;
      background: #1a2230;
      overflow: hidden;
    }
    .progress-fill {
      height: 100%;
      width: 0%;
      background: #2d66d8;
      transition: width 250ms ease;
    }
    .progress-meta {
      font-size: 12px;
      color: #b9c7df;
    }
    .progress-profiles {
      display: flex;
      flex-wrap: wrap;
      gap: 6px;
      max-height: 90px;
      overflow: auto;
    }
    .progress-pill {
      font-size: 11px;
      padding: 2px 6px;
      border-radius: 999px;
      border: 1px solid #2f3748;
      background: #0f131b;
      color: #cfd9ec;
    }
    .progress-pill.running { border-color: #3863ab; color: #9fc1ff; }
    .progress-pill.pass { border-color: #2f7b58; color: #94d7b0; }
    .progress-pill.fail, .progress-pill.infra_error { border-color: #8b4a4a; color: #f3b1b1; }
    .intent-options {
      display: grid;
      gap: 8px;
    }
    .intent-card {
      border: 1px solid #334057;
      border-radius: 6px;
      background: #101723;
      padding: 9px 10px;
      display: grid;
      grid-template-columns: 18px 1fr;
      gap: 8px;
      align-items: start;
      margin: 0;
    }
    .intent-card.active {
      border-color: #4d78bd;
      background: #13223a;
    }
    .intent-card.disabled {
      opacity: 0.6;
    }
    .intent-card input {
      width: 14px;
      height: 14px;
      margin: 1px 0 0;
    }
    .intent-label strong {
      display: block;
      font-size: 13px;
      color: #edf1f8;
    }
    .intent-label span {
      display: block;
      font-size: 12px;
      color: #aebbd0;
      line-height: 1.35;
      margin-top: 2px;
    }
    .matrix-wrap {
      overflow: auto;
      border: 1px solid #2f3748;
      border-radius: 6px;
    }
    .matrix-wrap table {
      min-width: 760px;
    }
    .matrix-counts {
      display: grid;
      grid-template-columns: repeat(4, minmax(0, 1fr));
      gap: 8px;
      margin-bottom: 8px;
    }
    .matrix-count {
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #101723;
      padding: 8px;
      display: grid;
      gap: 2px;
      min-width: 0;
    }
    .matrix-count strong {
      font-size: 16px;
      color: #edf1f8;
    }
    .matrix-count span {
      font-size: 11px;
      color: #aebbd0;
    }
    .matrix-count.pass { border-color: #2f7b58; background: #102119; }
    .matrix-count.fail, .matrix-count.error { border-color: #8b4a4a; background: #241516; }
    .matrix-count.check { border-color: #856b2c; background: #221d11; }
    .matrix-required-fail td {
      background: rgba(139, 74, 74, 0.16);
    }
    .matrix-row-pass td:first-child {
      border-left: 3px solid #2f7b58;
    }
    .matrix-row-fail td:first-child,
    .matrix-row-error td:first-child,
    .matrix-row-infra_error td:first-child {
      border-left: 3px solid #8b4a4a;
    }
    .matrix-status-pill {
      display: inline-block;
      min-width: 66px;
      border: 1px solid #3b4557;
      border-radius: 999px;
      padding: 2px 8px;
      text-align: center;
      font-size: 11px;
      font-weight: 800;
      letter-spacing: 0;
    }
    .matrix-status-pill.pass {
      border-color: #2f7b58;
      background: #102119;
      color: #8fe0ab;
    }
    .matrix-status-pill.fail,
    .matrix-status-pill.error,
    .matrix-status-pill.infra_error {
      border-color: #8b4a4a;
      background: #241516;
      color: #f1a0a0;
    }
    .matrix-status-pill.running {
      border-color: #3863ab;
      background: #111d31;
      color: #9fc1ff;
    }
    .ci-mode-panel {
      border: 1px solid #334057;
      border-radius: 6px;
      background: #101723;
      padding: 10px;
      display: grid;
      gap: 7px;
      margin-bottom: 10px;
    }
    .ci-mode-panel strong {
      font-size: 13px;
      color: #edf1f8;
    }
    .ci-mode-panel p {
      margin: 0;
      color: #b8c5da;
      font-size: 12px;
      line-height: 1.35;
    }
    .ci-mode-pills {
      display: flex;
      flex-wrap: wrap;
      gap: 6px;
    }
    .target-warning {
      color: #f3c777;
      margin-top: 2px;
    }
    .suite-preview {
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #0f131b;
      padding: 8px;
      margin-top: 8px;
      display: grid;
      gap: 8px;
    }
    .suite-preview table {
      margin-top: 2px;
    }
    table {
      border-collapse: collapse;
      width: 100%;
      font-size: 12px;
    }
    th, td {
      border: 1px solid #2f3748;
      padding: 6px 7px;
      text-align: left;
      vertical-align: top;
    }
    th { background: #19202b; }
    pre {
      margin: 0;
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #0d121a;
      padding: 10px;
      overflow: auto;
      font-size: 12px;
      line-height: 1.35;
      white-space: pre-wrap;
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    .history-table-wrap { max-height: 240px; overflow: auto; border: 1px solid #2f3748; border-radius: 6px; }
    .badge { font-size: 11px; padding: 2px 6px; border: 1px solid #2f3748; border-radius: 6px; background: #111722; color: #c8d2e6; }
    .split-actions { display: grid; grid-template-columns: 1fr 1fr; gap: 8px; margin-top: 8px; }
    .runtime-flow {
      border: 1px solid #2f3748;
      border-radius: 6px;
      padding: 10px;
      display: grid;
      gap: 8px;
    }
    .runtime-steps {
      display: grid;
      grid-template-columns: repeat(4, minmax(0, 1fr));
      gap: 6px;
    }
    .runtime-step {
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #0f131b;
      color: #95a3bc;
      font-size: 11px;
      padding: 5px 6px;
      text-align: center;
    }
    .runtime-step.active {
      border-color: #3863ab;
      color: #c8dbff;
      background: #122237;
    }
    .runtime-step.done {
      border-color: #2f7b58;
      color: #9bd9b4;
      background: #12231b;
    }
    .runtime-step.blocked {
      border-color: #8b4a4a;
      color: #f3b1b1;
      background: #2a1717;
    }
    .runtime-modes {
      display: grid;
      grid-template-columns: repeat(4, minmax(0, 1fr));
      gap: 8px;
    }
    .runtime-mode-btn.active {
      background: #1d5bd8;
      border-color: #2d66d8;
      color: #edf1f8;
    }
    .runtime-output > summary {
      cursor: pointer;
      font-size: 12px;
      color: #b5bfd1;
      margin-bottom: 6px;
      list-style: none;
    }
    .result-drilldown {
      border: 1px solid #2f3748;
      border-radius: 6px;
      background: #111722;
      overflow: hidden;
    }
    .result-drilldown > summary {
      cursor: pointer;
      padding: 9px 10px;
      color: #cbd6ea;
      font-size: 12px;
      font-weight: 600;
      list-style: none;
      border-bottom: 1px solid #2f3748;
    }
    .result-drilldown:not([open]) > summary { border-bottom: 0; }
    .result-drilldown > pre { border: 0; border-radius: 0; }
    .summary-cell-pass { color: #8fe0ab; font-weight: 700; }
    .summary-cell-fail, .summary-cell-error, .summary-cell-infra_error { color: #f1a0a0; font-weight: 700; }
    .summary-cell-running { color: #9fc1ff; font-weight: 700; }
    .hidden { display: none; }
    .mt8 { margin-top: 8px; }
    .mt10 { margin-top: 10px; }
    .inline-checkbox { width: auto; margin-right: 6px; }
    .clickable-row { cursor: pointer; }
    @media (max-width: 980px) {
      .layout {
        grid-template-columns: 1fr;
        height: auto;
        min-height: calc(100vh - 44px);
      }
      .panel {
        overflow: visible;
      }
      .results {
        height: auto;
        grid-template-rows: auto;
      }
      .workflow-steps {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
      .matrix-counts {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
      .runtime-steps,
      .runtime-modes {
        grid-template-columns: repeat(2, minmax(0, 1fr)) !important;
      }
    }
    @media (max-width: 560px) {
      .preview-banner {
        padding: 8px 10px;
      }
      .layout {
        padding: 8px;
        gap: 8px;
      }
      .section {
        padding: 10px;
      }
      .workflow-steps,
      .row,
      .segmented,
      .target-presets,
      .quad-actions,
      .matrix-counts,
      .runtime-steps,
      .runtime-modes,
      .split-actions {
        grid-template-columns: 1fr !important;
      }
      .step-title {
        align-items: flex-start;
        flex-direction: column;
        gap: 3px;
      }
      .matrix-wrap table {
        min-width: 620px;
      }
      .verdict-title {
        font-size: 15px;
      }
    }
  </style>
</head>
<body>
  <div class="preview-banner">Technical Preview — CI-first eBPF compatibility gate. Select target kernels, run validation, inspect drill-down evidence. Not for production runtime loading.</div>
  <div class="layout">
    <div class="panel">
      <h2>Compatibility Run Builder</h2>
      <div class="section workflow-strip">
        <div class="workflow-steps">
          <div class="workflow-step">1. Select targets</div>
          <div class="workflow-step">2. Provide BPF</div>
          <div class="workflow-step">3. Test intent</div>
          <div class="workflow-step">4. Run</div>
          <div class="workflow-step">5. Matrix</div>
        </div>
        <div class="hint">Built for the CI workflow Samy and Falco both pointed at: choose kernels, provide BPF objects, run the gate, inspect failures only when needed.</div>
      </div>

      <div class="section">
        <div class="step-title">
          <strong>1. Select Targets</strong>
          <span>kernel/distro matrix</span>
        </div>
        <div class="target-presets" id="targetPresets">
          <button type="button" class="secondary" data-preset="enterprise-broad">Enterprise Broad</button>
          <button type="button" class="secondary" data-preset="ubuntu-lts">Ubuntu LTS</button>
          <button type="button" class="secondary" data-preset="rhel-like">RHEL-like</button>
          <button type="button" class="secondary" data-preset="aws">AWS</button>
          <button type="button" class="secondary" data-preset="custom">Custom</button>
        </div>
        <div id="targetPresetHint" class="hint">Loading target catalog...</div>
        <div class="profile-header">
          <span>Run</span>
          <span>Kernel / distro profile</span>
          <span>Required</span>
        </div>
        <div id="profiles" class="profiles"></div>
        <div class="quad-actions">
          <button type="button" class="secondary" id="selectAll">Select All</button>
          <button type="button" class="secondary" id="clearAll">Clear All</button>
          <button type="button" class="secondary" id="requireSelected">Require Selected</button>
          <button type="button" class="secondary" id="clearRequired">Clear Required</button>
        </div>
      </div>

      <div class="section">
        <div class="step-title">
          <strong>2. Provide BPF</strong>
          <span>object or collection</span>
        </div>
        <div class="segmented">
          <button type="button" class="secondary active" id="modeSingle">Single Object</button>
          <button type="button" class="secondary" id="modeSuite">Collection / Suite</button>
        </div>

        <div id="singleInputMode">
          <label>Artifact Name</label>
          <input id="artifactName" placeholder="execsnoop">
          <label>BPF Input</label>
          <div class="row">
            <button type="button" class="secondary active" id="modeArtifact">Upload .bpf.o</button>
            <button type="button" class="secondary" id="modeSource">Compile Source</button>
          </div>
          <div id="artifactMode">
            <label>Artifact File</label>
            <input id="artifactFile" type="file">
          </div>
          <div id="sourceMode" class="hidden">
            <label>Source File</label>
            <input id="sourceFile" type="file">
            <label>Paste Source Code</label>
            <textarea id="sourceCode" placeholder="Paste .bpf.c source"></textarea>
            <label>Clang Flags (optional)</label>
            <input id="clangFlags" placeholder="-DDEBUG=1">
          </div>
        </div>

        <div id="suiteInputMode" class="hidden">
          <div class="ci-mode-panel">
            <strong>CI mode</strong>
            <p>Collections are designed to run in GitHub Actions on a self-hosted Linux/KVM runner. The suite lists BPF objects, manifests, and target policy; the output is a pass/fail matrix plus a detailed job summary.</p>
            <div class="ci-mode-pills">
              <span class="badge">suite YAML</span>
              <span class="badge">self-hosted KVM</span>
              <span class="badge">GitHub summary</span>
            </div>
          </div>
          <label>Suite YAML File</label>
          <input id="suiteFile" type="file" accept=".yaml,.yml,text/yaml">
          <label>or Paste Suite YAML</label>
          <textarea id="suiteText" placeholder="name: my-bpf-suite
defaults:
  matrix: matrices/dev-one.yaml
cases:
  - name: exec-tracepoint
    artifact: build/exec_tracepoint.bpf.o
    manifest: manifests/exec_tracepoint.yaml
  - name: network-xdp
    artifact: build/network_xdp.bpf.o"></textarea>
          <label>Suite Path in CI</label>
          <input id="suitePath" value="suites/project.yaml">
          <div id="suitePreview" class="suite-preview">
            <div class="hint">Paste a suite YAML to preview collection cases and generate GitHub Action configuration.</div>
          </div>
          <label>GitHub Action Preview</label>
          <pre id="suiteActionYaml" class="mono">Paste a suite YAML to generate a CI snippet.</pre>
          <button type="button" class="secondary mt8" id="copyActionYaml">Copy GitHub Action YAML</button>
        </div>
      </div>

      <details class="advanced-settings">
        <summary>Advanced single-object metadata</summary>
        <div class="section">
          <div class="row">
            <div>
              <label>Artifact Version</label>
              <input id="artifactVersion" placeholder="v1.0.0">
            </div>
            <div>
              <label>Artifact Variant</label>
              <input id="artifactVariant" placeholder="ringbuf-modern">
            </div>
          </div>
          <label>Artifact URI (optional)</label>
          <input id="artifactURI" placeholder="https://object-store.example.com/execsnoop-v1.0.0.bpf.o">
        </div>
      </details>

      <div class="section">
        <div class="step-title">
          <strong>3. Test Intent</strong>
          <span>what the gate proves</span>
        </div>
        <div class="intent-options">
          <label class="intent-card active" id="intentLoadAttach">
            <input type="radio" name="testIntent" value="load_attach" checked>
            <span class="intent-label">
              <strong>Load + attach</strong>
              <span>Default web path: libbpf load plus best-effort attach evidence.</span>
            </span>
          </label>
          <label class="intent-card disabled" id="intentLoadOnly">
            <input type="radio" name="testIntent" value="load_only" disabled>
            <span class="intent-label">
              <strong>Load only</strong>
              <span>Available in lower-level validator/agent flows, not exposed by this web gate yet.</span>
            </span>
          </label>
          <label class="intent-card disabled" id="intentBehavior">
            <input type="radio" name="testIntent" value="behavior" disabled>
            <span class="intent-label">
              <strong>Load + attach + behavior command</strong>
              <span>Use suite mode in CI when a collection needs Falco-style event or smoke-test assertions.</span>
            </span>
          </label>
        </div>
        <div id="testIntentHint" class="hint">This web gate currently proves load and attach compatibility. Behavior assertions belong in suite manifests and CI.</div>
      </div>

      <details class="advanced-settings">
        <summary>Advanced manifest and run settings</summary>
        <div class="section">
          <label>Manifest File (optional)</label>
          <input id="manifestFile" type="file">
          <label>or Manifest Text</label>
          <textarea id="manifestText" placeholder="name: demo
programs:
  - name: prog
    section: tracepoint/syscalls/sys_enter_execve"></textarea>
          <div id="writeAuthSection" class="hidden">
            <input id="writeApiKey" type="hidden" autocomplete="off">
            <input id="writeIdentityToken" type="hidden" autocomplete="off">
          </div>
          <div id="authHint" class="hint"></div>
          <div class="row">
            <div>
              <label>Timeout</label>
              <input id="timeout" value="8m">
            </div>
            <div>
              <label>Concurrency</label>
              <input id="concurrency" value="2">
            </div>
          </div>
        </div>
      </details>

      <div class="section">
        <div class="step-title">
          <strong>4. Run</strong>
          <span>gate selected targets</span>
        </div>
        <button id="runBtn">Run Compatibility Gate</button>
        <div id="runHint" class="hint">Single-object mode runs directly here. Collection mode generates the recommended CI suite configuration.</div>
      </div>
    </div>

    <div class="panel">
      <h2>Compatibility Matrix</h2>
      <div class="results">
        <div id="verdictBar" class="verdict-bar neutral">
          <div id="verdictTitle" class="verdict-title">No validation run yet</div>
          <div id="verdictMeta" class="verdict-meta">Select targets, provide a BPF object, then run the gate.</div>
        </div>
        <div id="status" class="status">Select target kernels and run validation.</div>
        <div class="progress-wrap">
          <div class="progress-track"><div id="progressFill" class="progress-fill"></div></div>
          <div id="progressMeta" class="progress-meta">0%</div>
          <div id="progressProfiles" class="progress-profiles"></div>
        </div>
        <div id="summary"></div>
        <details class="result-drilldown">
          <summary>Technical JSON report</summary>
          <pre id="resultJson" class="mono">{}</pre>
        </details>

        <details class="result-drilldown" id="evidenceDrilldown">
          <summary>Advanced evidence and history</summary>
          <div class="section">
          <div class="row">
            <div>
              <label>Artifact Filter</label>
              <input id="historyArtifactName" placeholder="execsnoop">
            </div>
            <div>
              <label>History Limit</label>
              <input id="historyLimit" value="100">
            </div>
          </div>
          <div class="split-actions">
            <button type="button" class="secondary" id="refreshHistory">Refresh History</button>
            <button type="button" id="runCompare">Compare Versions</button>
          </div>
          <div class="row">
            <div>
              <label>Base Version</label>
              <select id="baseVersion"></select>
            </div>
            <div>
              <label>Head Version</label>
              <select id="headVersion"></select>
            </div>
          </div>
          <div class="history-table-wrap">
            <table>
              <thead>
                <tr>
                  <th>Artifact</th>
                  <th>Version</th>
                  <th>Status</th>
                  <th>Req Pass/Fail</th>
                  <th>Created</th>
                </tr>
              </thead>
              <tbody id="historyRows"></tbody>
            </table>
          </div>
          <pre id="compareJson" class="mono mt8">{}</pre>

          <div class="mt10">
            <div class="row">
              <div>
                <label>Runtime Decision Limit</label>
                <input id="decisionLimit" value="100">
              </div>
            </div>
            <div class="split-actions">
              <button type="button" class="secondary" id="refreshDecisions">Refresh Runtime Decisions</button>
            </div>
            <div class="history-table-wrap">
              <table>
                <thead>
                  <tr>
                    <th>Decision</th>
                    <th>Operation</th>
                    <th>Status</th>
                    <th>Artifact</th>
                    <th>Version</th>
                    <th>Created</th>
                  </tr>
                </thead>
                <tbody id="decisionRows"></tbody>
              </table>
            </div>
            <pre id="decisionJson" class="mono mt8">{}</pre>
          </div>

          <div class="runtime-flow mt10">
            <div class="runtime-steps">
              <div class="runtime-step" id="runtimeStepProbe">1. Probe</div>
              <div class="runtime-step" id="runtimeStepSelect">2. Select</div>
              <div class="runtime-step" id="runtimeStepFetch">3. Fetch</div>
              <div class="runtime-step hidden" id="runtimeStepExecute">4. Execute</div>
            </div>

            <div class="runtime-modes" role="tablist" aria-label="Runtime Operation Mode">
              <button type="button" class="secondary runtime-mode-btn" id="runtimeModeProbe" aria-pressed="true">Probe</button>
              <button type="button" class="secondary runtime-mode-btn" id="runtimeModeSelect" aria-pressed="false">Select</button>
              <button type="button" class="secondary runtime-mode-btn" id="runtimeModeFetch" aria-pressed="false">Fetch</button>
              <button type="button" class="secondary runtime-mode-btn hidden" id="runtimeModeExecute" aria-pressed="false">Execute</button>
            </div>

            <div id="runtimeHint" class="hint"></div>

            <div class="row">
              <div>
                <label>Runtime Artifact Name</label>
                <input id="runtimeArtifactName" list="runtimeArtifactNames" placeholder="defaults to Artifact Name">
                <datalist id="runtimeArtifactNames"></datalist>
              </div>
              <div>
                <label>Runtime Version (optional)</label>
                <input id="runtimeVersion" placeholder="v1.0.0">
              </div>
            </div>

            <div class="row">
              <div>
                <label>Runtime Target Profile (optional)</label>
                <input id="runtimeTargetProfile" placeholder="ubuntu-22.04-5.15">
              </div>
              <div id="runtimeAttachModeWrap" class="hidden">
                <label>Runtime Attach Mode (optional)</label>
                <input id="runtimeAttachMode" placeholder="best-effort">
              </div>
            </div>

            <div id="runtimeExecuteFields" class="hidden">
              <div class="row">
                <div>
                  <label>Runtime Tenant (required for execute)</label>
                  <input id="runtimeTenant" placeholder="acme">
                </div>
                <div>
                  <label>Runtime Project (required for execute)</label>
                  <input id="runtimeProject" placeholder="aegis-bpf">
                </div>
              </div>
              <div class="row">
                <div>
                  <label>Registry Bearer Token (required for execute)</label>
                  <input id="runtimeRegistryToken" type="password" placeholder="Authorization: Bearer ...">
                </div>
                <div>
                  <label>Execute Approval Token (required for execute)</label>
                  <input id="runtimeApprovalToken" type="password" placeholder="X-Execute-Approval-Token">
                </div>
              </div>
              <div class="row">
                <div>
                  <label>Approved By (optional)</label>
                  <input id="runtimeApprovedBy" placeholder="operator@example.com">
                </div>
              </div>
            </div>

            <div class="row">
              <div id="runtimeProbeFeaturesWrap" class="hidden">
                <label><input id="runtimeProbeFeatures" class="inline-checkbox" type="checkbox" checked>probe_features</label>
              </div>
              <div id="runtimeRequireVerifiedHistoryWrap" class="hidden">
                <label><input id="runtimeRequireVerifiedHistory" class="inline-checkbox" type="checkbox" checked>require_verified_history</label>
              </div>
            </div>

            <button type="button" id="runtimeActionBtn">Probe Host</button>

            <details class="runtime-output" open>
              <summary>Technical Output</summary>
              <pre id="runtimeJson" class="mono">{}</pre>
            </details>
          </div>
          </div>
        </details>
      </div>
    </div>
  </div>

  <script nonce="__CSP_NONCE__">
    let mode = "artifact";
    let bpfInputMode = "single";
    let selectedPreset = "ubuntu-lts";
    const state = { profiles: [], history: [], decisions: [], suite: { name: "", cases: [] } };
    let apiConfig = null;
    let runInFlight = false;
    let evidenceLoaded = false;

    const byId = (id) => document.getElementById(id);
    const statusEl = byId("status");
    const verdictBarEl = byId("verdictBar");
    const verdictTitleEl = byId("verdictTitle");
    const verdictMetaEl = byId("verdictMeta");
    const resultJsonEl = byId("resultJson");
    const compareJsonEl = byId("compareJson");
    const decisionJsonEl = byId("decisionJson");
    const runtimeJsonEl = byId("runtimeJson");
    const progressFillEl = byId("progressFill");
    const progressMetaEl = byId("progressMeta");
    const progressProfilesEl = byId("progressProfiles");
    const runBtnEl = byId("runBtn");
    const runHintEl = byId("runHint");
    const targetPresetHintEl = byId("targetPresetHint");
    const suitePreviewEl = byId("suitePreview");
    const suiteActionYamlEl = byId("suiteActionYaml");
    const evidenceDrilldownEl = byId("evidenceDrilldown");
    const writeAuthSectionEl = byId("writeAuthSection");
    const authHintEl = byId("authHint");
    const runtimeHintEl = byId("runtimeHint");
    const runtimeArtifactNamesEl = byId("runtimeArtifactNames");
    const runtimeActionBtn = byId("runtimeActionBtn");
    const runtimeModeButtons = {
      probe: byId("runtimeModeProbe"),
      select: byId("runtimeModeSelect"),
      fetch: byId("runtimeModeFetch"),
      execute: byId("runtimeModeExecute")
    };
    const runtimeStepElements = {
      probe: byId("runtimeStepProbe"),
      select: byId("runtimeStepSelect"),
      fetch: byId("runtimeStepFetch"),
      execute: byId("runtimeStepExecute")
    };
    const runtimeStepsEl = document.querySelector(".runtime-steps");
    const runtimeModesEl = document.querySelector(".runtime-modes");
    const demoRuntimeExecuteDefaults = {
      tenant: "acme",
      project: "aegis-bpf",
      registryToken: "",
      approvalToken: "",
      approvedBy: "live-demo"
    };
    const exposeRuntimeExecuteInWebUI = false;
    let runtimeMode = "probe";
    const runtimeCompletedSteps = {
      probe: false,
      select: false,
      fetch: false,
      execute: false
    };

    function setStatus(text, error = false) {
      statusEl.textContent = text;
      statusEl.className = error ? "status error" : "status";
    }

    function setVerdict(kind, title, meta) {
      verdictBarEl.className = "verdict-bar " + (kind || "neutral");
      verdictTitleEl.textContent = title || "No validation run yet";
      verdictMetaEl.textContent = meta || "";
    }

    function setAuthHint(text, error = false) {
      authHintEl.textContent = text || "";
      authHintEl.className = error ? "hint error" : "hint";
    }

    function setRuntimeHint(text, error = false) {
      runtimeHintEl.textContent = text || "";
      runtimeHintEl.className = error ? "hint error" : "hint";
    }

    function deriveProfileHintFromProbe(probe) {
      const osID = String(probe && probe.os && probe.os.id || "").trim();
      const versionID = String(probe && probe.os && probe.os.version_id || "").trim();
      const release = String(probe && probe.kernel && probe.kernel.release || "").trim();
      const match = release.match(/^([0-9]+)\.([0-9]+)/);
      if (!osID || !versionID || !match) {
        return "";
      }
      return osID + "-" + versionID + "-" + match[1] + "." + match[2];
    }

    function modeLabel(mode) {
      switch (mode) {
        case "probe": return "Probe Host";
        case "select": return "Runtime Select";
        case "fetch": return "Runtime Fetch";
        case "execute": return "Runtime Execute";
        default: return "Run";
      }
    }

    function modeHelp(mode) {
      switch (mode) {
        case "probe":
          return "Detect host kernel and feature support; pre-fills target profile hint when possible.";
        case "select":
          return "Operator-only path: choose the best artifact variant from compatibility history.";
        case "fetch":
          return "Operator-only path: select and retrieve artifact payload with history verification.";
        case "execute":
          return "Operator-only path: controlled host load with explicit approval.";
        default:
          return "";
      }
    }

    function shouldApplyDemoRuntimeExecuteDefaults() {
      if (!apiConfig) {
        return false;
      }
      if (!exposeRuntimeExecuteInWebUI) {
        return false;
      }
      return !!apiConfig.allow_anonymous_write &&
        !!apiConfig.runtime_execute_enabled &&
        !apiConfig.write_api_key_configured &&
        !apiConfig.write_identity_verifier_enabled;
    }

    function setInputIfBlank(id, value) {
      const el = byId(id);
      if (!el) {
        return;
      }
      if (!el.value.trim()) {
        el.value = value;
      }
    }

    function applyDemoRuntimeExecuteDefaults() {
      if (!shouldApplyDemoRuntimeExecuteDefaults()) {
        return;
      }
      setInputIfBlank("runtimeTenant", demoRuntimeExecuteDefaults.tenant);
      setInputIfBlank("runtimeProject", demoRuntimeExecuteDefaults.project);
      setInputIfBlank("runtimeRegistryToken", demoRuntimeExecuteDefaults.registryToken);
      setInputIfBlank("runtimeApprovalToken", demoRuntimeExecuteDefaults.approvalToken);
      setInputIfBlank("runtimeApprovedBy", demoRuntimeExecuteDefaults.approvedBy);
    }

    function availableRuntimeArtifactNames() {
      const seen = new Set();
      const names = [];
      state.history.forEach((rec) => {
        const name = String(rec && rec.artifact_name || "").trim();
        if (!name || seen.has(name)) {
          return;
        }
        seen.add(name);
        names.push(name);
      });
      return names;
    }

    function refreshRuntimeArtifactSuggestions() {
      const names = availableRuntimeArtifactNames();
      runtimeArtifactNamesEl.innerHTML = "";
      names.forEach((name) => {
        const option = document.createElement("option");
        option.value = name;
        runtimeArtifactNamesEl.appendChild(option);
      });
      if (!byId("runtimeArtifactName").value.trim() && names.length > 0) {
        byId("runtimeArtifactName").value = names[0];
      }
    }

    function enhanceRuntimeErrorMessage(mode, message) {
      if ((mode === "select" || mode === "fetch") && message.includes("no artifact versions found for")) {
        const names = availableRuntimeArtifactNames();
        if (names.length > 0) {
          return message + ". Try one of: " + names.slice(0, 8).join(", ");
        }
      }
      return message;
    }

    function renderRuntimeSteps() {
      const order = ["probe", "select", "fetch", "execute"];
      order.forEach((step) => {
        const el = runtimeStepElements[step];
        if (!el) {
          return;
        }
        el.classList.remove("active", "done", "blocked");
        if (runtimeCompletedSteps[step]) {
          el.classList.add("done");
        }
        if (runtimeMode === step) {
          el.classList.add("active");
        }
      });
      if (!runtimeDeliveryActionsAvailable()) {
        if (runtimeStepElements.select) {
          runtimeStepElements.select.classList.add("blocked");
        }
        if (runtimeStepElements.fetch) {
          runtimeStepElements.fetch.classList.add("blocked");
        }
      }
      if (runtimeStepElements.execute && (!exposeRuntimeExecuteInWebUI || !apiConfig || !apiConfig.runtime_execute_enabled || !writeActionsAvailable())) {
        runtimeStepElements.execute.classList.add("blocked");
      }
    }

    function syncRuntimeModeUI() {
      const runtimeDeliveryAllowed = runtimeDeliveryActionsAvailable();
      const executeAllowed = exposeRuntimeExecuteInWebUI && !!apiConfig && !!apiConfig.runtime_execute_enabled && writeActionsAvailable();
      if (runtimeStepsEl) {
        runtimeStepsEl.style.gridTemplateColumns = exposeRuntimeExecuteInWebUI ? "repeat(4, minmax(0, 1fr))" : "repeat(3, minmax(0, 1fr))";
      }
      if (runtimeModesEl) {
        runtimeModesEl.style.gridTemplateColumns = exposeRuntimeExecuteInWebUI ? "repeat(4, minmax(0, 1fr))" : "repeat(3, minmax(0, 1fr))";
      }
      if (runtimeStepElements.execute) {
        runtimeStepElements.execute.style.display = exposeRuntimeExecuteInWebUI ? "block" : "none";
      }
      if (runtimeMode === "select" || runtimeMode === "fetch") {
        if (!runtimeDeliveryAllowed) {
          runtimeMode = "probe";
        }
      }
      if (runtimeMode === "execute" && !executeAllowed) {
        runtimeMode = "probe";
      }

      Object.entries(runtimeModeButtons).forEach(([modeKey, btn]) => {
        if (!btn) {
          return;
        }
        if (modeKey === "select" || modeKey === "fetch") {
          btn.disabled = !runtimeDeliveryAllowed;
          btn.title = runtimeDeliveryAllowed ? "" : "Runtime delivery is not open on this public demo";
        }
        if (modeKey === "execute") {
          btn.disabled = !executeAllowed;
          btn.style.display = executeAllowed ? "block" : "none";
          btn.title = executeAllowed ? "" : "Runtime execute is disabled on this public demo";
        }
        const active = runtimeMode === modeKey;
        btn.classList.toggle("active", active);
        btn.setAttribute("aria-pressed", active ? "true" : "false");
      });

      const showExecute = runtimeMode === "execute";
      byId("runtimeExecuteFields").style.display = showExecute ? "block" : "none";
      byId("runtimeAttachModeWrap").style.display = showExecute ? "block" : "none";
      byId("runtimeProbeFeaturesWrap").style.display = showExecute ? "block" : "none";
      byId("runtimeRequireVerifiedHistoryWrap").style.display =
        (runtimeMode === "fetch" || runtimeMode === "execute") ? "block" : "none";
      if (showExecute) {
        applyDemoRuntimeExecuteDefaults();
      }

      runtimeActionBtn.textContent = modeLabel(runtimeMode);
      runtimeActionBtn.disabled = false;
      runtimeActionBtn.title = "";

      if (!runtimeDeliveryAllowed && runtimeMode === "probe") {
        setRuntimeHint("Public demo mode: probe is available here. Selection, fetch, and execute are operator-only; see Results for prepared selection evidence.");
      } else {
        setRuntimeHint(modeHelp(runtimeMode), runtimeMode === "execute" && !executeAllowed);
      }
      renderRuntimeSteps();
    }

    function setRuntimeMode(nextMode) {
      if (!nextMode || !runtimeModeButtons[nextMode]) {
        return;
      }
      if ((nextMode === "select" || nextMode === "fetch") && !runtimeDeliveryActionsAvailable()) {
        setRuntimeHint("Runtime " + nextMode + " is not open on this public demo. Use the Results page for prepared selector evidence.", true);
        return;
      }
      if (nextMode === "execute" && (!exposeRuntimeExecuteInWebUI || !apiConfig || !apiConfig.runtime_execute_enabled || !writeActionsAvailable())) {
        setRuntimeHint("Runtime execute is disabled on this public demo.", true);
        return;
      }
      runtimeMode = nextMode;
      syncRuntimeModeUI();
    }

    function sleep(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    }

    function resetProgress() {
      progressFillEl.style.width = "0%";
      progressMetaEl.textContent = "0%";
      progressProfilesEl.innerHTML = "";
    }

    function renderProgress(job) {
      const percent = Math.max(0, Math.min(100, Number(job && job.percent) || 0));
      progressFillEl.style.width = percent + "%";

      const details = [];
      if (job && job.completed_profiles && job.total_profiles) {
        details.push(job.completed_profiles + "/" + job.total_profiles + " profiles");
      }
      if (job && job.stage) {
        details.push(job.stage);
      }
      if (job && job.message) {
        details.push(job.message);
      }
      progressMetaEl.textContent = percent + "%" + (details.length ? " • " + details.join(" • ") : "");
      if (percent > 0 && percent < 100) {
        setVerdict("running", "Running compatibility gate", progressMetaEl.textContent);
      }

      progressProfilesEl.innerHTML = "";
      const statuses = (job && job.profile_statuses) || {};
      const ids = Object.keys(statuses).sort();
      ids.forEach((id) => {
        const pill = document.createElement("span");
        const state = String(statuses[id] || "").trim() || "pending";
        pill.className = "progress-pill " + state;
        pill.textContent = id + ": " + state;
        progressProfilesEl.appendChild(pill);
      });
    }

    async function decodeJSONResponse(res) {
      const raw = await res.text();
      if (!raw || !raw.trim()) {
        return { data: {}, raw: "" };
      }
      try {
        return { data: JSON.parse(raw), raw };
      } catch (err) {
        return { data: {}, raw };
      }
    }

    async function requestJSON(url, options) {
      const res = await fetch(url, options);
      const decoded = await decodeJSONResponse(res);
      if (!res.ok) {
        let message = "";
        if (decoded.data && typeof decoded.data.error === "string" && decoded.data.error.trim()) {
          message = decoded.data.error.trim();
        } else if (decoded.raw.trim()) {
          message = decoded.raw.trim();
        } else {
          message = res.statusText || "request failed";
        }
        throw new Error("HTTP " + res.status + ": " + message);
      }
      return decoded.data || {};
    }

    function buildWriteHeaders(baseHeaders = {}) {
      const headers = Object.assign({}, baseHeaders);
      const key = byId("writeApiKey").value.trim();
      if (key && apiConfig && apiConfig.write_api_key_configured) {
        headers["X-API-Key"] = key;
      }
      const identityToken = byId("writeIdentityToken").value.trim();
      if (identityToken && apiConfig && apiConfig.write_identity_verifier_enabled) {
        headers["X-API-Identity-Token"] = identityToken;
      }
      return headers;
    }

    function runtimeArtifactNameOrFallback() {
      const runtimeName = byId("runtimeArtifactName").value.trim();
      if (runtimeName) {
        return runtimeName;
      }
      return byId("artifactName").value.trim();
    }

    function runtimeCommonBody() {
      return {
        artifact_name: runtimeArtifactNameOrFallback(),
        version: byId("runtimeVersion").value.trim(),
        target_profile: byId("runtimeTargetProfile").value.trim()
      };
    }

    function buildRuntimeExecuteHeaders() {
      const headers = buildWriteHeaders({ "Content-Type": "application/json" });
      const registryToken = byId("runtimeRegistryToken").value.trim();
      if (registryToken) {
        headers["Authorization"] = "Bearer " + registryToken;
      }
      const approvalToken = byId("runtimeApprovalToken").value.trim();
      if (approvalToken) {
        headers["X-Execute-Approval-Token"] = approvalToken;
      }
      const approvedBy = byId("runtimeApprovedBy").value.trim();
      if (approvedBy) {
        headers["X-Execute-Approved-By"] = approvedBy;
      }
      return headers;
    }

    function hasWriteCredentials() {
      const key = byId("writeApiKey").value.trim();
      const identityToken = byId("writeIdentityToken").value.trim();
      return (!!key && apiConfig && apiConfig.write_api_key_configured) ||
        (!!identityToken && apiConfig && apiConfig.write_identity_verifier_enabled);
    }

    function writeActionsAvailable() {
      return !!apiConfig && (!!apiConfig.allow_anonymous_write || hasWriteCredentials());
    }

    function runtimeDeliveryActionsAvailable() {
      return writeActionsAvailable() || (!!apiConfig && !!apiConfig.allow_anonymous_runtime_delivery);
    }

    function requireWriteCredentials(actionLabel) {
      if (writeActionsAvailable()) {
        return;
      }
      if (!apiConfig) {
        throw new Error("Demo configuration is not loaded yet");
      }
      throw new Error(actionLabel + " is an operator-only action in this public demo.");
    }

    function requireRuntimeDeliveryAccess(actionLabel) {
      if (runtimeDeliveryActionsAvailable()) {
        return;
      }
      if (!apiConfig) {
        throw new Error("Demo configuration is not loaded yet");
      }
      throw new Error(actionLabel + " is not open on this public demo.");
    }

    function setButtonActive(id, active) {
      const el = byId(id);
      if (el) {
        el.classList.toggle("active", !!active);
      }
    }

    function switchMode(nextMode) {
      mode = nextMode;
      byId("artifactMode").style.display = mode === "artifact" ? "block" : "none";
      byId("sourceMode").style.display = mode === "source" ? "block" : "none";
      setButtonActive("modeArtifact", mode === "artifact");
      setButtonActive("modeSource", mode === "source");
    }

    function switchBPFInputMode(nextMode) {
      bpfInputMode = nextMode;
      byId("singleInputMode").style.display = bpfInputMode === "single" ? "block" : "none";
      byId("suiteInputMode").style.display = bpfInputMode === "suite" ? "block" : "none";
      setButtonActive("modeSingle", bpfInputMode === "single");
      setButtonActive("modeSuite", bpfInputMode === "suite");
      if (bpfInputMode === "suite") {
        runBtnEl.textContent = "Generate CI Gate";
        runHintEl.textContent = "Collection mode is CI-first. The browser previews the suite and generates the GitHub Action configuration.";
        setStatus("Collection mode selected. Paste suite YAML to preview the BPF object set.");
        setVerdict("neutral", "Collection preview mode", "Use the generated GitHub Action on a self-hosted Linux/KVM runner for real suite execution.");
        updateSuitePreview();
      } else {
        runBtnEl.textContent = "Run Compatibility Gate";
        runHintEl.textContent = "Single-object mode runs directly here. Collection mode generates the recommended CI suite configuration.";
        setStatus("Single-object mode selected. Upload or compile one BPF object.");
        setVerdict("neutral", "No validation run yet", "Select targets, provide a BPF object, then run the gate.");
      }
    }

    byId("modeSingle").addEventListener("click", () => switchBPFInputMode("single"));
    byId("modeSuite").addEventListener("click", () => switchBPFInputMode("suite"));
    byId("modeArtifact").addEventListener("click", () => switchMode("artifact"));
    byId("modeSource").addEventListener("click", () => switchMode("source"));
    runtimeModeButtons.probe.addEventListener("click", () => setRuntimeMode("probe"));
    runtimeModeButtons.select.addEventListener("click", () => setRuntimeMode("select"));
    runtimeModeButtons.fetch.addEventListener("click", () => setRuntimeMode("fetch"));
    runtimeModeButtons.execute.addEventListener("click", () => setRuntimeMode("execute"));

    byId("suiteText").addEventListener("input", updateSuitePreview);
    byId("suitePath").addEventListener("input", updateSuitePreview);
    byId("suiteFile").addEventListener("change", () => {
      const file = byId("suiteFile").files[0];
      if (!file) {
        return;
      }
      const reader = new FileReader();
      reader.onload = () => {
        byId("suiteText").value = String(reader.result || "");
        updateSuitePreview();
      };
      reader.onerror = () => setStatus("Failed to read suite file", true);
      reader.readAsText(file);
    });
    byId("copyActionYaml").addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(suiteActionYamlEl.textContent || "");
        setStatus("GitHub Action YAML copied");
      } catch (err) {
        setStatus("Copy failed; select the generated YAML manually.", true);
      }
    });

    document.querySelectorAll("input[name='testIntent']").forEach((input) => {
      input.addEventListener("change", () => {
        document.querySelectorAll(".intent-card").forEach((card) => card.classList.remove("active"));
        const card = input.closest(".intent-card");
        if (card) {
          card.classList.add("active");
        }
      });
    });

    function refreshAuthHintFromConfig() {
      if (!apiConfig) {
        setAuthHint("Loading demo settings...");
        return;
      }
      writeAuthSectionEl.style.display = "none";
      const lines = [];
      let hasBlocking = false;
      if (apiConfig.allow_anonymous_write) {
        lines.push("Public demo mode: validation and protected demo actions run without credentials.");
      } else if (apiConfig.allow_anonymous_validate) {
        lines.push("Public demo mode: validation runs without credentials.");
        if (apiConfig.allow_anonymous_runtime_delivery) {
          lines.push("Runtime select/fetch are open for demo delivery proof.");
        } else if (!runtimeDeliveryActionsAvailable()) {
          lines.push("Operator-only actions are hidden in this UI.");
        }
      } else if (apiConfig.write_api_key_configured || apiConfig.write_identity_verifier_enabled) {
        lines.push("Protected server: validation requires operator credentials outside this public UI.");
        hasBlocking = !hasWriteCredentials();
      } else {
        lines.push("Validation is unavailable: server has no public validation mode configured.");
        hasBlocking = true;
      }
      if (!apiConfig.runtime_execute_enabled) {
        lines.push("Runtime execute is disabled on this server.");
      }
      if (lines.length == 0) {
        lines.push("Demo configuration loaded.");
      }
      setAuthHint(lines.join("\n"), hasBlocking);
      const compareButton = byId("runCompare");
      if (compareButton) {
        compareButton.disabled = !writeActionsAvailable();
        compareButton.title = writeActionsAvailable() ? "" : "Operator-only in this public demo";
      }
      syncRuntimeModeUI();
    }

    async function refreshAPIConfig() {
      apiConfig = await requestJSON("/api/config");
      applyDemoRuntimeExecuteDefaults();
      refreshAuthHintFromConfig();
    }

    function createProfileRow(profile) {
      const row = document.createElement("div");
      row.className = "profile";
      const include = document.createElement("input");
      include.type = "checkbox";
      include.checked = !!profile.transport_supported;
      include.dataset.kind = "include";
      include.dataset.id = profile.id;
      include.dataset.transportSupported = profile.transport_supported ? "true" : "false";

      const label = document.createElement("div");
      const title = document.createElement("div");
      const strong = document.createElement("strong");
      strong.textContent = profile.id;
      title.appendChild(strong);
      const meta = document.createElement("div");
      meta.className = "meta";
      const profileParts = [];
      if (profile.distro) profileParts.push(profile.distro);
      if (profile.version) profileParts.push(profile.version);
      if (profile.kernel_family) profileParts.push("kernel " + profile.kernel_family);
      if (profile.arch) profileParts.push(profile.arch);
      meta.textContent = profileParts.length > 0 ? profileParts.join(" • ") : "kernel target";
      label.append(title, meta);
      if (!profile.transport_supported || profile.image_cached === false) {
        const reason = document.createElement("div");
        reason.className = "meta target-warning";
        reason.textContent = !profile.transport_supported ? (profile.transport_note || "Unavailable for this run") : "VM image not cached yet";
        label.appendChild(reason);
      }

      const required = document.createElement("input");
      required.type = "checkbox";
      required.checked = !!profile.required_default && !!profile.transport_supported;
      required.disabled = !profile.transport_supported;
      required.dataset.kind = "required";
      required.dataset.id = profile.id;

      include.addEventListener("change", () => {
        selectedPreset = "custom";
        syncTargetPresetButtons();
        required.disabled = !include.checked || !profile.transport_supported;
        if (required.disabled) {
          required.checked = false;
        }
        updateTargetPresetHint();
      });
      required.addEventListener("change", () => {
        selectedPreset = "custom";
        syncTargetPresetButtons();
        updateTargetPresetHint();
      });

      row.append(include, label, required);
      return row;
    }

    function syncTargetPresetButtons() {
      document.querySelectorAll("button[data-preset]").forEach((btn) => {
        btn.classList.toggle("active", btn.dataset.preset === selectedPreset);
      });
    }

    function profileMatchesPreset(profile, preset) {
      const id = String(profile.id || "").toLowerCase();
      const distro = String(profile.distro || "").toLowerCase();
      const version = String(profile.version || "").toLowerCase();
      if (!profile.transport_supported) {
        return false;
      }
      switch (preset) {
        case "ubuntu-lts":
          return distro === "ubuntu" && ["18.04", "20.04", "22.04", "24.04"].includes(version) && !id.includes("minimal");
        case "rhel-like":
          return ["rhel", "rocky", "almalinux", "centos", "centos-stream"].includes(distro);
        case "aws":
          return distro.includes("amazon") || id.includes("amazonlinux") || id.includes("bottlerocket");
        case "enterprise-broad":
          if (distro === "ubuntu") {
            return ["20.04", "22.04", "24.04"].includes(version) && !id.includes("minimal");
          }
          if (distro === "debian") {
            return ["12", "13"].includes(version);
          }
          if (["rhel", "rocky", "almalinux", "centos", "centos-stream"].includes(distro)) {
            return ["8", "9", "10"].includes(version);
          }
          return ["oracle", "oraclelinux", "sles", "opensuse"].includes(distro) || distro.includes("amazon");
        default:
          return false;
      }
    }

    function requiredForPreset(profile, preset, selectedCount) {
      if (!profile.transport_supported) {
        return false;
      }
      if (profile.required_default) {
        return true;
      }
      return selectedCount <= 12;
    }

    function applyTargetPreset(preset) {
      if (preset === "custom") {
        selectedPreset = "custom";
        syncTargetPresetButtons();
        updateTargetPresetHint();
        return;
      }
      selectedPreset = preset;
      const selected = state.profiles.filter((profile) => profileMatchesPreset(profile, preset));
      const selectedIDs = new Set(selected.map((profile) => profile.id));
      document.querySelectorAll("input[data-kind='include']").forEach((input) => {
        input.checked = selectedIDs.has(input.dataset.id);
      });
      document.querySelectorAll("input[data-kind='required']").forEach((input) => {
        const profile = state.profiles.find((p) => p.id === input.dataset.id);
        const include = selectedIDs.has(input.dataset.id);
        input.disabled = !include || !profile || !profile.transport_supported;
        input.checked = include && !!profile && requiredForPreset(profile, preset, selected.length);
      });
      syncTargetPresetButtons();
      updateTargetPresetHint();
    }

    function updateTargetPresetHint() {
      const picks = selectedProfiles();
      const label = selectedPreset === "enterprise-broad" ? "Enterprise Broad" :
        selectedPreset === "ubuntu-lts" ? "Ubuntu LTS" :
        selectedPreset === "rhel-like" ? "RHEL-like" :
        selectedPreset === "aws" ? "AWS" : "Custom";
      targetPresetHintEl.textContent = label + ": " + picks.include.length + " target(s) selected, " + picks.required.length + " required for the gate.";
    }

    function appendCell(tr, value, className = "") {
      const td = document.createElement("td");
      td.textContent = String(value);
      if (className) {
        td.className = className;
      }
      tr.appendChild(td);
    }

    function normalizeStatus(status) {
      const normalized = String(status || "").trim().toLowerCase().replace(/[^a-z0-9_]+/g, "_");
      return normalized || "unknown";
    }

    function appendStatusCell(tr, status) {
      const td = document.createElement("td");
      const normalized = normalizeStatus(status);
      const pill = document.createElement("span");
      pill.className = "matrix-status-pill " + normalized;
      pill.textContent = String(status || "-").toUpperCase();
      td.appendChild(pill);
      tr.appendChild(td);
    }

    function appendMatrixCount(container, label, value, className = "") {
      const item = document.createElement("div");
      item.className = "matrix-count" + (className ? " " + className : "");
      const strong = document.createElement("strong");
      strong.textContent = String(value);
      const span = document.createElement("span");
      span.textContent = label;
      item.append(strong, span);
      container.appendChild(item);
    }

    function summaryStatusClass(status) {
      return "summary-cell-" + normalizeStatus(status);
    }

    function formatProfileEnv(target) {
      const env = target && target.profile ? target.profile : null;
      if (!env) return "-";
      const parts = [];
      if (env.distro) parts.push(env.distro);
      if (env.version) parts.push(env.version);
      if (env.kernel_family) parts.push("kfamily=" + env.kernel_family);
      if (env.arch) parts.push("arch=" + env.arch);
      return parts.length ? parts.join(" ") : "-";
    }

    function formatHostKernel(target) {
      const env = target && target.host ? target.host : null;
      if (!env) return "-";
      const kernel = env.kernel || "-";
      if (env.arch) return kernel + " (" + env.arch + ")";
      return kernel;
    }

    function formatTargetReason(target) {
      if (!target) return "-";
      if (String(target.status || "").toLowerCase() === "pass") return "-";
      if (target.classification_code) return target.classification_code;
      if (target.failed_stage) return target.failed_stage;
      if (target.classification_reason) return target.classification_reason;
      return "failed";
    }

    function cleanYAMLValue(raw) {
      let value = String(raw || "").trim();
      if ((value.startsWith("\"") && value.endsWith("\"")) || (value.startsWith("'") && value.endsWith("'"))) {
        value = value.slice(1, -1);
      }
      return value.trim();
    }

    function parseSuitePreview(text) {
      const suite = { name: "", cases: [] };
      let current = null;
      String(text || "").split(/\r?\n/).forEach((line) => {
        const trimmed = line.trim();
        if (!trimmed || trimmed.startsWith("#")) {
          return;
        }
        let match = trimmed.match(/^name:\s*(.+)$/);
        if (match && !current && !suite.name) {
          suite.name = cleanYAMLValue(match[1]);
          return;
        }
        match = line.match(/^\s*-\s+name:\s*(.+)$/);
        if (match) {
          current = { name: cleanYAMLValue(match[1]), artifact: "", manifest: "", artifactName: "" };
          suite.cases.push(current);
          return;
        }
        if (!current) {
          return;
        }
        match = line.match(/^\s+(artifact|manifest|artifact_name):\s*(.+)$/);
        if (!match) {
          return;
        }
        const key = match[1];
        const value = cleanYAMLValue(match[2]);
        if (key === "artifact") {
          current.artifact = value;
        } else if (key === "manifest") {
          current.manifest = value;
        } else if (key === "artifact_name") {
          current.artifactName = value;
        }
      });
      return suite;
    }

    function generateSuiteActionYAML(suite) {
      const suitePath = byId("suitePath").value.trim() || "suites/project.yaml";
      const suiteName = suite && suite.name ? suite.name : "bpf-compatibility";
      return [
        "name: BPF compatibility",
        "",
        "on:",
        "  pull_request:",
        "  push:",
        "    branches: [main]",
        "",
        "jobs:",
        "  bpfcompat:",
        "    name: " + suiteName,
        "    runs-on: [self-hosted, linux, x64, kvm]",
        "    steps:",
        "      - uses: actions/checkout@v4",
        "      - uses: Kernel-Guard/bpfcompat@v0.1.2",
        "        with:",
        "          suite: " + suitePath,
        "          suite-out: reports/bpfcompat-suite.json",
        "          suite-markdown: reports/bpfcompat-suite.md",
        "          timeout: 8m",
        "          concurrency: \"1\""
      ].join("\n");
    }

    function updateSuitePreview() {
      const suite = parseSuitePreview(byId("suiteText").value);
      state.suite = suite;
      suitePreviewEl.replaceChildren();
      const title = document.createElement("div");
      title.className = "hint";
      if (suite.cases.length === 0) {
        title.textContent = "No suite cases detected yet. Paste a suite YAML with cases[].name and cases[].artifact.";
        suitePreviewEl.appendChild(title);
      } else {
        title.textContent = "Suite " + (suite.name || "unnamed") + ": " + suite.cases.length + " BPF object case(s).";
        suitePreviewEl.appendChild(title);
        const table = document.createElement("table");
        const thead = document.createElement("thead");
        const headRow = document.createElement("tr");
        ["Case", "Artifact", "Manifest"].forEach((name) => {
          const th = document.createElement("th");
          th.textContent = name;
          headRow.appendChild(th);
        });
        thead.appendChild(headRow);
        table.appendChild(thead);
        const tbody = document.createElement("tbody");
        suite.cases.forEach((c) => {
          const tr = document.createElement("tr");
          appendCell(tr, c.artifactName || c.name || "-");
          appendCell(tr, c.artifact || "-");
          appendCell(tr, c.manifest || "-");
          tbody.appendChild(tr);
        });
        table.appendChild(tbody);
        suitePreviewEl.appendChild(table);
      }
      suiteActionYamlEl.textContent = generateSuiteActionYAML(suite);
    }

    async function loadProfiles() {
      const data = await requestJSON("/api/profiles");
      state.profiles = data.profiles || [];

      const container = byId("profiles");
      container.innerHTML = "";
      state.profiles.forEach((p) => container.appendChild(createProfileRow(p)));
      applyTargetPreset("ubuntu-lts");
    }

    function selectedProfiles() {
      const include = Array.from(document.querySelectorAll("input[data-kind='include']:checked"))
        .map((x) => x.dataset.id);
      const required = Array.from(document.querySelectorAll("input[data-kind='required']:checked"))
        .map((x) => x.dataset.id)
        .filter((id) => include.includes(id));
      return { include, required };
    }

    function renderSummary(report) {
      const container = byId("summary");
      if (!report || !report.targets) {
        container.replaceChildren();
        return;
      }
      const summaryStatus = String(report.summary && report.summary.status || "unknown");
      const targetCount = report.targets.length;
      const requiredFailed = report.targets.filter((t) => t.required && t.status !== "pass").length;
      const requiredPassed = report.targets.filter((t) => t.required && t.status === "pass").length;
      const optionalFailed = report.targets.filter((t) => !t.required && t.status !== "pass").length;
      const verdictKind = summaryStatus.toLowerCase() === "pass" ? "pass" : "fail";
      let verdictTitle = "PASS: all required targets passed";
      if (summaryStatus.toLowerCase() !== "pass" && requiredFailed > 0) {
        verdictTitle = "FAIL: " + requiredFailed + " required target(s) failed";
      } else if (summaryStatus.toLowerCase() !== "pass") {
        verdictTitle = "CHECK: optional target failure(s) found";
      }
      const verdictMeta = targetCount + " target(s) checked. Required pass/fail: " + requiredPassed + "/" + requiredFailed + ". Optional failures: " + optionalFailed + ".";
      setVerdict(verdictKind, verdictTitle, verdictMeta);

      const headline = document.createElement("div");
      headline.className = summaryStatus.toLowerCase() === "pass" ? "status" : "status error";
      headline.textContent = "Overall: " + summaryStatus + " • targets: " + targetCount + " • required pass/fail: " + requiredPassed + "/" + requiredFailed;

      const counts = document.createElement("div");
      counts.className = "matrix-counts";
      appendMatrixCount(counts, "required passed", requiredPassed, "pass");
      appendMatrixCount(counts, "required failed", requiredFailed, requiredFailed > 0 ? "fail" : "pass");
      appendMatrixCount(counts, "optional failed", optionalFailed, optionalFailed > 0 ? "check" : "pass");
      appendMatrixCount(counts, "targets checked", targetCount);

      const wrap = document.createElement("div");
      wrap.className = "matrix-wrap";
      const table = document.createElement("table");
      const thead = document.createElement("thead");
      const headRow = document.createElement("tr");
      ["Target", "Distro / kernel", "Pass/Fail", "Required", "Reason"].forEach((name) => {
        const th = document.createElement("th");
        th.textContent = name;
        headRow.appendChild(th);
      });
      thead.appendChild(headRow);
      table.appendChild(thead);

      const tbody = document.createElement("tbody");
      const targets = Array.from(report.targets).sort((a, b) => {
        const aRequiredFail = a.required && a.status !== "pass";
        const bRequiredFail = b.required && b.status !== "pass";
        if (aRequiredFail !== bRequiredFail) return aRequiredFail ? -1 : 1;
        const aFail = a.status !== "pass";
        const bFail = b.status !== "pass";
        if (aFail !== bFail) return aFail ? -1 : 1;
        return String(a.profile_id || "").localeCompare(String(b.profile_id || ""));
      });
      targets.forEach((t) => {
        const tr = document.createElement("tr");
        tr.classList.add("matrix-row-" + normalizeStatus(t.status));
        if (t.required && t.status !== "pass") {
          tr.classList.add("matrix-required-fail");
        }
        appendCell(tr, t.profile_id || "-");
        appendCell(tr, formatProfileEnv(t) + " • " + formatHostKernel(t));
        appendStatusCell(tr, t.status || "-");
        appendCell(tr, t.required ? "yes" : "optional");
        appendCell(tr, formatTargetReason(t));
        tbody.appendChild(tr);
      });
      table.appendChild(tbody);
      wrap.appendChild(table);
      container.replaceChildren(headline, counts, wrap);
    }

    document.querySelectorAll("button[data-preset]").forEach((btn) => {
      btn.addEventListener("click", () => applyTargetPreset(btn.dataset.preset));
    });

    byId("selectAll").addEventListener("click", () => {
      selectedPreset = "custom";
      syncTargetPresetButtons();
      document.querySelectorAll("input[data-kind='include']").forEach((x) => {
        x.checked = x.dataset.transportSupported === "true";
      });
      document.querySelectorAll("input[data-kind='required']").forEach((x) => {
        const include = document.querySelector("input[data-kind='include'][data-id='" + x.dataset.id + "']");
        x.disabled = !include || !include.checked;
      });
      updateTargetPresetHint();
    });
    byId("clearAll").addEventListener("click", () => {
      selectedPreset = "custom";
      syncTargetPresetButtons();
      document.querySelectorAll("input[data-kind='include']").forEach((x) => (x.checked = false));
      document.querySelectorAll("input[data-kind='required']").forEach((x) => {
        x.checked = false;
        x.disabled = true;
      });
      updateTargetPresetHint();
    });
    byId("requireSelected").addEventListener("click", () => {
      selectedPreset = "custom";
      syncTargetPresetButtons();
      const picks = selectedProfiles();
      document.querySelectorAll("input[data-kind='required']").forEach((x) => {
        if (picks.include.includes(x.dataset.id) && !x.disabled) {
          x.checked = true;
        }
      });
      updateTargetPresetHint();
    });
    byId("clearRequired").addEventListener("click", () => {
      selectedPreset = "custom";
      syncTargetPresetButtons();
      document.querySelectorAll("input[data-kind='required']").forEach((x) => (x.checked = false));
      updateTargetPresetHint();
    });

    async function runValidationJob(fd) {
      let startResp = null;
      try {
        startResp = await requestJSON("/api/validate/start", {
          method: "POST",
          headers: buildWriteHeaders(),
          body: fd
        });
      } catch (err) {
        if (String(err).includes("HTTP 404")) {
          setStatus("Server does not support async progress endpoint; running direct validation.");
          progressMetaEl.textContent = "Running without live progress (legacy server)";
          return await requestJSON("/api/validate", {
            method: "POST",
            headers: buildWriteHeaders(),
            body: fd
          });
        }
        throw err;
      }
      const jobID = String(startResp.job_id || "").trim();
      if (!jobID) {
        throw new Error("Validation job did not return job_id");
      }

      while (true) {
        const job = await requestJSON("/api/validate/status?job_id=" + encodeURIComponent(jobID));
        renderProgress(job);
        if (job.message) {
          setStatus(job.message);
        } else {
          setStatus("Running validation...");
        }

        if (job.state === "completed") {
          if (!job.result) {
            throw new Error("Validation completed without result payload");
          }
          return job.result;
        }
        if (job.state === "failed") {
          throw new Error(job.error || "Validation failed");
        }
        await sleep(1200);
      }
    }

    byId("runBtn").addEventListener("click", async () => {
      if (bpfInputMode === "suite") {
        updateSuitePreview();
        const count = state.suite.cases.length;
        if (count === 0) {
          setStatus("Paste a suite YAML before generating a CI gate.", true);
          setVerdict("error", "Suite YAML needs cases", "Add cases with name and artifact fields, then use the generated GitHub Action.");
          return;
        }
        setStatus("Suite preview ready. Run this collection through the generated GitHub Action.");
        setVerdict("neutral", "CI suite gate generated", count + " BPF object case(s) ready for self-hosted Linux/KVM execution.");
        return;
      }
      if (runInFlight) {
        setStatus("Validation already running. Please wait.");
        return;
      }
      runInFlight = true;
      runBtnEl.disabled = true;
      try {
        refreshAuthHintFromConfig();
        if (apiConfig && !apiConfig.allow_anonymous_write && !apiConfig.allow_anonymous_validate && !hasWriteCredentials()) {
          throw new Error("Validation is not open on this server. Use the public Results page or run the CLI locally.");
        }
        resetProgress();
        setStatus("Starting validation...");
        setVerdict("running", "Running compatibility gate", "Validating selected targets. Required failures will be shown first.");
        const fd = new FormData();

        fd.append("artifact_name", byId("artifactName").value.trim());
        fd.append("artifact_version", byId("artifactVersion").value.trim());
        fd.append("artifact_variant", byId("artifactVariant").value.trim());
        fd.append("artifact_uri", byId("artifactURI").value.trim());
        fd.append("timeout", byId("timeout").value.trim());
        fd.append("concurrency", byId("concurrency").value.trim());

        if (mode === "artifact") {
          if (byId("artifactFile").files[0]) {
            fd.append("artifact_file", byId("artifactFile").files[0]);
          }
        } else {
          if (byId("sourceFile").files[0]) {
            fd.append("source_file", byId("sourceFile").files[0]);
          }
          if (byId("sourceCode").value.trim()) {
            fd.append("source_code", byId("sourceCode").value);
          }
          if (byId("clangFlags").value.trim()) {
            fd.append("clang_flags", byId("clangFlags").value.trim());
          }
        }

        if (byId("manifestFile").files[0]) {
          fd.append("manifest_file", byId("manifestFile").files[0]);
        }
        if (byId("manifestText").value.trim()) {
          fd.append("manifest_text", byId("manifestText").value);
        }

        const picks = selectedProfiles();
        picks.include.forEach((id) => fd.append("profiles", id));
        picks.required.forEach((id) => fd.append("required_profiles", id));
        if (picks.include.length === 0) {
          throw new Error("Select at least one profile");
        }

        const data = await runValidationJob(fd);

        setStatus("Completed. Exit code " + data.exit_code);
        resultJsonEl.textContent = JSON.stringify(data, null, 2);
        renderSummary(data.report);
        if (evidenceDrilldownEl.open) {
          try {
            await refreshHistory();
          } catch (historyErr) {
            compareJsonEl.textContent = JSON.stringify({ warning: String(historyErr) }, null, 2);
          }
        }
      } catch (err) {
        setStatus(String(err), true);
        setVerdict("error", "Validation could not complete", String(err));
      } finally {
        runInFlight = false;
        runBtnEl.disabled = false;
      }
    });

    function refreshVersionSelectors() {
      const base = byId("baseVersion");
      const head = byId("headVersion");
      base.innerHTML = "";
      head.innerHTML = "";
      for (const rec of state.history) {
        const label = rec.artifact_name + "@" + rec.artifact_version + " (" + rec.summary_status + ")";
        const o1 = document.createElement("option");
        o1.value = rec.artifact_version;
        o1.textContent = label;
        base.appendChild(o1);

        const o2 = document.createElement("option");
        o2.value = rec.artifact_version;
        o2.textContent = label;
        head.appendChild(o2);
      }
      if (head.options.length > 0) {
        head.selectedIndex = 0;
      }
      if (base.options.length > 1) {
        base.selectedIndex = 1;
      }
    }

    async function refreshHistory() {
      const artifactName = byId("historyArtifactName").value.trim();
      const limit = byId("historyLimit").value.trim() || "100";
      const data = await requestJSON("/api/history/artifacts?artifact_name=" + encodeURIComponent(artifactName) + "&limit=" + encodeURIComponent(limit));
      state.history = data.records || [];

      const rows = byId("historyRows");
      rows.innerHTML = "";
      state.history.forEach((rec) => {
        const tr = document.createElement("tr");
        appendCell(tr, rec.artifact_name || "-");
        appendCell(tr, rec.artifact_version || "-");
        appendCell(tr, rec.summary_status || "-");
        appendCell(tr, String(rec.required_passed) + "/" + String(rec.required_failed));
        appendCell(tr, rec.created_at || "-");
        rows.appendChild(tr);
      });
      refreshVersionSelectors();
      refreshRuntimeArtifactSuggestions();
    }

    async function refreshDecisionHistory() {
      const limit = byId("decisionLimit").value.trim() || "100";
      const data = await requestJSON("/api/runtime/decisions?limit=" + encodeURIComponent(limit));
      state.decisions = data.records || [];

      const rows = byId("decisionRows");
      rows.innerHTML = "";
      state.decisions.forEach((rec) => {
        const tr = document.createElement("tr");
        appendCell(tr, rec.decision_id || "-");
        appendCell(tr, rec.operation || "-");
        appendCell(tr, rec.status || "-");
        appendCell(tr, rec.artifact_name || "-");
        appendCell(tr, rec.selected_version || rec.requested_version || "-");
        appendCell(tr, rec.created_at || "-");
        tr.classList.add("clickable-row");
        tr.addEventListener("click", () => {
          decisionJsonEl.textContent = JSON.stringify(rec, null, 2);
        });
        rows.appendChild(tr);
      });
      if (state.decisions.length === 0) {
        decisionJsonEl.textContent = "{}";
      } else {
        decisionJsonEl.textContent = JSON.stringify(state.decisions[0], null, 2);
      }
    }

    async function loadEvidenceIfNeeded() {
      if (!evidenceDrilldownEl.open || evidenceLoaded) {
        return;
      }
      evidenceLoaded = true;
      try {
        await refreshHistory();
      } catch (historyErr) {
        compareJsonEl.textContent = JSON.stringify({ warning: String(historyErr) }, null, 2);
      }
      try {
        await refreshDecisionHistory();
      } catch (decisionErr) {
        decisionJsonEl.textContent = JSON.stringify({ warning: String(decisionErr) }, null, 2);
      }
    }

    byId("refreshHistory").addEventListener("click", async () => {
      try {
        evidenceLoaded = true;
        await refreshHistory();
        setStatus("History refreshed");
      } catch (err) {
        setStatus(String(err), true);
      }
    });

    byId("refreshDecisions").addEventListener("click", async () => {
      try {
        evidenceLoaded = true;
        await refreshDecisionHistory();
        setStatus("Runtime decisions refreshed");
      } catch (err) {
        setStatus(String(err), true);
      }
    });
    evidenceDrilldownEl.addEventListener("toggle", loadEvidenceIfNeeded);

    byId("runCompare").addEventListener("click", async () => {
      try {
        requireWriteCredentials("Compare");
        const artifactName = byId("historyArtifactName").value.trim() || byId("artifactName").value.trim();
        if (!artifactName) {
          throw new Error("Artifact name is required for compare");
        }
        const body = {
          artifact_name: artifactName,
          base_version: byId("baseVersion").value,
          head_version: byId("headVersion").value
        };
        const data = await requestJSON("/api/compare", {
          method: "POST",
          headers: buildWriteHeaders({ "Content-Type": "application/json" }),
          body: JSON.stringify(body)
        });
        compareJsonEl.textContent = JSON.stringify(data, null, 2);
        setStatus("Compare completed");
      } catch (err) {
        setStatus(String(err), true);
      }
    });

    async function runRuntimeProbe() {
      const data = await requestJSON("/api/runtime/probe");
      runtimeJsonEl.textContent = JSON.stringify(data, null, 2);
      const hint = deriveProfileHintFromProbe(data.probe || {});
      if (!byId("runtimeTargetProfile").value.trim() && hint) {
        byId("runtimeTargetProfile").value = hint;
      }
      runtimeCompletedSteps.probe = true;
      setStatus("Runtime probe completed");
      if (runtimeDeliveryActionsAvailable()) {
        setRuntimeMode("select");
        setRuntimeHint("Host probe completed. Continue with Select.");
      } else {
        setRuntimeMode("probe");
        setRuntimeHint("Host probe completed. Selection and fetch are operator-only in this public demo.");
      }
      renderRuntimeSteps();
    }

    async function runRuntimeSelect() {
      requireRuntimeDeliveryAccess("Runtime select");
      const body = runtimeCommonBody();
      if (!body.artifact_name) {
        throw new Error("Runtime artifact name is required");
      }
      const data = await requestJSON("/api/runtime/select", {
        method: "POST",
        headers: buildWriteHeaders({ "Content-Type": "application/json" }),
        body: JSON.stringify(body)
      });
      runtimeJsonEl.textContent = JSON.stringify(data, null, 2);
      if (!byId("runtimeVersion").value.trim() && data.selection && data.selection.selected && data.selection.selected.artifact_version) {
        byId("runtimeVersion").value = data.selection.selected.artifact_version;
      }
      if (!byId("runtimeTargetProfile").value.trim() && data.selection && data.selection.host_profile_hint) {
        byId("runtimeTargetProfile").value = data.selection.host_profile_hint;
      }
      runtimeCompletedSteps.select = true;
      await refreshDecisionHistory();
      setStatus("Runtime select completed");
      setRuntimeMode("fetch");
      setRuntimeHint("Selection completed. Continue with Fetch to retrieve the selected artifact.");
      renderRuntimeSteps();
    }

    async function runRuntimeFetch() {
      requireRuntimeDeliveryAccess("Runtime fetch");
      const body = runtimeCommonBody();
      if (!body.artifact_name) {
        throw new Error("Runtime artifact name is required");
      }
      body.require_verified_history = byId("runtimeRequireVerifiedHistory").checked;
      const data = await requestJSON("/api/runtime/fetch", {
        method: "POST",
        headers: buildWriteHeaders({ "Content-Type": "application/json" }),
        body: JSON.stringify(body)
      });
      runtimeJsonEl.textContent = JSON.stringify(data, null, 2);
      if (!byId("runtimeVersion").value.trim() && data.selection && data.selection.selected && data.selection.selected.artifact_version) {
        byId("runtimeVersion").value = data.selection.selected.artifact_version;
      }
      runtimeCompletedSteps.fetch = true;
      await refreshDecisionHistory();
      setStatus("Runtime fetch completed");
      setRuntimeHint("Fetch completed. Technical output below includes selected version and fetch details.");
      renderRuntimeSteps();
    }

    async function runRuntimeExecute() {
      requireWriteCredentials("Runtime execute");
      const body = runtimeCommonBody();
      if (!body.artifact_name) {
        throw new Error("Runtime artifact name is required");
      }
      body.tenant = byId("runtimeTenant").value.trim();
      body.project = byId("runtimeProject").value.trim();
      body.attach_mode = byId("runtimeAttachMode").value.trim();
      body.probe_features = byId("runtimeProbeFeatures").checked;
      body.require_verified_history = byId("runtimeRequireVerifiedHistory").checked;
      body.allow_host_load = true;

      if (!body.tenant || !body.project) {
        throw new Error("Runtime execute requires tenant and project");
      }
      if (!byId("runtimeApprovalToken").value.trim()) {
        throw new Error("Runtime execute requires Execute Approval Token");
      }
      if (!byId("runtimeRegistryToken").value.trim()) {
        throw new Error("Runtime execute requires Registry Bearer Token");
      }
      if (body.require_verified_history !== true) {
        throw new Error("Runtime execute requires require_verified_history=true");
      }

      const data = await requestJSON("/api/runtime/execute", {
        method: "POST",
        headers: buildRuntimeExecuteHeaders(),
        body: JSON.stringify(body)
      });
      runtimeJsonEl.textContent = JSON.stringify(data, null, 2);
      runtimeCompletedSteps.execute = true;
      await refreshDecisionHistory();
      setStatus("Runtime execute completed");
      setRuntimeHint("Runtime execute completed.");
      renderRuntimeSteps();
    }

    runtimeActionBtn.addEventListener("click", async () => {
      try {
        if (runtimeMode === "probe") {
          await runRuntimeProbe();
          return;
        }
        if (runtimeMode === "select") {
          await runRuntimeSelect();
          return;
        }
        if (runtimeMode === "fetch") {
          await runRuntimeFetch();
          return;
        }
        if (runtimeMode === "execute") {
          await runRuntimeExecute();
          return;
        }
        throw new Error("Unknown runtime mode: " + runtimeMode);
      } catch (err) {
        const message = enhanceRuntimeErrorMessage(runtimeMode, String(err));
        setStatus(message, true);
        setRuntimeHint(message, true);
      }
    });

    (async () => {
      try {
        resetProgress();
        await refreshAPIConfig();
        await loadProfiles();
        updateSuitePreview();
        switchMode("artifact");
        switchBPFInputMode("single");
        setRuntimeMode("probe");
      } catch (err) {
        setStatus(String(err), true);
      }
    })();
  </script>
</body>
</html>`
