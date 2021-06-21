package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

func main() {
	// 编码/解码
	example()
	// 多个数字一起编码/解码
	multiExample()

	ByteOrder()
	uvarint()
}

func example() {
	var order = binary.LittleEndian

	var data = math.Pi

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, order, data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", buf.Bytes())

	var origin float64
	if err := binary.Read(buf, order, &origin); err != nil {
		log.Fatal(err)
	}
	fmt.Println(origin)
}

func multiExample() {
	buf := new(bytes.Buffer)

	var data = []interface{}{
		uint16(61374),
		int8(-65),
		uint8(254),
	}

	for _, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(buf.Bytes())

	var origin = struct {
		A uint16
		B int8
		C uint8
	}{}
	if err := binary.Read(buf, binary.LittleEndian, &origin); err != nil {
		log.Fatal(err)
	}
	fmt.Println(origin)
}

func ByteOrder()  {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint16(b[:2], '1')
	binary.LittleEndian.PutUint16(b[:2], 0x07d0)
	fmt.Printf("% x\n", b)

	x1 := binary.LittleEndian.Uint16(b[:2])
	x2 := binary.LittleEndian.Uint16(b[2:])
	fmt.Printf("%#04x %#04x\n", x1, x2)
}

func uvarint()  {
	var n uint64 = 256
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, n)

	i, err := binary.ReadUvarint(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}