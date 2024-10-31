package segments

import (
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentNixShell(p *types.Powerline) []types.Segment {
	var nixShell string
	nixShell, _ = os.LookupEnv("IN_NIX_SHELL")
	if nixShell == "" {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "nix-shell",
		Content:    "\uf313",
		Foreground: p.theme.NixShellFg,
		Background: p.theme.NixShellBg,
	}}
}
