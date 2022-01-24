package main

import (
	"encoding/json"
	"fmt"
)

type msg struct {
	Event string `json:"event"`
	Data json.RawMessage `json:"data"`
}

func main()  {
	str := `{"event":"test","data":{"name":"1","age":18}}`
	m := &msg{}
	err := json.Unmarshal([]byte(str), m)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("msg:%+v\n", m)
	fmt.Printf("data:%s\n", m.Data)
}
