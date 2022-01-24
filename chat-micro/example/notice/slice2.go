package main

import (
	"bytes"
	"fmt"
)

// 错误使用 slice 的拼接示例
//拼接的结果不是正确的 AAAAsuffix/BBBBBBBBB，因为 dir1、 dir2 两个 slice 引用的数据都是 path 的底层数组，
//dir1 同时也修改了 path，也导致了 dir2 的修改
func joinErr() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	println(sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1))        // AAAA
	println("dir2: ", string(dir2))        // BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path))    // AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // uffixBBBB

	println("new path: ", string(path))    // AAAAsuffix/uffixBBBB    // 错误结果
}

// 使用 full slice expression
func join() {

	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	//第三个参数是用来控制 dir1 的新容量，再往 dir1 中 append 超额元素时，将分配新的 buffer 来保存。而不是覆盖原来的 path 底层数组
	dir1 := path[:sepIndex:sepIndex]        // 此时 cap(dir1) 指定为4， 而不是先前的 16
	dir2 := path[sepIndex+1:]
	dir1 = append(dir1, "suffix"...)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // BBBBBBBBB
	println("new path: ", string(path))    // AAAAsuffix/BBBBBBBBB
}

// 超过容量将重新分配数组来拷贝值、重新存储
//当你从一个已存在的 slice 创建新 slice 时，二者的数据指向相同的底层数组。如果你的程序使用这个特性，那需要注意 “旧”（stale） slice 问题。
//某些情况下，向一个 slice 中追加元素而它指向的底层数组容量不足时，将会重新分配一个新数组来存储数据。而其他 slice 还指向原来的旧底层数组。
func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)    // 3 3 [1 2 3 ]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2)    // 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}
	// 此时的 s1 与 s2 是指向同一个底层数组的
	fmt.Println(s1)        // [1 22 23]
	fmt.Println(s2)        // [22 23]

	s2 = append(s2, 4)    // 向容量为 2 的 s2 中再追加元素，此时将分配新数组来存

	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s1)        // [1 22 23]    // 此时的 s1 不再更新，为旧数据
	fmt.Println(s2)        // [32 33 14]
}