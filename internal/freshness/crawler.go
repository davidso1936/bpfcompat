package freshness

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Entry is one kernel release from a kernel-crawler list.json. The upstream
// entries also carry header-package URLs; only the release identity matters
// here, so the rest is dropped at decode time (the x86_64 list is ~90 MB).
type Entry struct {
	KernelRelease string `json:"kernelrelease"`
	Target        string `json:"target"`
}

// Inventory maps a kernel-crawler distro key (ubuntu, rocky, amazonlinux2023,
// ...) to its kernel entries for one architecture.
type Inventory map[string][]Entry

// Newest returns the highest kernel release matching ref, or "" when nothing
// matches.
func (inv Inventory) Newest(ref CrawlerRef) string {
	best := ""
	for _, e := range inv[ref.Distro] {
		if ref.Target != "" && e.Target != ref.Target {
			continue
		}
		if !strings.HasPrefix(e.KernelRelease, ref.ReleasePrefix) {
			continue
		}
		if ref.ReleaseContains != "" && !strings.Contains(e.KernelRelease, ref.ReleaseContains) {
			continue
		}
		if best == "" || CompareKernelReleases(e.KernelRelease, best) > 0 {
			best = e.KernelRelease
		}
	}
	return best
}

// ParseInventory decodes a kernel-crawler list.json stream.
func ParseInventory(r io.Reader) (Inventory, error) {
	var inv Inventory
	if err := json.NewDecoder(r).Decode(&inv); err != nil {
		return nil, fmt.Errorf("decode crawler inventory: %w", err)
	}
	return inv, nil
}

// DefaultCrawlerBase is where falcosecurity/kernel-crawler publishes its
// weekly per-arch inventories as <base>/<arch>/list.json.
const DefaultCrawlerBase = "https://falcosecurity.github.io/kernel-crawler"

// FetchInventory downloads and parses <base>/<arch>/list.json. base must be
// an https URL; use LoadInventoryFile for local copies.
func FetchInventory(ctx context.Context, base, arch string, timeout time.Duration) (Inventory, error) {
	listURL := strings.TrimSuffix(base, "/") + "/" + arch + "/list.json"
	parsed, err := url.Parse(listURL)
	if err != nil {
		return nil, fmt.Errorf("crawler URL: %w", err)
	}
	if parsed.Scheme != "https" {
		return nil, fmt.Errorf("crawler URL must be https, got %q", parsed.Scheme)
	}

	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, parsed.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("crawler request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", parsed, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: HTTP %d", parsed, resp.StatusCode)
	}
	return ParseInventory(resp.Body)
}

// LoadInventoryFile parses a local copy of a kernel-crawler list.json.
func LoadInventoryFile(path string) (Inventory, error) {
	f, err := os.Open(path) // #nosec G304 -- operator-supplied path by design
	if err != nil {
		return nil, fmt.Errorf("open crawler inventory: %w", err)
	}
	defer f.Close()
	return ParseInventory(f)
}
