package cmd

import (
	"io"
	"os"
)

var (
	OutWriter io.Writer = os.Stdout
	ErrWriter io.Writer = os.Stderr
)
