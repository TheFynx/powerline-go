package segments

import (
	"strings"
	"time"

	"github.com/thefynx/powerline-go/types"
)

func segmentTime(p *types.Powerline) []types.Segment {
	return []types.Segment{{
		Name:       "time",
		Content:    time.Now().Format(strings.TrimSpace(p.cfg.Time)),
		Foreground: p.theme.TimeFg,
		Background: p.theme.TimeBg,
	}}
}
