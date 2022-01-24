package util

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// StringSliceReflectEqual 判断 string和slice 是否相等
// 因为使用了反射，所以效率较低，可以看benchmark结果
func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

// StringSliceEqual 判断 string和slice 是否相等
// 使用了传统的遍历方式
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// reflect.DeepEqual的结果保持一致
	if (a == nil) != (b == nil) {
		return false
	}

	// bounds check 边界检查
	// 避免越界
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// SliceShuffle shuffle a slice
func SliceShuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// Uint64SliceReverse 对uint64 slice 反转
func Uint64SliceReverse(a []uint64) []uint64 {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// StringSliceContains 字符串切片中是否包含另一个字符串
// 来自go源码 net/http/server.go
func StringSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// IsInSlice 判断某一值是否在slice中
// 因为使用了反射，所以时间开销比较大，使用中根据实际情况进行选择
func IsInSlice(value interface{}, sli interface{}) bool {
	switch reflect.TypeOf(sli).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(sli)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// InIntSlice 判断整型值是否在slice中
func InIntSlice(v int, sli []int) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}

// InInt64Slice 判断整型值是否在slice中
func InInt64Slice(v int64, sli []int64) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}

// InuInt32Slice 判断是否在切片中存在
func InuInt32Slice(v uint32, sli []uint32) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}

// InStringSlice 判断字符串值是否在slice中
func InStringSlice(v string, sli []string) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}

// Uint64ShuffleSlice 对slice进行随机
func Uint64ShuffleSlice(a []uint64) []uint64 {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// see: https://yourbasic.org/golang/

// Uint64DeleteElemInSlice 从slice删除元素
// fast version, 会改变顺序
// i：slice的索引值
// s: slice
func Uint64DeleteElemInSlice(i int, s []uint64) []uint64 {
	if i < 0 || i > len(s)-1 {
		return s
	}
	// Remove the element at index i from s.
	s[i] = s[len(s)-1] // Copy last element to index i.
	s[len(s)-1] = 0    // Erase last element (write zero value).
	s = s[:len(s)-1]   // Truncate slice.

	return s
}

// Uint64DeleteElemInSliceWithOrder 从slice删除元素
// slow version, 保持原有顺序
// i：slice的索引值
// s: slice
func Uint64DeleteElemInSliceWithOrder(i int, s []uint64) []uint64 {
	if i < 0 || i > len(s)-1 {
		return s
	}
	// Remove the element at index i from s.
	copy(s[i:], s[i+1:]) // Shift s[i+1:] left one index.
	s[len(s)-1] = 0      // Erase last element (write zero value).
	s = s[:len(s)-1]     // Truncate slice.

	return s
}

// SliceIntToString1 整型数组拼接成字符串
func SliceIntToString1(s []int) string {
	if len(s) < 1 {
		return ""
	}

	ss := strconv.Itoa(s[0])
	for i := 1; i < len(s); i++ {
		ss += "," + strconv.Itoa(s[i])
	}

	return ss
}

// SliceIntToString2 整型数组拼接成字符串 性能最优
func SliceIntToString2(s []int) string {
	if len(s) < 1 {
		return ""
	}

	var str strings.Builder
	str.WriteString(strconv.Itoa(s[0]))
	for i := 1; i < len(s); i++ {
		str.WriteString(",")
		str.WriteString(strconv.Itoa(s[i]))
	}

	return str.String()
}

// SliceUInt32ToString2 整型数组拼接成字符串 性能最优
func SliceUInt32ToString2(s []uint32) string {
	if len(s) < 1 {
		return ""
	}

	var str strings.Builder
	str.WriteString(strconv.Itoa(int(s[0])))
	for i := 1; i < len(s); i++ {
		str.WriteString(",")
		str.WriteString(strconv.Itoa(int(s[i])))
	}

	return str.String()
}

// SliceInt64ToString2 整型数组拼接成字符串 性能最优
func SliceInt64ToString2(s []int64, sep string) string {
	if len(s) < 1 {
		return ""
	}

	var str strings.Builder
	str.WriteString(strconv.Itoa(int(s[0])))
	for i := 1; i < len(s); i++ {
		str.WriteString(sep)
		str.WriteString(strconv.Itoa(int(s[i])))
	}

	return str.String()
}

// SliceIntToString3 整型数组拼接成字符串 适用大数组
func SliceIntToString3(s []int) string {
	if len(s) < 1 {
		return ""
	}

	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}

	return *(*string)(unsafe.Pointer(&b))
}

// SliceToInt 字符串数组转换成整型数组
func SliceToInt(ss []string) (ii []int) {
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		ii = append(ii, i)
	}
	return ii
}

// SliceTouInt32 切换转int32
func SliceTouInt32(ss []string) (ii []uint32) {
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		ii = append(ii, uint32(i))
	}
	return ii
}

// FilterBigIntSlice 过滤切片元素 适合大切片
func FilterBigIntSlice(a []int, f func(v int) bool) []int{
	for i := 0; i < len(a); i++ {
		if !f(a[i]) {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

// FilterSmallIntSlice 过滤切片元素 适合小切片
func FilterSmallIntSlice(a []int, f func(v int) bool) []int {
	ret := make([]int, 0, len(a))
	for _, val := range a {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// FilterSmallUInt32Slice 过滤切片
func FilterSmallUInt32Slice(a []uint32, f func(v uint32) bool) []uint32 {
	ret := make([]uint32, 0, len(a))
	for _, val := range a {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}