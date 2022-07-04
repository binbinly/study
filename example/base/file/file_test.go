package file

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	if IsCI() {
		return
	}
	m.Run()
}

const fileName = "demo.log"

func TestCreateFile(t *testing.T) {
	CreateFile(fileName)
}

func TestWriteFile(t *testing.T) {
	WriteFile(fileName, "Hello,World")
}

func TestAppendToFile(t *testing.T) {
	AppendToFile(fileName, "Hello,Pibigstar")
}

func TestReadFile(t *testing.T) {
	ReadFile(fileName)
}

func TestMkdirFile(t *testing.T) {
	MkOneDir("demo")
	MkAllDir("test/user/one")
}

func TestReadAllDir(t *testing.T) {
	ReadAllDir(".")
}

func TestDeleteFile(t *testing.T) {
	DeleteFile(fileName)
}

func TestFileAbs(t *testing.T) {
	t.Log(GetFileAbs("file.go"))
}

func TestInode(t *testing.T) {
	t.Log(Inode("file.go"))
}

func TestCopy(t *testing.T) {
	fileMd5, err := GetFileMd5("file.go")
	if err != nil {
		t.Error(err)
	}
	t.Log(fileMd5)
}

func IsCI() bool {
	name, _ := os.Hostname()
	if name != "pibigstar" && runtime.GOOS == "linux" {
		return true
	}
	return false
}