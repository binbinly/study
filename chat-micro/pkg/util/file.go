package util

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//IsFileExist 文件是否存在
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// ImageType 图片类型
func ImageType(read io.Reader) string {
	src, err := ioutil.ReadAll(read)
	if err != nil {
		return ""
	}
	reader := bytes.NewReader(src)
	_, imgType, err := image.Decode(reader)
	if err != nil {
		return ""
	}
	return imgType
}

// GetFileExt 获取文件后缀名
func GetFileExt(filename string) string {
	index := strings.LastIndex(filename, ".")
	return filename[index:]
}

// Mkdir 创建目录
func Mkdir(path string) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}

//FileCreate 创建文件
func FileCreate(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content.Bytes())
	if err != nil {
		return err
	}
	return nil
}