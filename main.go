package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/thefynx/powerline-go/internal/constants"
	"github.com/thefynx/powerline-go/internal/core"
	"github.com/thefynx/powerline-go/internal/helpers"
	"github.com/thefynx/powerline-go/internal/segments"
	"github.com/thefynx/powerline-go/internal/types"
)

var modules = map[string]func(*types.Powerline) []types.Segment{
	"aws":                 segments.AWS,
	"bzr":                 segments.Bzr,
	"cwd":                 segments.Cwd,
	"direnv":              segments.Direnv,
	"docker":              segments.Docker,
	"docker-context":      segments.DockerContext,
	"dotenv":              segments.DotEnv,
	"duration":            segments.Duration,
	"exit":                segments.ExitCode,
	"fossil":              segments.Fossil,
	"gcp":                 segments.GCP,
	"git":                 segments.Git,
	"gitlite":             segments.GitLite,
	"goenv":               segments.Goenv,
	"hg":                  segments.Hg,
	"svn":                 segments.Subversion,
	"host":                segments.Host,
	"jobs":                segments.Jobs,
	"kube":                segments.Kube,
	"load":                segments.Load,
	"newline":             segments.Newline,
	"perlbrew":            segments.Perlbrew,
	"plenv":               segments.PlEnv,
	"perms":               segments.Perms,
	"rbenv":               segments.Rbenv,
	"root":                segments.Root,
	"rvm":                 segments.Rvm,
	"shell-var":           segments.ShellVar,
	"shenv":               segments.ShEnv,
	"ssh":                 segments.SSH,
	"termtitle":           segments.TermTitle,
	"terraform-workspace": segments.TerraformWorkspace,
	"time":                segments.Time,
	"node":                segments.Node,
	"user":                segments.User,
	"venv":                segments.VirtualEnv,
	"vgo":                 segments.VirtualGo,
	"vi-mode":             segments.ViMode,
	"wsl":                 segments.WSL,
	"nix-shell":           segments.NixShell,
}

func comments(lines ...string) string {
	return " " + strings.Join(lines, "\n"+" ")
}

func commentsWithDefaults(lines ...string) string {
	return comments(lines...) + "\n"
}

func main() {
	flag.Parse()

	cfg := core.Defaults
	err := cfg.Load()
	if err != nil {
		println("Error loading config")
		println(err.Error())
	}

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "cwd-mode":
			cfg.CwdMode = *core.Args.CwdMode
		case "cwd-max-depth":
			cfg.CwdMaxDepth = *core.Args.CwdMaxDepth
		case "cwd-max-dir-size":
			cfg.CwdMaxDirSize = *core.Args.CwdMaxDirSize
		case "colorize-hostname":
			cfg.ColorizeHostname = *core.Args.ColorizeHostname
		case "hostname-only-if-ssh":
			cfg.HostnameOnlyIfSSH = *core.Args.HostnameOnlyIfSSH
		case "alternate-ssh-icon":
			cfg.SshAlternateIcon = *core.Args.SshAlternateIcon
		case "east-asian-width":
			cfg.EastAsianWidth = *core.Args.EastAsianWidth
		case "newline":
			cfg.PromptOnNewLine = *core.Args.PromptOnNewLine
		case "static-prompt-indicator":
			cfg.StaticPromptIndicator = *core.Args.StaticPromptIndicator
		case "venv-name-size-limit":
			cfg.VenvNameSizeLimit = *core.Args.VenvNameSizeLimit
		case "jobs":
			cfg.Jobs = *core.Args.Jobs
		case "git-assume-unchanged-size":
			cfg.GitAssumeUnchangedSize = *core.Args.GitAssumeUnchangedSize
		case "git-disable-stats":
			cfg.GitDisableStats = strings.Split(*core.Args.GitDisableStats, ",")
		case "git-mode":
			cfg.GitMode = *core.Args.GitMode
		case "mode":
			cfg.Mode = *core.Args.Mode
		case "theme":
			cfg.Theme = *core.Args.Theme
		case "shell":
			cfg.Shell = *core.Args.Shell
		case "modules":
			cfg.Modules = strings.Split(*core.Args.Modules, ",")
		case "modules-right":
			cfg.ModulesRight = strings.Split(*core.Args.ModulesRight, ",")
		case "priority":
			cfg.Priority = strings.Split(*core.Args.Priority, ",")
		case "max-width":
			cfg.MaxWidthPercentage = *core.Args.MaxWidthPercentage
		case "truncate-segments-width":
			cfg.TruncateSegmentWidth = *core.Args.TruncateSegmentWidth
		case "error":
			cfg.PrevError = *core.Args.PrevError
		case "numeric-exit-codes":
			cfg.NumericExitCodes = *core.Args.NumericExitCodes
		case "ignore-repos":
			cfg.IgnoreRepos = strings.Split(*core.Args.IgnoreRepos, ",")
		case "shorten-gke-names":
			cfg.ShortenGKENames = *core.Args.ShortenGKENames
		case "shorten-eks-names":
			cfg.ShortenEKSNames = *core.Args.ShortenEKSNames
		case "shorten-openshift-names":
			cfg.ShortenOpenshiftNames = *core.Args.ShortenOpenshiftNames
		case "shell-var":
			cfg.ShellVar = *core.Args.ShellVar
		case "shell-var-no-warn-empty":
			cfg.ShellVarNoWarnEmpty = *core.Args.ShellVarNoWarnEmpty
		case "trim-ad-domain":
			cfg.TrimADDomain = *core.Args.TrimADDomain
		case "path-aliases":
			for _, pair := range strings.Split(*core.Args.PathAliases, ",") {
				kv := strings.SplitN(pair, "=", 2)
				cfg.PathAliases[kv[0]] = kv[1]
			}
		case "duration":
			cfg.Duration = *core.Args.Duration
		case "duration-min":
			cfg.DurationMin = *core.Args.DurationMin
		case "duration-low-precision":
			cfg.DurationLowPrecision = *core.Args.DurationLowPrecision
		case "eval":
			cfg.Eval = *core.Args.Eval
		case "condensed":
			cfg.Condensed = *core.Args.Condensed
		case "ignore-warnings":
			cfg.IgnoreWarnings = *core.Args.IgnoreWarnings
		case "time":
			cfg.Time = *core.Args.Time
		case "vi-mode":
			cfg.ViMode = *core.Args.ViMode
		}
	})

	if strings.HasSuffix(cfg.Theme, ".json") {
		file, err := os.ReadFile(cfg.Theme)
		if err == nil {
			theme := cfg.Themes[core.Defaults.Theme]
			err = json.Unmarshal(file, &theme)
			if err == nil {
				cfg.Themes[cfg.Theme] = theme
			} else {
				println("Error reading theme")
				println(err.Error())
			}
		}
	}

	if strings.HasSuffix(cfg.Mode, ".json") {
		file, err := os.ReadFile(cfg.Mode)
		if err == nil {
			symbols := cfg.Modes[core.Defaults.Mode]
			err = json.Unmarshal(file, &symbols)
			if err == nil {
				cfg.Modes[cfg.Mode] = symbols
			} else {
				println("Error reading mode")
				println(err.Error())
			}
		}
	}

	p := core.NewPowerline(cfg, helpers.GetValidCwd(), constants.AlignLeft)
	if p.SupportsRightModules() && p.HasRightModules() && !cfg.Eval {
		panic("Flag '-modules-right' requires '-eval' mode.")
	}

	fmt.Print(p.Draw())
}
