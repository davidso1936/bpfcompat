# Supply-Chain & Trust Posture

This document records the supply-chain controls that ship as code in this
repository and the maintainer-side settings that must be configured in the
GitHub repository UI/API (they cannot be set from committed files).

## In-repo controls (automated)

| Control | Where | Trigger |
|---|---|---|
| CodeQL static analysis (Go, security-and-quality) | `.github/workflows/codeql.yml` | push/PR to `main`, weekly, manual |
| OpenSSF Scorecard | `.github/workflows/scorecard.yml` | push to `main`, weekly, `branch_protection_rule`, manual |
| Dependency updates (gomod + github-actions) | `.github/dependabot.yml` | weekly (Mondays) |
| Vulnerability scan | `govulncheck` in `.github/workflows/ci.yml` | every PR |
| Lint | `golangci-lint` in `.github/workflows/ci.yml` | every PR |
| SBOM (CycloneDX) | `.github/workflows/release-artifacts.yml` | push to `main`, tags, manual |
| Keyless signing (cosign / Sigstore OIDC) | `.github/workflows/release-artifacts.yml` | tag releases (`v*`) |
| SHA256 checksums | `.github/workflows/release-artifacts.yml` | all builds |

### Verifying a signed release

```bash
cosign verify-blob \
  --certificate SHA256SUMS.crt --signature SHA256SUMS.sig \
  --certificate-identity-regexp 'https://github.com/Kernel-Guard/bpfcompat/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  SHA256SUMS
sha256sum -c SHA256SUMS
```

## Maintainer-side settings (configure once, in the GitHub UI/API)

These are not files in the repo. Track their status here.

- [ ] **Branch protection on `main`**: require PRs, require status checks
      (`ci`, `codeql`), require up-to-date branches, restrict force-pushes.
      Settings → Branches → Add rule, or:
      `gh api -X PUT repos/Kernel-Guard/bpfcompat/branches/main/protection ...`
- [ ] **Secret scanning + push protection**: Settings → Code security → enable
      "Secret scanning" and "Push protection".
- [ ] **Dependabot alerts + security updates**: Settings → Code security →
      enable "Dependabot alerts" and "Dependabot security updates".
- [ ] **Code scanning default setup OFF if CodeQL workflow is used** (avoid
      double analysis): Settings → Code security → Code scanning.
- [ ] **OpenSSF Best Practices badge**: register the project at
      <https://www.bestpractices.dev/>, then add the returned badge:
      `[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/<ID>/badge)](https://www.bestpractices.dev/projects/<ID>)`
      to the README badge row. (Scorecard is automated; this one needs a
      one-time self-assessment.)
- [ ] **Scorecard badge publishing**: the first successful `scorecard.yml` run
      on `main` populates <https://scorecard.dev/viewer/?uri=github.com/Kernel-Guard/bpfcompat>
      and resolves the README badge.

## Notes

- `publish_results: true` in the Scorecard workflow requires the repository to
  be public; it is a no-op signal otherwise.
- CodeQL uses `build-mode: none`, so it does not need the C/libbpf validator
  toolchain (the validator is a separate non-Go component).
- The `Kernel-Guard` GitHub org and the "KernelGuard" site branding should be
  reconciled before a wider public launch; see the project roadmap.
