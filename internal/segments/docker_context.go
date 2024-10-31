package segments

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/thefynx/powerline-go/types"
)

type DockerContextConfig struct {
	CurrentContext string `json:"currentContext"`
}

func segmentDockerContext(p *types.Powerline) []types.Segment {
	context := "default"
	home, _ := os.LookupEnv("HOME")
	contextFolder := filepath.Join(home, ".docker", "contexts")
	configFile := filepath.Join(home, ".docker", "config.json")
	contextEnvVar := os.Getenv("DOCKER_CONTEXT")

	if contextEnvVar != "" {
		context = contextEnvVar
	} else {
		stat, err := os.Stat(contextFolder)
		if err == nil && stat.IsDir() {
			dockerConfigFile, err := ioutil.ReadFile(configFile)
			if err == nil {
				var dockerConfig DockerContextConfig
				err = json.Unmarshal(dockerConfigFile, &dockerConfig)
				if err == nil && dockerConfig.CurrentContext != "" {
					context = dockerConfig.CurrentContext
				}
			}
		}
	}

	// Don’t show the default context
	if context == "default" {
		return []types.Segment{}
	}

	return []types.Segment{{
		Name:       "docker-context",
		Content:    "🐳" + context,
		Foreground: p.theme.PlEnvFg,
		Background: p.theme.PlEnvBg,
	}}
}
