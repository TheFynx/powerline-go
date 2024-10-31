package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentShEnv(p *types.Powerline) []types.Segment {
	env, _ := os.LookupEnv("SHENV_VERSION")
	if env == "" {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "shenv",
		Content:    env,
		Foreground: p.theme.ShEnvFg,
		Background: p.theme.ShEnvBg,
	}}
}
