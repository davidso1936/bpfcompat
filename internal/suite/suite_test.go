package suite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLoadBytesRejectsUnknownFields(t *testing.T) {
	_, err := LoadBytes([]byte(`
name: demo
unknown: true
cases:
  - name: one
    artifact: a.bpf.o
    matrix: matrix.yaml
`))
	if err == nil {
		t.Fatalf("expected unknown field error")
	}
}

func TestLoadBytesValidatesCases(t *testing.T) {
	_, err := LoadBytes([]byte(`
name: demo
cases:
  - name: one
    artifact: a.bpf.o
`))
	if err == nil || !strings.Contains(err.Error(), "matrix is required") {
		t.Fatalf("expected missing matrix error, got %v", err)
	}
}

func TestLoadBytesValidatesValidationModeAndBehaviorTest(t *testing.T) {
	_, err := LoadBytes([]byte(`
name: demo
defaults:
  matrix: matrix.yaml
  validation_mode: mystery
cases:
  - name: one
    artifact: a.bpf.o
`))
	if err == nil || !strings.Contains(err.Error(), "defaults.validation_mode") {
		t.Fatalf("expected invalid default validation mode error, got %v", err)
	}

	_, err = LoadBytes([]byte(`
name: demo
defaults:
  matrix: matrix.yaml
cases:
  - name: one
    artifact: a.bpf.o
    test:
      mode: smoke
      command: /bin/true
`))
	if err == nil || !strings.Contains(err.Error(), "test.mode") {
		t.Fatalf("expected invalid behavior mode error, got %v", err)
	}

	spec, err := LoadBytes([]byte(`
name: demo
defaults:
  matrix: matrix.yaml
  validation_mode: load_only
cases:
  - name: one
    artifact: a.bpf.o
    validation_mode: load_attach
  - name: behavior
    artifact: b.bpf.o
    test:
      mode: behavior
      command: /bin/true
      timeout: 10s
      expect:
        exit_code: 0
`))
	if err != nil {
		t.Fatalf("expected valid suite, got %v", err)
	}
	if spec.Defaults.ValidationMode != "load_only" || spec.Cases[0].ValidationMode != "load_attach" || spec.Cases[1].Test == nil {
		t.Fatalf("unexpected parsed suite: %+v", spec)
	}
}

func TestBuildRunnerConfigResolvesSuiteRelativePaths(t *testing.T) {
	spec := Spec{
		Name: "demo",
		Defaults: Defaults{
			Matrix:    "../matrices/dev-one.yaml",
			WorkDir:   "../.bpfcompat",
			ReportDir: "../reports/suites/demo",
			Timeout:   "8m",
		},
		Cases: []Case{
			{
				Name:     "functional",
				Artifact: "../examples/functional-execve/functional_execve.bpf.o",
				Manifest: "../examples/functional-execve/manifest-dev-one.yaml",
			},
		},
	}

	cfg, summary, err := buildRunnerConfig("/repo/suites", spec, spec.Cases[0], RunOptions{Concurrency: 1})
	if err != nil {
		t.Fatalf("build config: %v", err)
	}

	assertEqual(t, cfg.ArtifactPath, "/repo/examples/functional-execve/functional_execve.bpf.o")
	assertEqual(t, cfg.ManifestPath, "/repo/examples/functional-execve/manifest-dev-one.yaml")
	assertEqual(t, cfg.MatrixPath, "/repo/matrices/dev-one.yaml")
	assertEqual(t, cfg.OutPath, "/repo/reports/suites/demo/functional.json")
	assertEqual(t, cfg.MarkdownPath, "/repo/reports/suites/demo/functional.md")
	assertEqual(t, cfg.WorkDir, "/repo/.bpfcompat")
	if cfg.Timeout != 8*time.Minute {
		t.Fatalf("unexpected timeout: %s", cfg.Timeout)
	}
	if cfg.Concurrency != 1 {
		t.Fatalf("unexpected concurrency: %d", cfg.Concurrency)
	}
	assertEqual(t, summary.ReportJSONPath, "/repo/reports/suites/demo/functional.json")
}

func TestBuildRunnerConfigSupportsBehaviorTestManifest(t *testing.T) {
	reportDir := t.TempDir()
	baseDir := filepath.Join(reportDir, "suites")
	spec := Spec{
		Name: "demo",
		Defaults: Defaults{
			Matrix:         "../matrices/dev-one.yaml",
			ReportDir:      reportDir,
			ValidationMode: "load_only",
		},
		Cases: []Case{
			{
				Name:           "behavior",
				Artifact:       "../examples/functional-execve/functional_execve.bpf.o",
				ArtifactName:   "functional_execve",
				ValidationMode: "load_attach",
				Test: &Test{
					Mode:    "behavior",
					Command: "/bin/true",
					Expect: Expect{
						ExitCode: intPtr(0),
					},
				},
			},
		},
	}

	cfg, summary, err := buildRunnerConfig(baseDir, spec, spec.Cases[0], RunOptions{Concurrency: 1})
	if err != nil {
		t.Fatalf("build config: %v", err)
	}

	assertEqual(t, cfg.ValidationMode, "behavior")
	assertEqual(t, summary.ValidationMode, "behavior")
	if !strings.HasSuffix(cfg.ManifestPath, "behavior.behavior-manifest.yaml") {
		t.Fatalf("unexpected generated manifest path: %s", cfg.ManifestPath)
	}
	raw, err := os.ReadFile(cfg.ManifestPath)
	if err != nil {
		t.Fatalf("read generated manifest: %v", err)
	}
	for _, want := range []string{"functional_tests:", "behavior-behavior", "command: /bin/true"} {
		if !strings.Contains(string(raw), want) {
			t.Fatalf("generated manifest missing %q:\n%s", want, raw)
		}
	}
}

func TestRenderMarkdownIncludesCases(t *testing.T) {
	md := RenderMarkdown(Summary{
		Name:     "demo",
		Status:   "fail",
		ExitCode: 2,
		Cases: []CaseSummary{
			{
				Name:               "one",
				Status:             "pass",
				RequiredPassed:     1,
				TotalProfiles:      1,
				ReportMarkdownPath: "one.md",
				ValidationMode:     "load_only",
				Targets: []CaseTargetSummary{
					{ProfileID: "ubuntu-24.04-6.8", Status: "pass", Required: true},
				},
			},
			{
				Name:               "two",
				Status:             "fail",
				RequiredFailed:     1,
				TotalProfiles:      1,
				ReportMarkdownPath: "two.md",
				ValidationMode:     "behavior",
				BehaviorStatus:     "fail",
				BehaviorPassed:     1,
				BehaviorFailed:     1,
				Targets: []CaseTargetSummary{
					{ProfileID: "ubuntu-24.04-6.8", Status: "fail", ClassificationCode: "MISSING_BTF"},
				},
			},
		},
	})
	for _, want := range []string{"# bpfcompat Compatibility Suite", "`demo`", "`one`", "`two`", "0/1", "`behavior`", "`fail 1/1`", "## Collection Matrix", "`ubuntu-24.04-6.8`", "`pass required`", "`fail MISSING_BTF`"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q:\n%s", want, md)
		}
	}
}

func intPtr(v int) *int {
	return &v
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("unexpected value: got=%q want=%q", got, want)
	}
}
