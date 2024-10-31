// +build windows

package segments

import (
	"os"

	pwl "github.com/thefynx/powerline-go"
)

func segmentPerms(p *types.Powerline) []types.Segment {
	cwd := p.cwd
	const W_USR = 0002
	// Check user's permissions on directory in a portable but probably slower way
	fileInfo, _ := os.Stat(cwd)
	if fileInfo.Mode()&W_USR == W_USR {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "perms",
		Content:    p.symbols.Lock,
		Foreground: p.theme.ReadonlyFg,
		Background: p.theme.ReadonlyBg,
	}}
}
