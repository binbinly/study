package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func main()  {

	var data = []string{"test", "hello", "go"}

	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	if err := w.Write(data); err != nil {
		log.Fatal(err)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))


	r := csv.NewReader(&buf)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}