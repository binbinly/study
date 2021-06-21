package archive

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

//ZipWrite 压缩成zip
func ZipWrite(files []*File) (buf *bytes.Buffer, err error) {
	buf = new(bytes.Buffer)

	// 初始化writer
	w := zip.NewWriter(buf)
	defer w.Close()

	// 设置压缩级别，不指定则默认
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	var f io.Writer
	for _, file := range files {

		// 根据文件名称，writer创建文件
		f, err = w.Create(file.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "file create")
		}
		// 创建的文件写入内容
		_, err = f.Write(file.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "file write")
		}
	}
	return
}

//ZipRead 读取zip压缩文件
func ZipRead(path string) ([]*File, error) {

	// 根据文件路径，获取zip文件内容
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, errors.Wrapf(err, "open reader")
	}
	// 方法最后调用关闭
	defer r.Close()

	// 循环读取多个文件内容
	var rc io.ReadCloser
	var b []byte
	var files []*File
	for _, f := range r.File {
		err = func(f *zip.File) error {
			// 打开文件
			rc, err = f.Open()
			if err != nil {
				return errors.Wrapf(err, "file open")
			}
			defer rc.Close()

			// 读取文件内容
			b, err = ioutil.ReadAll(rc)
			if err != nil {
				return errors.Wrapf(err, "file read all")
			}
			file := &File{
				Name: f.Name,
				Body: b,
			}
			files = append(files, file)
			return nil
		}(f)
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}
