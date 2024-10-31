package segments

import (
	"github.com/thefynx/powerline-go/internal/helpers"
	"github.com/thefynx/powerline-go/types"
)

func segmentViMode(p *types.Powerline) []types.Segment {
	mode := p.cfg.ViMode
	if mode == "" {
		helpers.Warn("'--vi-mode' is not set.")
		return []types.Segment{}
	}

	switch mode {
	case "vicmd":
		return []types.Segment{{
			Name:       "vi-mode",
			Content:    "C",
			Foreground: p.theme.ViModeCommandFg,
			Background: p.theme.ViModeCommandBg,
		}}
	default: // usually "viins" or "main"
		return []types.Segment{{
			Name:       "vi-mode",
			Content:    "I",
			Foreground: p.theme.ViModeInsertFg,
			Background: p.theme.ViModeInsertBg,
		}}
	}
}
