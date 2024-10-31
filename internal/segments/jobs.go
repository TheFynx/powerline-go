package segments

import (
	"strconv"

	"github.com/thefynx/powerline-go/types"
)

func segmentJobs(p *types.Powerline) []types.Segment {
	if p.cfg.Jobs <= 0 {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "jobs",
		Content:    strconv.Itoa(p.cfg.Jobs),
		Foreground: p.theme.JobsFg,
		Background: p.theme.JobsBg,
	}}
}
