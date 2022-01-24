package util

import (
	"reflect"
	"regexp"
)

// IsZero 检查是否是零值
func IsZero(i ...interface{}) bool {
	for _, j := range i {
		v := reflect.ValueOf(j)
		if isZero(v) {
			return true
		}
	}
	return false
}

// ValidateMobile 验证手机号
func ValidateMobile(phone string) bool {
	// 170、171、165、162、167 排除
	regular := "^(((13[0-9]{1})|(15[0-9]{1})|(16[1346890]{1})|(17[2-8]{1})|(18[0-9]{1})|(19[0-9]{1})|(14[5-7]{1}))+\\d{8})$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	case reflect.Invalid:
		return true
	default:
		z := reflect.Zero(v.Type())
		return reflect.DeepEqual(z.Interface(), v.Interface())
	}
}
