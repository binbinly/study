package internal

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

const (
	compareNotEqual int = iota - 2
	compareLess
	compareEqual
	compareGreater
)

// Assert is a simple implementation of assertion, only for internal usage
type Assert struct {
	T        *testing.T
	CaseName string
}

// NewAssert return instance of Assert
func NewAssert(t *testing.T, caseName string) *Assert {
	return &Assert{T: t, CaseName: caseName}
}

// Equal check if expected is equal with actual
func (a *Assert) Equal(expected, actual interface{}) {
	if compare(expected, actual) != compareEqual {
		makeTestFailed(a.T, a.CaseName, expected, actual)
	}
}

// NotEqual check if expected is not equal with actual
func (a *Assert) NotEqual(expected, actual interface{}) {
	if compare(expected, actual) == compareEqual {
		expectedInfo := fmt.Sprintf("not %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Greater check if expected is greater than actual
func (a *Assert) Greater(expected, actual interface{}) {
	if compare(expected, actual) != compareGreater {
		expectedInfo := fmt.Sprintf("> %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// GreaterOrEqual check if expected is greater than or equal with actual
func (a *Assert) GreaterOrEqual(expected, actual interface{}) {
	isGreatOrEqual := compare(expected, actual) == compareGreater || compare(expected, actual) == compareEqual
	if !isGreatOrEqual {
		expectedInfo := fmt.Sprintf(">= %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Less check if expected is less than actual
func (a *Assert) Less(expected, actual interface{}) {
	if compare(expected, actual) != compareLess {
		expectedInfo := fmt.Sprintf("< %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// LessOrEqual check if expected is less than or equal with actual
func (a *Assert) LessOrEqual(expected, actual interface{}) {
	isLessOrEqual := compare(expected, actual) == compareLess || compare(expected, actual) == compareEqual
	if !isLessOrEqual {
		expectedInfo := fmt.Sprintf("<= %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// IsNil check if value is nil
func (a *Assert) IsNil(value interface{}) {
	if value != nil {
		makeTestFailed(a.T, a.CaseName, nil, value)
	}
}

// IsNotNil check if value is not nil
func (a *Assert) IsNotNil(value interface{}) {
	if value == nil {
		makeTestFailed(a.T, a.CaseName, "not nil", value)
	}
}

// compare x and y return :
// x > y -> 1, x < y -> -1, x == y -> 0, x != y -> -2
func compare(x, y interface{}) int {
	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)

	if vx.Type() != vy.Type() {
		return compareNotEqual
	}

	switch vx.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		xInt := vx.Int()
		yInt := vy.Int()
		if xInt > yInt {
			return compareGreater
		} else if xInt == yInt {
			return compareEqual
		} else {
			return compareLess
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		xUint := vx.Uint()
		yUint := vy.Uint()
		if xUint > yUint {
			return compareGreater
		}else if xUint == yUint {
			return compareEqual
		} else {
			return compareLess
		}
	case reflect.Float32, reflect.Float64:
		xFloat := vx.Float()
		yFloat := vy.Float()
		if xFloat > yFloat {
			return compareGreater
		} else if xFloat == yFloat {
			return compareEqual
		} else {
			return compareLess
		}
	case reflect.String:
		xString := vx.String()
		yString := vy.String()
		if xString > yString {
			return compareGreater
		} else if xString == yString {
			return compareEqual
		} else {
			return compareLess
		}
	default:
		if reflect.DeepEqual(x, y) {
			return compareEqual
		}
		return compareNotEqual
	}
}

// makeTestFailed make test failed and log error info
func makeTestFailed(t *testing.T, caseName string, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(2)
	errInfo := fmt.Sprintf("Case %v failed. file: %v, line: %v, expected: %v, actual: %v.", caseName, file, line, expected, actual)
	t.Error(errInfo)
	t.FailNow()
}