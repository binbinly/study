package main

import (
	"fmt"
	"path"
)

func main()  {

	var pathStr = "example/testdata/test/"

	fmt.Println(path.Base(pathStr))

	fmt.Println(path.Clean("a//c"))

	fmt.Println(path.Dir(pathStr))

	fmt.Println(path.Ext("/a/test.html"))

	fmt.Println(path.IsAbs("/dev/null"))

	dir, file := path.Split(pathStr)
	fmt.Println(dir, file)

	fmt.Println(path.Join("a", "b/c"))

}
