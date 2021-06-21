package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	example()
	exampleType()
	exampleValue()
	exampleMake()
}

type Hello struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Arr  []string `json:"arr"`
}

func (h *Hello) Say() {
	fmt.Println("Hello World")
}

func (h Hello) Says() {
	fmt.Println("Hello world s")
}

func exampleType() {
	hello := Hello{}
	t := reflect.TypeOf(hello)

	reflect.TypeOf(&hello).Elem()

	t.Kind()
	fmt.Println("reflect.Struct:", t.Kind() == reflect.Struct)

	fmt.Println("Name:", t.Name())

	fmt.Println("PkgPath", t.PkgPath())

	fmt.Println("String", t.String())

	fmt.Println("Size", t.Size())

	fmt.Println("Align", t.Align())

	fmt.Println("FieldAlign", t.FieldAlign())

	fmt.Println("NumMethod", t.NumMethod())

	fmt.Println("Method", t.Method(0))

	fmt.Println(t.MethodByName("Says"))

	fmt.Println("NumField", t.NumField())

	fmt.Println("Field", t.Field(0))

	fmt.Println("Field Tag(json)", t.Field(0).Tag.Get("jsons"))

	fmt.Println(t.Field(2).Tag.Lookup("json"))

	fmt.Println("FieldByIndex", t.FieldByIndex([]int{1}))

	fmt.Println(t.FieldByName("Name"))

	fmt.Println(t.FieldByNameFunc(func(s string) bool {
		return s == "Name"
	}))

	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil))

	fmt.Println(fileType.Implements(writerType))

	fmt.Println(fileType.AssignableTo(writerType))

	fmt.Println(fileType.ConvertibleTo(writerType))
}

func exampleValue() {
	hello := Hello{}

	v := reflect.ValueOf(hello)

	reflect.ValueOf(&hello).Elem()

	v.Type()

	v.Kind()
	fmt.Println("reflect.String:", v.Kind() == reflect.String)

	fmt.Println(v.IsValid())

	fmt.Println(v.NumMethod())

	fmt.Println(v.CanAddr())

	fmt.Println(v.CanInterface())

	fmt.Println(v.Interface())

	fmt.Println(v.CanSet())

}

func exampleMake()  {

	var c chan int
	var fn func(int) int
	var m map[string]string
	var sl = []int{1,2,3}

	reflect.MakeChan(reflect.TypeOf(c), 1)

	reflect.MakeFunc(reflect.TypeOf(fn), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{args[0]}
	})

	reflect.MakeMap(reflect.TypeOf(m))

	reflect.MakeMapWithSize(reflect.TypeOf(m), 10)

	reflect.MakeSlice(reflect.ValueOf(sl).Type(), reflect.ValueOf(sl).Len(), reflect.ValueOf(sl).Cap())

	reflect.ChanOf(reflect.BothDir, reflect.TypeOf(c))

	reflect.FuncOf([]reflect.Type{reflect.TypeOf(15)}, []reflect.Type{reflect.TypeOf("Hello")}, false)

	reflect.MapOf(reflect.TypeOf(15), reflect.TypeOf("Hello"))

	reflect.ArrayOf(5, reflect.TypeOf(15))

	reflect.SliceOf(reflect.ValueOf(sl).Type())

	reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag: `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag: `json:"age"`,
		},
	})
}

func example()  {

	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag: `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag: `json:"age"`,
		},
	})

	reflect.New(typ)

	reflect.PtrTo(typ)

	reflect.Zero(typ)

	sl := []int{1,5,6,7,3}
	reflect.Swapper(sl)

	var dst []byte
	var src = []byte("Hello World")

	reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))

	reflect.DeepEqual("1", 1)

}