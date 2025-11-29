package main

import (
	"os"

	"github.com/vibe-chung/gh-vibe/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
