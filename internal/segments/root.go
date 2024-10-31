package segments

import "github.com/thefynx/powerline-go/types"

func segmentRoot(p *types.Powerline) []types.Segment {
	var foreground, background uint8
	if p.cfg.PrevError == 0 || p.cfg.StaticPromptIndicator {
		foreground = p.theme.CmdPassedFg
		background = p.theme.CmdPassedBg
	} else {
		foreground = p.theme.CmdFailedFg
		background = p.theme.CmdFailedBg
	}

	return []types.Segment{{
		Name:       "root",
		Content:    p.shell.RootIndicator,
		Foreground: foreground,
		Background: background,
	}}
}
