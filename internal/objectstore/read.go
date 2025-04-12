package objectstore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadObjectRaw(hash string) (header string, body []byte, err error) {
	if len(hash) < 2 {
		return "", nil, errors.New("invalid hash")
	}
	path := filepath.Join(".mygit", "objects", hash[:2], hash[2:])
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read object file: %w", err)
	}

	// Gitオブジェクトは "blob <size>\0<content>" の形式
	parts := strings.SplitN(string(raw), "\u0000", 2)
	if len(parts) != 2 {
		return "", nil, errors.New("invalid object format")
	}
	return parts[0], []byte(parts[1]), nil
}

func ReadObjectContent(hash string) ([]byte, error) {
	_, body, err := ReadObjectRaw(hash)
	return body, err
}

func ReadObjectType(hash string) (string, error) {
	header, _, err := ReadObjectRaw(hash)
	if err != nil {
		return "", err
	}
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid header format")
	}
	return parts[0], nil
}

func ReadObjectSize(hash string) (int, error) {
	header, _, err := ReadObjectRaw(hash)
	if err != nil {
		return 0, err
	}
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 {
		return 0, errors.New("invalid header format")
	}

	size, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid size: %w", err)
	}
	return size, nil
}
