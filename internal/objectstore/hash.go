// Copyright © 2025 mook-jp <mook24.jp@gmail.com>
package objectstore

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mook-jp/ggit/internal/repository"
)

// Hashを作成してファイルを作成する
func HashObject(filename string, writeFlag bool, out io.Writer) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// blob <size>\0<content>の形式でハッシュ用のデータを作成
	header := fmt.Sprintf("blob %d\u0000", len(data))
	full := append([]byte(header), data...)

	// ファイルのハッシュ値を計算する
	hash := sha1.Sum(full)
	hashStr := fmt.Sprintf("%x", hash)

	if writeFlag {
		// オブジェクトルートを取得する
		repoRoot, err := repository.FindRepoRoot(".")
		if err != nil {
			return "", err
		}

		// データを zlib で圧縮する
		var full_zlib bytes.Buffer
		w := zlib.NewWriter(&full_zlib)
		_, err = w.Write(full)
		if err != nil {
			return "", fmt.Errorf("failed to write zlib: %w", err)
		}
		w.Close()

		// ディレクトリとファイルに分けてフォルダとファイルを作成する
		// 例: .mygit/objects/ab/cd1234567890abcdef1234567890abcdef1234
		// 2文字目までをディレクトリ名にして、3文字目以降をファイル名にする
		// 一つのディレクトリに何万個もファイルが入ると遅くなるので、2文字目までをディレクトリ名にする
		objectPath := filepath.Join(repoRoot, ".mygit", "objects", hashStr[:2], hashStr[2:])
		if err := os.MkdirAll(filepath.Dir(objectPath), 0755); err != nil {
			return "", fmt.Errorf("failed to create object dir: %w", err)
		}

		// ファイルを作成する
		f, err := os.Create(objectPath)
		if err != nil {
			return "", fmt.Errorf("failed to create object file: %w", err)
		}
		defer f.Close()

		// ファイルに書き込む
		if _, err := f.Write(full_zlib.Bytes()); err != nil {
			return "", fmt.Errorf("failed to write object: %w", err)
		}
	}

	return hashStr, nil
}
