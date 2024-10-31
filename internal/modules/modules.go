package modules

import (
	segments "github.com/thefynx/powerline-go/internal/segment"
	"github.com/thefynx/powerline-go/internal/types"
)

// ModuleMap contains all available powerline modules
var ModuleMap = map[string]func(*types.Powerline) []types.Segment{
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
