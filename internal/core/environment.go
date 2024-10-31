package core

import "runtime"

func HomeEnvName() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	return env
}
