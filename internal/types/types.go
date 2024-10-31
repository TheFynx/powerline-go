package types

import "github.com/thefynx/powerline-go/internal/constants"

type Powerline struct {
	cfg            Config
	cwd            string
	userInfo       User
	userIsAdmin    bool
	hostname       string
	username       string
	theme          Theme
	shell          ShellInfo
	reset          string
	symbols        SymbolTemplate
	priorities     map[string]int
	ignoreRepos    map[string]bool
	Segments       [][]Segment
	curSegment     int
	align          constants.Alignment
	rightPowerline *Powerline
}

// GetTheme returns the theme configuration
func (p *Powerline) GetTheme() Theme {
	return p.theme
}

type prioritizedSegments struct {
	i    int
	segs []Segment
}

// ShellInfo holds information about the shell
type ShellInfo struct {
	RootIndicator         string
	ColorTemplate         string
	EscapedDollar         string
	EscapedBacktick       string
	EscapedBackslash      string
	EvalPromptPrefix      string
	EvalPromptSuffix      string
	EvalPromptRightPrefix string
	EvalPromptRightSuffix string
	VarPrefix             string
	VarSuffix             string
	VarSeparator          string
	VarSeparatorRight     string
	FuncPrefix            string
	FuncSuffix            string
	SegmentSeparator      string
	PromptPrefix          string
	PromptSuffix          string
	RightSegSeparator     string
	RightPromptPrefix     string
	RightPromptSuffix     string
	NoExitCode            bool
	PromptVarSuffix       string
}

// SymbolTemplate holds symbol configuration
type SymbolTemplate struct {
	Separator               string
	SeparatorThin           string
	SeparatorRight          string
	SeparatorRightThin      string
	Lock                    string
	Network                 string
	NetworkAlternate        string
	Separator2              string
	SeparatorThin2          string
	SeparatorRight2         string
	SeparatorRightThin2     string
	Branch                  string
	LockAndNetwork          string
	LockAndNetworkAlternate string
}

// Theme holds color configuration
type Theme struct {
	Reset                  uint8
	UsernameFg             uint8
	UsernameBg             uint8
	UsernameRootBg         uint8
	HostnameFg             uint8
	HostnameBg             uint8
	HomeSpecialDisplay     bool
	HomeFg                 uint8
	HomeBg                 uint8
	PathFg                 uint8
	PathBg                 uint8
	CwdFg                  uint8
	SeparatorFg            uint8
	ReadonlyFg             uint8
	ReadonlyBg             uint8
	SshFg                  uint8
	SshBg                  uint8
	DockerMachineFg        uint8
	DockerMachineBg        uint8
	RepoCleanFg            uint8
	RepoCleanBg            uint8
	RepoDirtyFg            uint8
	RepoDirtyBg            uint8
	JobsFg                 uint8
	JobsBg                 uint8
	CmdPassedFg            uint8
	CmdPassedBg            uint8
	CmdFailedFg            uint8
	CmdFailedBg            uint8
	SvnChangesFg           uint8
	SvnChangesBg           uint8
	GitAheadFg             uint8
	GitAheadBg             uint8
	GitBehindFg            uint8
	GitBehindBg            uint8
	GitStagedFg            uint8
	GitStagedBg            uint8
	GitNotStagedFg         uint8
	GitNotStagedBg         uint8
	GitUntrackedFg         uint8
	GitUntrackedBg         uint8
	GitConflictedFg        uint8
	GitConflictedBg        uint8
	VirtualEnvFg           uint8
	VirtualEnvBg           uint8
	PerlbrewFg             uint8
	PerlbrewBg             uint8
	TimeFg                 uint8
	TimeBg                 uint8
	AWSFg                  uint8
	AWSBg                  uint8
	HostnameColorizedFgMap map[string]uint8
}

// Config holds the powerline configuration
type Config struct {
	ColorizeHostname     bool
	HostnameOnlyIfSSH    bool
	SshAlternateIcon     bool
	EastAsianWidth       bool
	PromptOnNewLine      bool
	Mode                 string
	Theme                string
	Shell                string
	Modules              []string
	ModulesRight         []string
	Priority             []string
	MaxWidthPercentage   int
	TruncateSegmentWidth int
	PrevError            int
	NumericExitCodes     bool
	IgnoreRepos          []string
	ShortenGKENames      bool
	ShortenEKSNames      bool
	ShellVar             string
	ShellVarNoWarnEmpty  bool
	PathAliases          map[string]string
	Duration             string
	DurationMin          string
	Eval                 bool
	Condensed            bool
	IgnoreWarnings       bool
	Themes               map[string]Theme
	Modes                map[string]SymbolTemplate
	Shells               map[string]ShellInfo
}

// User represents a user account
type User struct {
	Uid      string
	Gid      string
	Username string
	Name     string
	HomeDir  string
}
