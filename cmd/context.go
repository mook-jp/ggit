// Copyright © 2025 mook-jp <mook24.jp@gmail.com>
package cmd

import (
	"io"
	"os"
)

var (
	OutWriter io.Writer = os.Stdout
	ErrWriter io.Writer = os.Stderr
)
