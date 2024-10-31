package segments

import (
	"os"
	"path"

	"github.com/thefynx/powerline-go/types"
)

func segmentPerlbrew(p *types.Powerline) []types.Segment {
	env, _ := os.LookupEnv("PERLBREW_PERL")
	if env == "" {
		return []types.Segment{}
	}

	envName := path.Base(env)
	return []types.Segment{{
		Name:       "perlbrew",
		Content:    envName,
		Foreground: p.theme.PerlbrewFg,
		Background: p.theme.PerlbrewBg,
	}}
}
