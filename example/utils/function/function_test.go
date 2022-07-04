package function

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"lib/utils/internal"
)

func TestAfter(t *testing.T) {
	arr := []string{"a", "b"}
	f := After(len(arr), func(i int) int {
		t.Log("print done")
		return i
	})
	type cb func(args ...interface{}) []reflect.Value
	print := func(i int, s string, fn cb) {
		t.Logf("print: arr[%d] is %s \n", i, s)
		v := fn(i)
		if v != nil {
			vv := v[0].Int()
			if vv != 1 {
				t.FailNow()
			}
		}
	}
	t.Log("print: arr is", arr)
	for i := 0; i < len(arr); i++ {
		print(i, arr[i], f)
	}
}

func TestBefore(t *testing.T) {
	assert := internal.NewAssert(t, "TestBefore")

	arr := []string{"a", "b", "c", "d", "e"}
	f := Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...interface{}) []reflect.Value
	appendStr := func(i int, s string, fn cb) {
		v := fn(i)
		t.Log("v", i, v[0])
		res = append(res, v[0].Int())
	}

	for i := 0; i < len(arr); i++ {
		appendStr(i, arr[i], f)
	}

	expected := []int64{0, 1, 2, 2, 2}
	assert.Equal(expected, res)
}

func TestCurry(t *testing.T) {
	assert := internal.NewAssert(t, "TestCurry")

	add := func(a, b int) int {
		return a + b
	}
	var addCurry Fn = func(values ...interface{}) interface{} {
		return add(values[0].(int), values[1].(int))
	}
	add1 := addCurry.Curry(1)
	assert.Equal(3, add1(2))
}

func TestCompose(t *testing.T) {
	assert := internal.NewAssert(t, "TestCompose")

	toUpper := func(a ...interface{}) interface{} {
		return strings.ToUpper(a[0].(string))
	}
	toLower := func(a ...interface{}) interface{} {
		return strings.ToLower(a[0].(string))
	}

	expected := toUpper(toLower("aBCde"))
	cf := Compose(toUpper, toLower)
	res := cf("aBCde")

	assert.Equal(expected, res)
}

func TestDelay(t *testing.T) {
	var print = func(s string) {
		t.Log(s)
	}
	Delay(2*time.Second, print, "test delay")
}

func TestDebounced(t *testing.T) {
	assert := internal.NewAssert(t, "TestDebounced")

	count := 0
	add := func() {
		count++
	}

	debouncedAdd := Debounced(add, 50*time.Microsecond)
	debouncedAdd()
	debouncedAdd()
	debouncedAdd()
	debouncedAdd()

	time.Sleep(100 * time.Millisecond)
	assert.Equal(1, count)

	debouncedAdd()
	time.Sleep(100 * time.Millisecond)
	assert.Equal(2, count)
}

func TestSchedule(t *testing.T) {
	assert := internal.NewAssert(t, "TestSchedule")

	var res []string
	appendStr := func(s string) {
		res = append(res, s)
	}

	stop := Schedule(1*time.Second, appendStr, "*")
	time.Sleep(5 * time.Second)
	close(stop)

	expected := []string{"*", "*", "*", "*", "*"}
	assert.Equal(expected, res)
}
