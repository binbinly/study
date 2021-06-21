package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	example()

	//exampleDir()

	exampleFile()

	exampleChange()

	// 文件操作模式
	exampleFileMode()
	// 环境变量
	exampleEnv()
	// 进程操作
	exampleProcess()
	// 获取命令行参数
	exampleArgs()
}

const (
	DirPath     = "example/testdata/"
	FilePath    = "example/testdata/test_os.txt"
	NewFilePath = "example/testdata/test_os_new.txt"
)

func example() {

	fmt.Println(os.Hostname())

	fmt.Println("aqesize", os.Getpagesize())

	fmt.Println("uid:", os.Getuid())

	fmt.Println("euid", os.Geteuid())

	fmt.Println("gid", os.Getgid())

	fmt.Println("egid", os.Getegid())

	fmt.Println(os.Getgroups())

	fmt.Println("pid", os.Getpid())

	fmt.Println("ppid", os.Getppid())

	fmt.Println(os.Environ())

	fmt.Println(os.Getwd())

	fmt.Println(os.TempDir())

	fmt.Println(os.UserCacheDir())

	fmt.Println(os.UserHomeDir())

}

func exampleDir() {

	if err := os.Mkdir(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func exampleFile() {

	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
}

func exampleChange() {

	if _, err := os.Stat(FilePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
		} else {
			log.Fatal(err)
		}
	}

	if err := os.Chmod(FilePath, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.Chown(FilePath, 501, 20); err != nil {
		log.Fatal(err)
	}

	if err := os.Lchown(FilePath, 501, 20); err != nil {
		log.Fatal(err)
	}

	mtime := time.Date(2019, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2019, time.March, 2, 4, 5, 6, 0, time.UTC)

	if err := os.Chtimes(FilePath, atime, mtime); err != nil {
		log.Fatal(err)
	}

}

func exampleFileMode()  {

	fi, err := os.Lstat(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("permissitions:%#o\n", fi.Mode().Perm())

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")


	}

}

func exampleEnv()  {

	if err := os.Setenv("NAME", "Gopher"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("NAME"))

	fmt.Println(os.LookupEnv("NAME"))

	fmt.Println(os.Expand("Hello $NAME", func(s string) string {
		return "Gopher"

	}))

	fmt.Println(os.ExpandEnv("Hello ${NAME}"))

	if err := os.Unsetenv("NAME"); err != nil {
		log.Fatal(err)
	}
}

func exampleProcess()  {

	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	p, err := os.StartProcess("/usr/bin/vim", []string{"/usr/bin/vim", "temp.txt"}, attr)
	if err != nil {
		log.Fatal(err)
	}

	p, err = os.FindProcess(p.Pid)
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Kill(); err != nil {
		log.Fatal(err)
	}
}

func exampleArgs()  {
	args := os.Args[1:]
	fmt.Println(args)
}