package main

import (
	"chat/pkg/app"
	"chat/pkg/log"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	file()
}

func file() io.Writer {
	filename := "./a/b/c/info.log"
	dirname := filepath.Dir(filename)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		log.Fatal(err)
	}
	// if we got here, then we need to create a file
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fh.WriteString("aaaaaa")
	return fh
}

func token()  {
	fmt.Printf("max:%v\n", 1 << 10)
	p := map[string]interface{}{"user_id": 1, "username": "test"}
	t, _ := app.Sign(nil, p, "UCAYyw9S5Q9oS2Bh1GhXZZmOawfiGSZXbuYR6KcYvidfhoGOcwsk8zb7vwpsd37o", 86400)
	fmt.Printf("token:%v\n", t)
}

func j()  {
	str := `{"id":1,"name":"san","data":{"a":1}}`
	type msg struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Data json.RawMessage `json:"data"`
	}
	m := &msg{}
	err := json.Unmarshal([]byte(str), m)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("msg:%v\n", m)
}