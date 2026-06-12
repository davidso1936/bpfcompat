// Package freshness compares the kernel releases last validated per VM
// profile against the per-distro kernel inventory published weekly by
// falcosecurity/kernel-crawler. The crawler indexes header packages, not
// bootable images, so it acts as a freshness oracle for the image pipeline:
// when a distro ships a kernel newer than the one a profile's image last
// booted, the matrix evidence for that profile is behind and the image (and
// validation run) should be refreshed.
package freshness

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Baselines is the committed record of the last-validated guest kernel per
// profile plus how to find that profile's kernel series in kernel-crawler
// output.
type Baselines struct {
	Baselines []Baseline `yaml:"baselines"`
}

// Baseline records one profile's last observed kernel and its crawler mapping.
type Baseline struct {
	Profile  string      `yaml:"profile"`
	Kernel   string      `yaml:"kernel,omitempty"`
	Recorded string      `yaml:"recorded,omitempty"`
	Note     string      `yaml:"note,omitempty"`
	Crawler  *CrawlerRef `yaml:"crawler,omitempty"`
}

// CrawlerRef selects the kernel-crawler entries that correspond to one
// profile's kernel series: the top-level distro key, an optional target
// flavor (for example ubuntu-generic vs ubuntu-kvm), a release prefix that
// pins the series (5.15.0-), and an optional substring for distros that mix
// several major releases under one key (el9 vs el10 under rocky).
type CrawlerRef struct {
	Distro          string `yaml:"distro"`
	Target          string `yaml:"target,omitempty"`
	ReleasePrefix   string `yaml:"release_prefix"`
	ReleaseContains string `yaml:"release_contains,omitempty"`
	Arch            string `yaml:"arch,omitempty"`
}

// DefaultArch is the crawler architecture used when a baseline does not set
// one explicitly.
const DefaultArch = "x86_64"

// Result statuses. Stale is the actionable one; the rest are informational.
const (
	StatusFresh     = "fresh"      // baseline >= newest shipping kernel in the series
	StatusStale     = "stale"      // distro ships a newer kernel than last validated
	StatusUncovered = "uncovered"  // no crawler mapping for this profile
	StatusNoEntries = "no-entries" // mapping matched nothing (typically an EOL series)
	StatusNoKernel  = "no-kernel"  // baseline has a mapping but no recorded kernel yet
)

// Result is the freshness verdict for one baseline entry.
type Result struct {
	Profile  string `json:"profile"`
	Baseline string `json:"baseline,omitempty"`
	Newest   string `json:"newest,omitempty"`
	Status   string `json:"status"`
	Reason   string `json:"reason,omitempty"`
}

// Evaluate compares every baseline against the crawler inventory for its
// architecture. fetch is called at most once per distinct arch.
func Evaluate(b Baselines, fetch func(arch string) (Inventory, error)) ([]Result, error) {
	inventories := map[string]Inventory{}
	results := make([]Result, 0, len(b.Baselines))

	for _, entry := range b.Baselines {
		res := Result{Profile: entry.Profile, Baseline: entry.Kernel}
		switch {
		case entry.Crawler == nil:
			res.Status = StatusUncovered
			res.Reason = entry.Note
			if res.Reason == "" {
				res.Reason = "no kernel-crawler mapping"
			}
		case entry.Kernel == "":
			res.Status = StatusNoKernel
			res.Reason = "no validated kernel recorded yet"
		default:
			arch := entry.Crawler.Arch
			if arch == "" {
				arch = DefaultArch
			}
			inv, ok := inventories[arch]
			if !ok {
				fetched, err := fetch(arch)
				if err != nil {
					return nil, fmt.Errorf("fetch crawler inventory for %s: %w", arch, err)
				}
				inv = fetched
				inventories[arch] = inv
			}
			newest := inv.Newest(*entry.Crawler)
			if newest == "" {
				res.Status = StatusNoEntries
				res.Reason = "no crawler entries match the series (EOL or unpublished)"
				break
			}
			res.Newest = newest
			if CompareKernelReleases(newest, entry.Kernel) > 0 {
				res.Status = StatusStale
				res.Reason = "distro ships a newer kernel than the last validated image"
			} else {
				res.Status = StatusFresh
			}
		}
		results = append(results, res)
	}
	return results, nil
}

