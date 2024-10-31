package helpers

import (
	"os"
	"strings"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetValidCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		var exists bool
		cwd, exists = os.LookupEnv("PWD")
		if !exists {
			Warn("Your current directory is invalid.")
			print("> ")
			os.Exit(1)
		}
	}

	parts := strings.Split(cwd, string(os.PathSeparator))
	up := cwd

	for len(parts) > 0 && !PathExists(up) {
		parts = parts[:len(parts)-1]
		up = strings.Join(parts, string(os.PathSeparator))
	}
	if cwd != up {
		Warn("Your current directory is invalid. Lowest valid directory: " + up)
	}
	return cwd
}
