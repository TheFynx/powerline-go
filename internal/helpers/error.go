package helpers

import "github.com/thefynx/powerline-go/internal/core"

func Warn(msg string) {
	if *core.Args.IgnoreWarnings {
		return
	}

	print("[powerline-go]", msg)
}
