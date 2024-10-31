package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentVirtualGo(p *types.Powerline) []types.Segment {
	env, _ := os.LookupEnv("VIRTUALGO")
	if env == "" {
		return []types.Segment{}
	}

	return []types.Segment{{
		Name:       "vgo",
		Content:    env,
		Foreground: p.theme.VirtualGoFg,
		Background: p.theme.VirtualGoBg,
	}}
}
