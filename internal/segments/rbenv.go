package segments

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/thefynx/powerline-go/types"
)

const rubyVersionFileSuffix = "/.ruby-version"
const globalVersionFileSuffix = "/.rbenv/version"

func runRbenvCommand(cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)
	out, err := command.Output()
	return string(out), err
}

// check RBENV_VERSION variable
func checkEnvForRbenvVersion() (string, error) {
	rbenvVersion := os.Getenv("RBENV_VERSION")
	if len(rbenvVersion) <= 0 {
		return "", errors.New("Not found in RBENV_VERSION")
	}
	return rbenvVersion, nil
}

// check existence of .ruby_version in tree until root path
func checkForRubyVersionFileInTree() (string, error) {
	var (
		workingDirectory string
		err              error
	)

	workingDirectory, err = os.Getwd()
	if err == nil {
		for workingDirectory != "/" {
			rubyVersion, rubyVersionErr := ioutil.ReadFile(workingDirectory + rubyVersionFileSuffix)
			if rubyVersionErr == nil {
				return strings.TrimSpace(string(rubyVersion)), nil
			}

			workingDirectory = filepath.Dir(workingDirectory)
		}
	}

	return "", errors.New("No .ruby_version file found in tree")
}

// check for global version
func checkForGlobalVersion() (string, error) {
	homeDir, _ := os.UserHomeDir()
	globalRubyVersion, err := ioutil.ReadFile(homeDir + globalVersionFileSuffix)
	if err != nil {
		return "", errors.New("No global version file found in tree")
	}
	return strings.TrimSpace(string(globalRubyVersion)), nil
}

// retrieve rbenv version output
func checkForRbenvOutput() (string, error) {
	// spawn rbenv and print out version
	out, err := runRbenvCommand("rbenv", "version")
	if err != nil {
		return "", errors.New("Not found in rbenv output")
	}
	items := strings.Split(out, " ")
	if len(items) <= 0 {
		return "", errors.New("Not found in rbenv output")
	}

	return items[0], nil
}

func segmentRbenv(p *types.Powerline) []types.Segment {
	var (
		segment string
		err     error
	)

	segment, err = checkEnvForRbenvVersion()
	if err != nil {
		segment, err = checkForRubyVersionFileInTree()
	}
	if err != nil {
		segment, err = checkForGlobalVersion()
	}
	if err != nil {
		segment, err = checkForRbenvOutput()
	}
	if err != nil {
		return []types.Segment{}
	}
	return []types.Segment{{
		Name:       "rbenv",
		Content:    segment,
		Foreground: p.theme.TimeFg,
		Background: p.theme.TimeBg,
	}}
}