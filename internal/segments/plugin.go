package segments

import (
	"encoding/json"
	"os/exec"

	"github.com/thefynx/powerline-go/types"
)

func segmentPlugin(p *types.Powerline, plugin string) ([]types.Segment, bool) {
	output, err := exec.Command("powerline-go-" + plugin).Output()
	if err != nil {
		return nil, false
	}
	segments := []types.Segment{}
	err = json.Unmarshal(output, &segments)
	if err != nil {
		// The plugin was found but no valid data was returned. Ignore it
		return []types.Segment{}, true
	}
	return segments, true
}
