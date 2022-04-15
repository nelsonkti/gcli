package xfile

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name   string
	Data   []byte
	Format string
}

// 创建文件夹
func MkdirFolder(path string, modePerm uint32) (string, error) {

	folder, err := filepath.Abs(path)
	if err != nil {
		return path, err
	}

	err = os.MkdirAll(folder, os.FileMode(modePerm))

	return folder, nil
}

func LoadFile(path string) (*FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Name:   info.Name(),
		Data:   data,
		Format: format(info.Name()),
	}, nil
}

func format(name string) string {
	if p := strings.Split(name, "."); len(p) > 1 {
		return p[len(p)-1]
	}
	return ""
}
