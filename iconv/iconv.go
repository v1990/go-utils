// Package iconv :interface convert
package iconv

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"math"
	"reflect"
	"strconv"
)

// IsNil check interface value is nil
func IsNil(v interface{}) (isNil bool) {
	if v == nil {
		return true
	}

	vv := reflect.ValueOf(v)
	switch vv.Kind() {
	case
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice,
		reflect.Ptr:
		return vv.IsNil()
	}

	return false
}

// ToString interface to string
func ToString(v interface{}, def string) string {
	if v == nil {
		return def
	}

	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case complex64:
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)
	case complex128:
		return strconv.FormatComplex(complex128(v), 'f', -1, 128)

	case *string:
		if v == nil {
			return def
		}
		return *v
	case *bool:
		if v == nil {
			return def
		}

		return strconv.FormatBool(*v)
	case *int:
		if v == nil {
			return def
		}

		return strconv.FormatInt(int64(*v), 10)
	case *int8:
		if v == nil {
			return def
		}

		return strconv.FormatInt(int64(*v), 10)
	case *int16:
		if v == nil {
			return def
		}

		return strconv.FormatInt(int64(*v), 10)
	case *int32:
		if v == nil {
			return def
		}

		return strconv.FormatInt(int64(*v), 10)
	case *int64:
		if v == nil {
			return def
		}

		return strconv.FormatInt(int64(*v), 10)
	case *uint:
		if v == nil {
			return def
		}

		return strconv.FormatUint(uint64(*v), 10)
	case *uint8:
		if v == nil {
			return def
		}

		return strconv.FormatUint(uint64(*v), 10)
	case *uint16:
		if v == nil {
			return def
		}

		return strconv.FormatUint(uint64(*v), 10)
	case *uint32:
		if v == nil {
			return def
		}

		return strconv.FormatUint(uint64(*v), 10)
	case *uint64:
		if v == nil {
			return def
		}

		return strconv.FormatUint(uint64(*v), 10)
	case *float32:
		if v == nil {
			return def
		}

		return strconv.FormatFloat(float64(*v), 'f', -1, 32)
	case *float64:
		if v == nil {
			return def
		}

		return strconv.FormatFloat(float64(*v), 'f', -1, 64)
	case *complex64:
		if v == nil {
			return def
		}

		return strconv.FormatComplex(complex128(*v), 'f', -1, 64)
	case *complex128:
		if v == nil {
			return def
		}

		return strconv.FormatComplex(complex128(*v), 'f', -1, 128)
	case fmt.Stringer:
		if v == nil {
			return def
		}

		return v.String()
	case error:
		if v == nil {
			return def
		}

		return v.Error()
	}

	return fmt.Sprintf("%v", v)
}

func ToInt64(v interface{}, def int64) int64 {
	if v, ok := ParseInt64(v); ok {
		return v
	}
	return def
}

func ToUInt64(v interface{}, def uint64) uint64 {
	if v, ok := ParseUInt64(v); ok {
		return v
	}
	return def
}

type I2Int interface{ ToInt() int }

type ProtoEnumNumber interface {
	Number() protoreflect.EnumNumber
}

func ParseInt64(v interface{}) (int64, bool) {
	if v == nil {
		return 0, false
	}

	{
		switch v := v.(type) {
		case int8:
			return int64(v), true
		case int16:
			return int64(v), true
		case int32:
			return int64(v), true
		case int64:
			return int64(v), true
		case int:
			return int64(v), true

		case uint8:
			return int64(v), true
		case uint16:
			return int64(v), true
		case uint32:
			return int64(v), true
		case uint64:
			if v <= math.MaxInt64 {
				return int64(v), true
			}
		case uint:
			if v <= math.MaxInt64 {
				return int64(v), true
			}

		case string:
			if vv, err := strconv.ParseInt(v, 10, 64); err == nil {
				return vv, true
			}

		case *int8:
			if v == nil {
				return 0, false
			}
			return int64(*v), true
		case *int16:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *int32:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *int64:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *int:
			if v == nil {
				return 0, false
			}

			return int64(*v), true

		case *uint8:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *uint16:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *uint32:
			if v == nil {
				return 0, false
			}

			return int64(*v), true
		case *uint64:
			if v == nil {
				return 0, false
			}

			if (*v) <= math.MaxInt64 {
				return int64(*v), true
			}
		case *uint:
			if v == nil {
				return 0, false
			}

			if (*v) <= math.MaxInt64 {
				return int64(*v), true
			}

		case *string:
			if v == nil {
				return 0, false
			}

			if vv, err := strconv.ParseInt(*v, 10, 64); err == nil {
				return vv, true
			}

		// NOTE: interface{}((*int64)(nil))!=nil
		//  See: TestInterfaceNil

		case I2Int:
			if v == nil {
				return 0, false
			}
			return int64(v.ToInt()), true

		case ProtoEnumNumber:
			if v == nil {
				return 0, false
			}
			return int64(v.Number()), true
		}
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if vv := val.Uint(); vv <= math.MaxInt64 {
			return int64(vv), true
		}
	}

	return 0, false
}

func ParseUInt64(v interface{}) (uint64, bool) {
	if v == nil {
		return 0, false
	}

	{
		switch v := v.(type) {
		case int8:
			if v < 0 {
				return 0, false
			}
			return uint64(v), true
		case int16:
			if v < 0 {
				return 0, false
			}

			return uint64(v), true
		case int32:
			if v < 0 {
				return 0, false
			}

			return uint64(v), true
		case int64:
			if v < 0 {
				return 0, false
			}

			return uint64(v), true
		case int:
			if v < 0 {
				return 0, false
			}

			return uint64(v), true

		case uint8:
			return uint64(v), true
		case uint16:
			return uint64(v), true
		case uint32:
			return uint64(v), true
		case uint64:
			return uint64(v), true
		case uint:
			return uint64(v), true

		case string:
			if vv, err := strconv.ParseUint(v, 10, 64); err == nil {
				return vv, true
			}

		case *int8:
			if v == nil || *v < 0 {
				return 0, false
			}
			return uint64(*v), true
		case *int16:
			if v == nil || *v < 0 {
				return 0, false
			}

			return uint64(*v), true
		case *int32:
			if v == nil || *v < 0 {
				return 0, false
			}

			return uint64(*v), true
		case *int64:
			if v == nil || *v < 0 {
				return 0, false
			}

			return uint64(*v), true
		case *int:
			if v == nil || *v < 0 {
				return 0, false
			}
			return uint64(*v), true

		case *uint8:
			if v == nil {
				return 0, false
			}
			return uint64(*v), true
		case *uint16:
			if v == nil {
				return 0, false
			}
			return uint64(*v), true
		case *uint32:
			if v == nil {
				return 0, false
			}
			return uint64(*v), true
		case *uint64:
			if v == nil {
				return 0, false
			}
			return uint64(*v), true
		case *uint:
			if v == nil {
				return 0, false
			}
			return uint64(*v), true

		case *string:
			if v == nil {
				return 0, false
			}

			if vv, err := strconv.ParseUint(*v, 10, 64); err == nil {
				return vv, true
			}

		// NOTE: interface{}((*int64)(nil))!=nil
		//  See: TestInterfaceNil

		case I2Int:
			if v == nil {
				return 0, false
			}
			return uint64(v.ToInt()), true

		case ProtoEnumNumber:
			if v == nil {
				return 0, false
			}
			return uint64(v.Number()), true
		}
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if vv := val.Int(); vv > 0 {
			return uint64(vv), true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint(), true
	}

	return 0, false
}
