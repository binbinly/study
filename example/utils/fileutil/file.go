package fileutil

import (
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Dir get dir path, without last name.
func Dir(path string) string {
	return filepath.Dir(path)
}

// Name get file/dir name
func Name(path string) string {
	return filepath.Base(path)
}

// Ext get filename ext. alias of path.Ext()
func Ext(aPath string) string {
	return path.Ext(aPath)
}

// IsExist checks if a file or directory exists
func IsExist(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreateFile create a file in path
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}
	defer file.Close()

	return true
}

// IsDir checks if the path is directory or not
func IsDir(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile reports whether the named file or directory exists.
func IsFile(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}

// IsAbsPath is abs path.
func IsAbsPath(aPath string) bool {
	return path.IsAbs(aPath)
}

// RemoveFile remove the path file
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile copy src file to dest file
func CopyFile(srcFilePath, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 4096)
	for {
		n, err := srcFile.Read(tmp)
		distFile.Write(tmp[:n])
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

//ClearFile write empty string to path file
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

//ReadFileToString return string of file content
func ReadFileToString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine read file line by line
func ReadFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}
		res = append(res, string(line))
	}

	return res, nil
}

// ListFileNames return all file names in the path
func ListFileNames(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}

	f, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	sz := len(f)
	if sz == 0 {
		return []string{}, nil
	}

	var res []string
	for i := 0; i < sz; i++ {
		if !f[i].IsDir() {
			res = append(res, f[i].Name())
		}
	}

	return res, nil
}

// IsLink checks if a file is symbol link or not
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// Mode return file's mode and permission
func Mode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// MimeType return file mime type
// param `file` should be string(file path) or *os.File
func MimeType(file interface{}) string {
	var mediaType string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediaType
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediaType
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediaType
		}
		return http.DetectContentType(buffer)
	}
	return mediaType
}