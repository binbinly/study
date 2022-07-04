package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// 使用ioutil读取文件
func ReadFile(filename string)  {
	data, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Println(string(data))
}

//读取文件夹
func ReadAllDir(path string)  {
	files, err := ioutil.ReadDir(path)
	check(err)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 这种会覆盖掉原先内容
func WriteFile(filename, data string)  {
	err := ioutil.WriteFile(filename, []byte(data), os.ModePerm)
	check(err)
}

// 追加内容到文件末尾
func AppendToFile(filename, data string)  {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer func() {
		file.Close()
	}()
	check(err)
	file.Write([]byte(data))
}

// 创建文件并返回文件指针
func CreateFile(filename string)  {
	// 如果源文件已存在，会清空该文件的内容
	// 如果多级目录，某一个目录不存在，则会返回PathError
	file, err := os.Create(filename)
	defer func() {
		file.Close()
	}()
	check(err)
}

// 创建单个文件夹
func MkOneDir(dir string)  {
	err := os.Mkdir(dir, os.ModePerm)
	check(err)
	os.RemoveAll(dir)
}

// 创建多层文件夹
func MkAllDir(dirs string)  {
	// 如果不存在，才创建
	if !IsExist(dirs) {
		err := os.MkdirAll(dirs, os.ModePerm)
		check(err)
		fmt.Println("dirs", strings.Split(dirs, "/"))
		os.RemoveAll(strings.Split(dirs, "/")[0])
	}
}

// 删除文件
func DeleteFile(filename string)  {
	err := os.Remove(filename)
	check(err)
}

func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 返回该文件的绝对路径
func GetFileAbs(path string) string {
	if absPath, err := filepath.Abs(path); err == nil {
		return absPath
	}
	return ""
}

// 获取文件inode号
func Inode(filename string) uint64 {
	fileinfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return 0
	}
	stat, ok := fileinfo.Sys().(*syscall.Stat_t)
	if !ok {
		return 0
	}
	return stat.Ino
}