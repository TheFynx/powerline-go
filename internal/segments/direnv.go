package segments

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/thefynx/powerline-go/types"
)

func segmentDirenv(p *types.Powerline) []types.Segment {
	content := os.Getenv("DIRENV_DIR")
	if content == "" {
		return []types.Segment{}
	}
	if strings.TrimPrefix(content, "-") == p.userInfo.HomeDir {
		content = "~"
	} else {
		content = filepath.Base(content)
	}

	return []types.Segment{{
		Name:       "direnv",
		Content:    content,
		Foreground: p.theme.DotEnvFg,
		Background: p.theme.DotEnvBg,
	}}
}
