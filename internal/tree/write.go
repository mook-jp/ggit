package tree

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Write(baseDir string) (string, error) {
	// ツリーに格納するリストを作成する
	var entries []byte

	// カレントディレクトリを走査
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || strings.HasPrefix(path, ".mygit") {
			return nil // エラーまたはディレクトリ、.mygitディレクトリは無視
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// blob ヘッダを含めた全体を生成
		header := fmt.Sprintf("blob %d\u0000", len(data))
		full := append([]byte(header), data...)

		// ハッシュを取得
		// スライスに変更しないと append でエラーになる
		sum := sha1.Sum(full)
		hash := sum[:]

		// zlib 圧縮
		var compressed bytes.Buffer
		w := zlib.NewWriter(&compressed)
		if _, err := w.Write(full); err != nil {
			return err
		}
		w.Close()

		// zlib 圧縮したデータを保存
		objPath := filepath.Join(".mygit", "objects", hex.EncodeToString(hash[:2]), hex.EncodeToString(hash[2:]))
		if err := os.MkdirAll(filepath.Dir(objPath), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(objPath, compressed.Bytes(), 0644); err != nil {
			return err
		}

		// treeエントリを作成(mode, name, null, binary, hash)
		// -------------------------------------------------------------------------------------------------------------
		// |name    |value(sample)  |Description                                                                       |
		// |--------|---------------|----------------------------------------------------------------------------------|
		// |mode    |"100644"       |ファイルのアクセスモード（例：644 = 通常ファイル、755 = 実行可能、40000 = ディレクトリ）|
		// |filename|"hello.txt"    |ファイル名（スペースで mode と区切られ、\0 で終わる）                                 |
		// |\u0000  |null byte      |ファイル名とハッシュ値の区切り                                                       |
		// |hash    |20 byte binary |対象ファイルの blob ハッシュ（SHA-1）をバイナリのまま                                 |
		// -------------------------------------------------------------------------------------------------------------
		// 例：100644 hello.txt\0 + 20 byte binary
		mode := "100644"
		filename := filepath.Base(path)
		entry := fmt.Sprintf("%s %s\u0000", mode, filename)
		entryBytes := append([]byte(entry), hash...)
		entries = append(entries, entryBytes...)
		return nil
	})
	if err != nil {
		return "", err
	}

	// treeオブジェクトを作成
	treeHeader := fmt.Sprintf("tree %d\u0000", len(entries))
	treeFull := append([]byte(treeHeader), entries...)

	// ハッシュを計算
	treeHush := sha1.Sum(treeFull)

	// ツリーオブジェクトを圧縮
	var treeCompressed bytes.Buffer
	w := zlib.NewWriter(&treeCompressed)
	if _, err := w.Write(treeFull); err != nil {
		return "", err
	}
	w.Close()

	// ツリーオブジェクトを保存
	treePath := filepath.Join(".mygit", "objects", hex.EncodeToString(treeHush[:2]), hex.EncodeToString(treeHush[2:]))
	if err := os.MkdirAll(filepath.Dir(treePath), 0755); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", treeHush[:]), nil
}
