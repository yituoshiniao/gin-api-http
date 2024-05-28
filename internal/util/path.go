package util

import (
	"context"
	"os"
	"path/filepath"

	"github.com/yituoshiniao/kit/xlog"
)

func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateFile(filePath string) error {
	if !Exists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func CreateDir(ctx context.Context, filePath string) error {
	dir := filepath.Dir(filePath)
	if !Exists(dir) {
		err := CreateFile(dir)
		if err != nil {
			xlog.S(ctx).Errorw("os.Mkdir--错误", "err", err)
			return err
		}
	}
	return nil
}
