package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func jsonUn() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	//Go 默认会将数值当做 float64 处理
	fmt.Printf("%T\n", result["status"])    // float64
	var status = uint64(result["status"].(float64))
	fmt.Println("Status value: ", status)
}

// 指定字段类型
func jsonUn2() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status, _ = result["status"].(json.Number).Int64()
	fmt.Println("Status value: ", status)
}

// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
// 将数据转为 decode 为 string
func jsonUn3() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("Status value: ", status)
}

// struct 中指定字段类型
func jsonStruct() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("Result: %+v", result)
}

// 状态名称可能是 int 也可能是 string，指定为 json.RawMessage 类型
func main() {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct{
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		if err != nil{
			log.Fatalln(err)
		}

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}