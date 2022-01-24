package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 在 defer 函数中参数会提前求值
// 对 defer 延迟执行的函数，它的参数会在声明时候就会求出具体值，而不是在执行时才求值：
func Defer() {
	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++
}

// 命令行参数指定目录名
// 遍历读取目录下的文件
func main() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	dir := os.Args[1]
	start, err := os.Stat(dir)
	if err != nil || !start.IsDir() {
		os.Exit(2)
	}

	var targets []string
	filepath.Walk(dir, func(fPath string, fInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fInfo.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fPath)
		return nil
	})

	for _, target := range targets {
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target:", target, "error:", err)    //error:too many open files
			break
		}
		defer f.Close()    // 在每次 for 语句块结束时，不会关闭文件资源

		// 使用 f 资源
	}
}
