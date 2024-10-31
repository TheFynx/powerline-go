package segments

import (
	"net/url"
	"os"

	"github.com/thefynx/powerline-go/types"
)

func segmentDocker(p *types.Powerline) []types.Segment {
	var docker string
	dockerMachineName, _ := os.LookupEnv("DOCKER_MACHINE_NAME")
	dockerHost, _ := os.LookupEnv("DOCKER_HOST")

	if dockerMachineName != "" {
		docker = dockerMachineName
	} else if dockerHost != " " {
		u, err := url.Parse(dockerHost)
		if err == nil {
			docker = u.Host
		}
	}

	if docker == "" {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "docker",
		Content:    docker,
		Foreground: p.theme.DockerMachineFg,
		Background: p.theme.DockerMachineBg,
	}}
}
