package segments

import (
	"io/ioutil"
	"os"

	"github.com/thefynx/powerline-go/types"
)

const wsFile = "./.terraform/environment"

func segmentTerraformWorkspace(p *types.Powerline) []types.Segment {
	stat, err := os.Stat(wsFile)
	if err != nil {
		return []types.Segment{}
	}
	if stat.IsDir() {
		return []types.Segment{}
	}
	workspace, err := ioutil.ReadFile(wsFile)
	if err != nil {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "terraform-workspace",
		Content:    string(workspace),
		Foreground: p.theme.TFWsFg,
		Background: p.theme.TFWsBg,
	}}
}
