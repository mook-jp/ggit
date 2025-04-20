package repository

import (
	"errors"
	"os"
	"path/filepath"
)

func FindRepoRoot(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}

	for {
		mygitPath := filepath.Join(dir, ".mygit")

		info, err := os.Stat(mygitPath)
		if err == nil && info.IsDir() {
			return dir, nil
		}

		// ルートまで到達してしまった場合
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", errors.New(".mygit repository not found")
}
