//go:build !windows
// +build !windows

package segments

import (
	"github.com/thefynx/powerline-go/types"
	"golang.org/x/sys/unix"
)

func segmentPerms(p *types.Powerline) []types.Segment {
	cwd := p.cwd
	if unix.Access(cwd, unix.W_OK) == nil {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "perms",
		Content:    p.symbols.Lock,
		Foreground: p.theme.ReadonlyFg,
		Background: p.theme.ReadonlyBg,
	}}
}
