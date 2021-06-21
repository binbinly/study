package archive

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestTar(t *testing.T)  {
	var files = []*File{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}

	// 写tar文件数据流
	buf, err := TarWrite(files)
	if err != nil {
		t.Fatal(err)
	}

	var filepath = "test.tar"
	// 自动生成并写入文件
	if err = ioutil.WriteFile(filepath, buf.Bytes(), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	fs, err := TarRead(filepath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, files, fs)
}

func TestZip(t *testing.T)  {
	var files = []*File{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling license.")},
	}

	// 写tar文件数据流
	buf, err := ZipWrite(files)
	if err != nil {
		t.Fatal(err)
	}

	var filepath = "test.zip"
	// 自动生成并写入文件
	if err = ioutil.WriteFile(filepath, buf.Bytes(), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	fs, err := ZipRead(filepath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, files, fs)
}