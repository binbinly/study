package reflect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	A int    `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

// 通过反射修改非导出字段
func TestChangeNotExportFiled(t *testing.T) {
	var r TestStruct
	r.A = 100
	r.B = "B"
	r.C = "C"
	getValue := reflect.ValueOf(r)
	if getValue.Kind() != reflect.Struct {
		panic("need struct kind")
	}
	getType := reflect.TypeOf(r)

	for i := 0; i < getValue.NumField(); i++ {
		t.Logf("name: %s, type: %s, value: %v\n", getType.Field(i).Name, getValue.Field(i).Type(), getValue.Field(i).Interface())
	}

	t.Logf("%+v", r)
}

// 根据反射判断字段类型
func TestInterface(t *testing.T) {
	var value interface{}
	value = "pibigstar"
	switch value.(type) {
	case string:
		v, ok := value.(string)
		if ok {
			t.Logf("String ==> %s \n", v)
		}
	case map[string]string:
		v, ok := value.(map[string]string)
		if ok {
			t.Logf("Map ==> %v \n", v)
		}
	default:
		bs, _ := json.Marshal(value)
		t.Logf("Others ==> %s \n", string(bs))
	}
}

// 反射基本操作
func TestReflect(t *testing.T) {
	var str = "hello world"

	v := reflect.ValueOf(str)
	t.Log("value:", v)
	t.Log("value:", v.String())

	t.Log("type:", v.Type())
	t.Log("kind:", v.Kind())

	canSet := v.CanSet()
	t.Log("can set:", canSet)

	v = reflect.ValueOf(&str)
	v = v.Elem()
	v.SetString("new world")

	t.Log("value:", v)

	test := TestStruct{A: 23, B: "Hello World"}
	s := reflect.ValueOf(&test).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Logf("%s: Type ==>%s Value==> %v \n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("new world")
	t.Logf("%+v", test)
}

// 获取tag
func TestGetTag(t *testing.T) {
	s := TestStruct{}
	rt := reflect.TypeOf(s)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fmt.Println(f.Tag.Get("json"))
	}
}

// 处理不定数量的chan
func TestChan(t *testing.T) {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 20)
	ch3 := make(chan string, 30)
	cases := createCases(ch1, ch2, ch3)

	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() && ok {
			t.Log("recv:", recv)
		} else {
			t.Log("send:", cases[chosen].Send)
		}
	}
}

func createCases(chs ...chan string) []reflect.SelectCase {
	var cases []reflect.SelectCase
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	for i, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(fmt.Sprintf("Hello: %d", i)), // 发送的send值
		})
	}
	return cases
}