package segments

import (
	"os"

	"github.com/thefynx/powerline-go/internal/helpers"
	"github.com/thefynx/powerline-go/types"
)

func segmentShellVar(p *types.Powerline) []types.Segment {
	shellVarName := p.cfg.ShellVar
	varContent, varExists := os.LookupEnv(shellVarName)

	if !varExists {
		if shellVarName != "" {
			helpers.Warn("Shell variable " + shellVarName + " does not exist.")
		}
		return []types.Segment{}
	}

	if varContent == "" {
		if !p.cfg.ShellVarNoWarnEmpty {
			helpers.Warn("Shell variable " + shellVarName + " is empty.")
		}
		return []types.Segment{}
	}

	return []types.Segment{{
		Name:       "shell-var",
		Content:    varContent,
		Foreground: p.theme.ShellVarFg,
		Background: p.theme.ShellVarBg,
	}}
}
