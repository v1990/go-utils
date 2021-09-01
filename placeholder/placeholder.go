package placeholder

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	ErrFormat = errors.New("format err") // 格式错误，一般是缺少左/右通配符
)

var (
	DefaultPlaceholder = NewPlaceHolder("${", "}")
)

type MissKeysError []string

func (e MissKeysError) Error() string {
	return fmt.Sprintf("miss keys: %v", e.Keys())
}
func (e MissKeysError) Keys() []string {
	return []string(e)
}
func (e MissKeysError) ToError() error {
	if len(e) == 0 {
		return nil
	}
	return e
}

// PlaceHolder 占位符替换器
type PlaceHolder struct {
	left, right string
}

func NewPlaceHolder(left string, right string) *PlaceHolder {
	return &PlaceHolder{left: left, right: right}
}

// Validate 验证通配符是否成对匹配
func (p PlaceHolder) Validate(content []byte) error {
	// TODO 验证通配符是否成对匹配
	return nil
}

// ExtractKeys 提取文本中使用通配符的key
func (p PlaceHolder) ExtractKeys(content []byte) (keys []string, err error) {
	l, r := []byte(p.left), []byte(p.right)
	ll, rl := len(l), len(r)

	i := 0
	km := make(map[string]bool)
	for {
		n := bytes.Index(content[i:], l)
		if n < 0 {
			break
		}
		if isEscape, _ := p.checkEscape(content[i : i+n]); isEscape {
			continue
		}
		i += n + ll

		m := bytes.Index(content[i:], r)
		if m < 0 {
			return nil, ErrFormat
		}
		key := string(content[i : i+m])
		i += m + rl

		if !km[key] {
			keys = append(keys, key)
			km[key] = true
		}

	}

	return keys, nil
}

// 在b从后往前检查转义字符
// @return is 是否转义
// @return escaped 转义后的内容
func (p PlaceHolder) checkEscape(b []byte) (is bool, escaped []byte) {
	r := len(b)
	for i := r - 1; i >= 0; i-- {
		if b[i] != '\\' {
			break
		}
		is = !is
		if is {
			r--
		}
	}

	return is, b[:r]
}
func (p PlaceHolder) Replace(content []byte, data map[string]string) ([]byte, error) {
	if len(content) == 0 {
		return content, nil
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(content)))
	l, r := []byte(p.left), []byte(p.right)
	ll, rl := len(l), len(r)

	i := 0
	var missKeysErr MissKeysError
	for {
		// find left
		n := bytes.Index(content[i:], l)
		if n < 0 {
			buf.Write(content[i:])
			break
		}
		// 处理转义字符
		isEscape, escaped := p.checkEscape(content[i : i+n])
		i += n + ll
		buf.Write(escaped)
		// 如果是转义，把前缀补回去，然后跳过
		if isEscape {
			buf.Write(l)
			continue
		}

		// find right
		m := bytes.Index(content[i:], r)
		if m < 0 {
			buf.Write(content[i:])
			return buf.Bytes(), ErrFormat
		}
		// extract key
		key := string(content[i : i+m])
		i += m + rl

		if v, ok := data[key]; ok {
			buf.WriteString(v)
		} else {
			// write: ${key}
			buf.Write(l)
			buf.WriteString(key)
			buf.Write(r)
			missKeysErr = append(missKeysErr, key)
		}
	}

	return buf.Bytes(), missKeysErr.ToError()

}

func (p PlaceHolder) MakePlaceholder(key string) string {
	return p.left + key + p.right
}

func Replace(content []byte, data map[string]string) ([]byte, error) {
	return DefaultPlaceholder.Replace(content, data)
}

func IsMissKeysErr(err error) (MissKeysError, bool) {
	switch e := err.(type) {
	case MissKeysError:
		return e, true
	}
	return nil, false
}
