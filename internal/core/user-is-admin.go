//go:build !windows
// +build !windows

package core

import (
	"os"
)

func userIsAdmin() bool {
	return os.Getuid() == 0
}
