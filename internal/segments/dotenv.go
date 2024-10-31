package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentDotEnv(p *types.Powerline) []types.Segment {
	files := []string{".env", ".envrc"}
	dotEnv := false
	for _, file := range files {
		stat, err := os.Stat(file)
		if err == nil && !stat.IsDir() {
			dotEnv = true
			break
		}
	}
	if !dotEnv {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "dotenv",
		Content:    "\u2235",
		Foreground: p.theme.DotEnvFg,
		Background: p.theme.DotEnvBg,
	}}
}
