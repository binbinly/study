package main

import (
	"log"
	"reflect"
)

func main()  {
	slice := &[]int64{0,1,2,3,4,5,6,7,8,9}
	/*log.Println(reflect.TypeOf(slice).Len())
	  log.Println(reflect.TypeOf(slice).Type().Len())
	  log.Println(reflect.ValueOf(slice).Type().Len())
	  log.Println(reflect.ValueOf(slice).Elem().Type().Len())
	*/
	log.Println(reflect.ValueOf(slice).Elem().Len())
}