package segments

import (
	"net/url"
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentWSL(p *types.Powerline) []types.Segment {
	var WSL string
	WSLMachineName, _ := os.LookupEnv("WSL_DISTRO_NAME")
	WSLHost, _ := os.LookupEnv("NAME")

	if WSLMachineName != "" {
		WSL = WSLMachineName
	} else if WSLHost != " " {
		u, err := url.Parse(WSLHost)
		if err == nil {
			WSL = u.Host
		}
	}

	if WSL != "" {
		return []types.Segment{{
			Name:       "WSL",
			Content:    WSL,
			Foreground: p.theme.WSLMachineFg,
			Background: p.theme.WSLMachineBg,
		}}
	}
	return []types.Segment{}
}
