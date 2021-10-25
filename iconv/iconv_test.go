package iconv

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type stringer string

func (s stringer) String() string { return "stringer:" + string(s) }

func TestToString(t *testing.T) {
	s1 := "test"
	tests := []struct {
		name string
		val  interface{}
		want string
	}{
		{"nil", nil, ""},
		{"nilInt", (*int)(nil), ""},
		{"nilInterface", interface{}((*int)(nil)), ""},

		{"float32", float32(3.141592654), "3.1415927"},
		{"float64", float64(3.141592654), "3.141592654"},
		{"error", errors.New("test"), "test"},
		{"string", string("test"), "test"},
		{"stringPtr", (*string)(&s1), s1},
		{"stringer", stringer("test"), "stringer:test"},
		{"complex64", complex64(complex(1, 2)), "(1+2i)"},
		{"complex128", complex128(complex(1, 2)), "(1+2i)"},
		{"complex128-1", 1i * 1i, "(-1+0i)"},
		{"complex128-2", 2 + 1i, "(2+1i)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.val, ""); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceNil(t *testing.T) {
	var v interface{} = (*int)(nil)
	var res string
	if v == nil {
		res = fmt.Sprintf("interface %#v is nil", v)
	} else {
		res = fmt.Sprintf("interface %#v is NOT nil", v)
	}

	require.Equal(t, "interface (*int)(nil) is NOT nil", res)

}

func TestIsNil(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name      string
		v         interface{}
		wantIsNil bool
	}{
		{"t0", 0, false},
		{"t1", nil, true},
		{"t2", error(nil), true},
		{"t2", (*int)(nil), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsNil := IsNil(tt.v); gotIsNil != tt.wantIsNil {
				t.Errorf("IsNil() = %v, want %v", gotIsNil, tt.wantIsNil)
			}
		})
	}
}
