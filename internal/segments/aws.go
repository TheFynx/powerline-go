package segments

import (
	"os"

	"github.com/thefynx/powerline-go/internal/types"
)

func AWS(p *types.Powerline) []types.Segment {
	profile := os.Getenv("AWS_PROFILE")
	region := os.Getenv("AWS_DEFAULT_REGION")
	if profile == "" {
		profile = os.Getenv("AWS_VAULT")
		if profile == "" {
			return []types.Segment{}
		}
	}
	var r string
	if region != "" {
		r = " (" + region + ")"
	}
	return []types.Segment{{
		Name:       "aws",
		Content:    profile + r,
		Foreground: p.GetTheme().AWSFg,
		Background: p.GetTheme().AWSBg,
	}}
}
