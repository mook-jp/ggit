package repository

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

// CopyDir recursively copies src to dst.
func CopyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dst, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// ファイルをコピー
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

func TestFindRepoRoot(t *testing.T) {
	tmpDir := t.TempDir()

	// .mygit を作成するパス
	repoRoot := filepath.Join(tmpDir, "project")
	workDir := filepath.Join(repoRoot, "subdir", "subsub")

	if err := os.MkdirAll(filepath.Join(workDir), 0755); err != nil {
		t.Fatalf("failed to create nested workdir: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(filepath.Join(repoRoot), ".mygit"), 0755); err != nil {
		t.Fatalf("failed to create .mygit: %v", err)
	}

	// カレントディレクトリを入れ子に深い場所に移動
	if err := os.Chdir(workDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	found, err := FindRepoRoot(".")
	if err != nil {
		t.Fatalf("expected to find .mygit repo, but got error: %v", err)
	}

	if found != repoRoot {
		t.Errorf("expected repo root to be %s, got %s", repoRoot, found)
	}
}

func TestFindRepoRoot_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	_, err := FindRepoRoot(".")
	if err == nil {
		t.Fatal("expected error when .mygit not found, got nil")
	}
}
