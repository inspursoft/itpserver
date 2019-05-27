package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func CheckDirs(dir ...string) (targetPath string, err error) {
	targetPath = filepath.Join(dir...)
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		err = os.MkdirAll(targetPath, 0755)
	}
	return
}

func CheckFileExt(fileName string, ext string) bool {
	return filepath.Ext(fileName) == ext
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
