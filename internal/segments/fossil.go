package segments

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/thefynx/powerline-go/types"
)

func getFossilStatus() (bool, bool, bool) {
	hasModifiedFiles := false
	hasUntrackedFiles := false
	hasMissingFiles := false

	out, err := exec.Command("fossil", "changes", "--differ").Output()
	if err == nil {
		output := strings.Split(string(out), "\n")
		for _, line := range output {
			if line != "" {
				if strings.HasPrefix(line, "EXTRA") {
					hasUntrackedFiles = true
				} else if strings.HasPrefix(line, "MISSING") {
					hasMissingFiles = true
				} else {
					hasModifiedFiles = true
				}
			}
		}
	}
	return hasModifiedFiles, hasUntrackedFiles, hasMissingFiles
}

func segmentFossil(p *types.Powerline) []types.Segment {
	out, _ := exec.Command("fossil", "branch", "current").Output()
	output := strings.SplitN(string(out), "\n", 2)
	if len(output) > 0 && output[0] != "" {
		branch := output[0]
		hasModifiedFiles, hasUntrackedFiles, hasMissingFiles := getFossilStatus()

		var foreground, background uint8
		var content string
		if hasModifiedFiles || hasUntrackedFiles || hasMissingFiles {
			foreground = p.theme.RepoDirtyFg
			background = p.theme.RepoDirtyBg

			extra := ""

			if hasUntrackedFiles {
				extra += "+"
			}

			if hasMissingFiles {
				extra += "!"
			}

			if hasUntrackedFiles {
				extra += "?"
			}

			content = fmt.Sprintf("%s %s", branch, extra)
		} else {
			foreground = p.theme.RepoCleanFg
			background = p.theme.RepoCleanBg

			content = branch
		}

		return []types.Segment{{
			Name:       "fossil",
			Content:    content,
			Foreground: foreground,
			Background: background,
		}}
	}
	return []types.Segment{}
}