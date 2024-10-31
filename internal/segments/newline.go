package segments

import "github.com/thefynx/powerline-go/types"

func segmentNewline(p *types.Powerline) []types.Segment {
	return []types.Segment{{NewLine: true}}
}
