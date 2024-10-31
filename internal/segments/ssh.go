package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentSSH(p *types.Powerline) []types.Segment {
	sshClient, _ := os.LookupEnv("SSH_CLIENT")
	if sshClient == "" {
		return []types.Segment{}
	}
	var networkIcon string
	if p.cfg.SshAlternateIcon {
		networkIcon = p.symbols.NetworkAlternate
	} else {
		networkIcon = p.symbols.Network
	}

	return []types.Segment{{
		Name:       "ssh",
		Content:    networkIcon,
		Foreground: p.theme.SSHFg,
		Background: p.theme.SSHBg,
	}}
}
