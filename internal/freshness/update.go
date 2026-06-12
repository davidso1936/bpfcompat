package freshness

import (
	"github.com/kernel-guard/bpfcompat/pkg/schema"
)

// UpdateFromReport merges the host kernels observed in a matrix report into
// the baselines: matching profiles get their kernel and recorded date
// refreshed, unknown profiles are appended without a crawler mapping (they
// show up as uncovered until someone adds one). Returns the profile ids that
// were updated and added.
func UpdateFromReport(b *Baselines, rep schema.ReportV01) (updated, added []string) {
	recorded := rep.Run.StartedAt
	if len(recorded) > 10 {
		recorded = recorded[:10]
	}

	index := map[string]int{}
	for i, entry := range b.Baselines {
		index[entry.Profile] = i
	}

	for ti := range rep.Targets {
		target := &rep.Targets[ti]
		if target.Host == nil || target.Host.Kernel == "" {
			continue
		}
		if i, ok := index[target.ProfileID]; ok {
			if b.Baselines[i].Kernel == target.Host.Kernel && b.Baselines[i].Recorded != "" {
				continue
			}
			b.Baselines[i].Kernel = target.Host.Kernel
			b.Baselines[i].Recorded = recorded
			updated = append(updated, target.ProfileID)
			continue
		}
		b.Baselines = append(b.Baselines, Baseline{
			Profile:  target.ProfileID,
			Kernel:   target.Host.Kernel,
			Recorded: recorded,
		})
		index[target.ProfileID] = len(b.Baselines) - 1
		added = append(added, target.ProfileID)
	}
	return updated, added
}
