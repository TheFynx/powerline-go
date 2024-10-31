package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentPlEnv(p *types.Powerline) []types.Segment {
	env, _ := os.LookupEnv("PLENV_VERSION")
	if env == "" {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "plenv",
		Content:    env,
		Foreground: p.theme.PlEnvFg,
		Background: p.theme.PlEnvBg,
	}}
}
