package function

import (
	"fmt"
	"reflect"
)

func invokeFunc(fn interface{}, args ...interface{}) []reflect.Value {
	fv := functionValue(fn)
	params := make([]reflect.Value, len(args))
	for i, item := range args {
		params[i] = reflect.ValueOf(item)
	}
	return fv.Call(params)
}

func unsafeInvokeFunc(fn interface{}, args ...interface{}) []reflect.Value {
	fv := reflect.ValueOf(fn)
	params := make([]reflect.Value, len(args))
	for i, item := range args {
		params[i] = reflect.ValueOf(item)
	}
	return fv.Call(params)
}

func functionValue(fn interface{}) reflect.Value {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", fn))
	}

	return v
}

func mustBeFunction(fn interface{})  {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", fn))
	}
}