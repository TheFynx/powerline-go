package segments

import (
	"fmt"
	"strings"

	"github.com/thefynx/powerline-go/types"
)

func segmentGitLite(p *types.Powerline) []types.Segment {
	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand("git", "--no-optional-locks", "rev-parse", "--show-toplevel")
		if err != nil {
			return []types.Segment{}
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return []types.Segment{}
		}
	}

	out, err := runGitCommand("git", "--no-optional-locks", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return []types.Segment{}
	}

	status := strings.TrimSpace(out)
	var branch string

	if status == "HEAD" {
		branch = getGitDetachedBranch(p)
	} else {
		branch = status
	}

	if p.cfg.GitMode != "compact" && len(p.symbols.RepoBranch) > 0 {
		branch = fmt.Sprintf("%s %s", p.symbols.RepoBranch, branch)
	}

	return []types.Segment{{
		Name:       "git-branch",
		Content:    branch,
		Foreground: p.theme.RepoCleanFg,
		Background: p.theme.RepoCleanBg,
	}}
}