// StaleCount returns how many results are actionable.
func StaleCount(results []Result) int {
	n := 0
	for _, r := range results {
		if r.Status == StatusStale {
			n++
		}
	}
	return n
}

var digitRuns = regexp.MustCompile(`\d+`)

// CompareKernelReleases orders two kernel release strings by their numeric
// components (5.15.0-184-generic > 5.15.0-92-generic). Non-numeric text only
// separates the numbers; it never participates in the ordering. Returns
// -1, 0, or 1.
func CompareKernelReleases(a, b string) int {
	av := digitRuns.FindAllString(a, -1)
	bv := digitRuns.FindAllString(b, -1)
	for i := 0; i < len(av) && i < len(bv); i++ {
		ai, _ := strconv.Atoi(av[i])
		bi, _ := strconv.Atoi(bv[i])
		if ai != bi {
			if ai > bi {
				return 1
			}
			return -1
		}
	}
	switch {
	case len(av) > len(bv):
		return 1
	case len(av) < len(bv):
		return -1
	default:
		return 0
	}
}

// LoadBaselines parses a committed baselines file with strict field checking.
func LoadBaselines(data []byte) (Baselines, error) {
	var b Baselines
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&b); err != nil {
		return Baselines{}, fmt.Errorf("parse baselines: %w", err)
	}
	seen := map[string]bool{}
	for i, entry := range b.Baselines {
		if entry.Profile == "" {
			return Baselines{}, fmt.Errorf("baseline %d: profile is required", i)
		}
		if seen[entry.Profile] {
			return Baselines{}, fmt.Errorf("baseline %d: duplicate profile %q", i, entry.Profile)
		}
		seen[entry.Profile] = true
		if entry.Crawler != nil {
			if entry.Crawler.Distro == "" {
				return Baselines{}, fmt.Errorf("baseline %s: crawler.distro is required", entry.Profile)
			}
			if entry.Crawler.ReleasePrefix == "" {
				return Baselines{}, fmt.Errorf("baseline %s: crawler.release_prefix is required", entry.Profile)
			}
		}
	}
	return b, nil
}

const baselinesHeader = `# Last-validated guest kernel per VM profile, compared weekly against the
# kernel inventory published by falcosecurity/kernel-crawler
# (https://falcosecurity.github.io/kernel-crawler/). Refresh entries from a
# matrix run with:
#   bpfcompat kernel-freshness --update-from-report <report.json>
# Entries without a crawler mapping are reported as uncovered, not stale.
`

// MarshalBaselines renders the baselines file deterministically, sorted by
// profile id, with the explanatory header.
func MarshalBaselines(b Baselines) ([]byte, error) {
	sorted := make([]Baseline, len(b.Baselines))
	copy(sorted, b.Baselines)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Profile < sorted[j].Profile })

	var buf bytes.Buffer
	buf.WriteString(baselinesHeader)
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(Baselines{Baselines: sorted}); err != nil {
		return nil, fmt.Errorf("encode baselines: %w", err)
	}
	if err := enc.Close(); err != nil {
		return nil, fmt.Errorf("encode baselines: %w", err)
	}
	return buf.Bytes(), nil
}

// Markdown renders the freshness results as a report suitable for a GitHub
// job summary.
func Markdown(results []Result) string {
	var b strings.Builder
	b.WriteString("# Kernel freshness vs kernel-crawler\n\n")
	stale := StaleCount(results)
	if stale == 0 {
		b.WriteString("All covered profiles are validated against the newest kernel their distro ships.\n\n")
	} else {
		fmt.Fprintf(&b, "**%d profile(s) are behind what their distro currently ships.** Refresh the cached image and re-run the matrix to update evidence.\n\n", stale)
	}
	b.WriteString("| Profile | Last validated | Newest shipping | Status | Notes |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, r := range results {
		status := r.Status
		if r.Status == StatusStale {
			status = "**stale**"
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n",
			r.Profile, orDash(r.Baseline), orDash(r.Newest), status, orDash(r.Reason))
	}
	b.WriteString("\nInventory: [falcosecurity/kernel-crawler](https://falcosecurity.github.io/kernel-crawler/) (header-package index used as a freshness oracle; images stay vendor cloud images).\n")
	return b.String()
}

func orDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
