package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const Dir = "example/testdata"
const File = "example/testdata/test"

func main()  {
	var data = []byte("go is a general-purpose language designed with systems programming in mind.")

	var r = bytes.NewReader(data)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	ioutil.WriteFile(File, data, os.ModePerm)

	bf, err := ioutil.ReadFile(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bf))

	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		fmt.Println(file.ModTime())
		fmt.Println(file.Mode())
		fmt.Println(file.IsDir())
		fmt.Println(file.Size())
		fmt.Println(file.Sys())
	}

	rc := ioutil.NopCloser(r)
	if err = rc.Close(); err != nil {
		log.Fatal(err)
	}

	exampleTempFile()

	exampleTempDir()
}

func exampleTempDir() {
	content := []byte("temporary file's content")

	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	fmt.Println("dir:", dir)

	tmpFn := filepath.Join(dir, "tmpfile")
	if err = ioutil.WriteFile(tmpFn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func exampleTempFile()  {
	data := []byte("temporary file's content")

	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	fmt.Println("file:", tmpFile)

	if _, err = tmpFile.Write(data); err != nil {
		log.Fatal(err)
	}

	if err = tmpFile.Close();err != nil {
		log.Fatal(err)
	}
}