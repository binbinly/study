package maputil

import (
	"reflect"
	"strings"
)

// MergeStringMap simple merge two string map. merge src to dst map
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}

		dst[k] = v
	}
	return dst
}

// Keys get all keys of the given map.
func Keys(mp interface{}) (keys []string) {
	rftVal := reflect.ValueOf(mp)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}

	if rftVal.Kind() != reflect.Map {
		return
	}

	for _, key := range rftVal.MapKeys() {
		keys = append(keys, key.String())
	}
	return
}

// Values get all values from the given map.
func Values(mp interface{}) (values []interface{}) {
	rftTyp := reflect.TypeOf(mp)
	if rftTyp.Kind() == reflect.Ptr {
		rftTyp = rftTyp.Elem()
	}

	if rftTyp.Kind() != reflect.Map {
		return
	}

	rftVal := reflect.ValueOf(mp)
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}