// Copyright © 2025 mook-jp <mook24.jp@gmail.com>
package initrepo

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Options struct {
	BaseDir       string
	InitialBranch string
	Stdout        io.Writer
	Stderr        io.Writer
}

var (
	ErrRepoAlreadyExists = errors.New(".my git already exists")
)

func InitRepo(opts Options) error {
	gitDir := filepath.Join(opts.BaseDir, ".mygit")

	if _, err := os.Stat(gitDir); err == nil {
		fmt.Fprintln(opts.Stderr, "Warning: .mygit already exists")
		return ErrRepoAlreadyExists
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking .mygit: %w", err)
	}

	dirs := []string{
		filepath.Join(gitDir, "objects"),
		filepath.Join(gitDir, "resfs"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	headPath := filepath.Join(gitDir, "HEAD")
	headContent := fmt.Sprintf("ref: refs/heads/%s\n", opts.InitialBranch)
	if err := os.WriteFile(headPath, []byte(headContent), 0644); err != nil {
		return fmt.Errorf("failed to write HEAD file: %w", err)
	}

	fmt.Fprintln(opts.Stdout, "Initialized empty MyGit repository in .mygit")
	return nil
}
