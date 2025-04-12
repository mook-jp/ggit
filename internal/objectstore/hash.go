package objectstore

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Hashを作成してファイルを作成する
func HashObject(filename string, write bool, out io.Writer) (string, error) {
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

	if write {
		// ディレクトリとふぃあるに分けてフォルダとファイルを作成する
		// 例: .mygit/objects/ab/cd1234567890abcdef1234567890abcdef1234
		// 2文字目までをディレクトリ名にして、3文字目以降をファイル名にする
		// 一つのディレクトリに何万個もファイルが入ると遅くなるので、2文字目までをディレクトリ名にする
		objectPath := filepath.Join(".mygit", "objects", hashStr[:2], hashStr[2:])
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
		if _, err := f.Write(full); err != nil {
			return "", fmt.Errorf("failed to write object: %w", err)
		}
	}

	return hashStr, nil
}
