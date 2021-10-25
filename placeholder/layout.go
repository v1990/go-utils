package placeholder

import (
	"bytes"
)

const (
	txtChunk = iota
	keyChunk
)

type Layout struct {
	chunks   [][]byte
	types    []int
	capacity int
}

func ParseLayout(p PlaceHolder, content []byte) (Layout, error) {
	var x Layout
	x.capacity = len(content)

	l, r := []byte(p.left), []byte(p.right)
	ll, rl := len(l), len(r)

	pos := 0
	for {
		// find left
		n := bytes.Index(content[pos:], l)
		if n < 0 {
			x.put(txtChunk, content[pos:])
			break
		}
		// 处理转义字符
		isEscape, escaped := checkEscape(content[pos : pos+n])
		pos += n + ll
		x.put(txtChunk, escaped)

		if isEscape { // 如果是转义，把 left 补回去
			x.put(txtChunk, l)
			continue
		}

		// find right
		m := bytes.Index(content[pos:], r)
		if m < 0 {
			return Layout{}, newParseErr(pos-ll, content)
		}
		if mm := bytes.Index(content[pos:pos+m], l); mm >= 0 {
			return Layout{}, newParseErr(pos-ll, content)
		}

		// extract key
		x.put(keyChunk, content[pos:pos+m])

		pos += m + rl
	}

	return x, nil
}

func (l Layout) Keys() []string {
	dup := make(map[string]bool)
	keys := make([]string, 0)
	for i, t := range l.types {
		if t == keyChunk {
			key := string(l.chunks[i])
			if !dup[key] {
				keys = append(keys, key)
			}
		}
	}

	return keys
}
func (l Layout) Execute(ctx map[string]string) ([]byte, error) {
	var missKeys []string
	buf := bytes.NewBuffer(make([]byte, 0, l.capacity))

	for i, chunk := range l.chunks {
		switch l.types[i] {
		case txtChunk:
			buf.Write(chunk)
		case keyChunk:
			key := string(chunk)
			val, ok := ctx[key]
			if !ok {
				missKeys = append(missKeys, key)
			} else {
				buf.WriteString(val)
			}
		}

	}
	var err error
	if len(missKeys) > 0 {
		err = MissKeysError(missKeys)
	}

	return buf.Bytes(), err
}

func (l *Layout) put(typ int, chunk []byte) {
	if len(chunk) == 0 {
		return
	}
	n := len(l.types)
	if typ == txtChunk && n > 0 && l.types[n-1] == txtChunk { // 合并txt
		l.chunks[n-1] = append(l.chunks[n-1], chunk...)
	} else {
		l.chunks = append(l.chunks, chunk)
		l.types = append(l.types, typ)
	}
}

// 在b从后往前检查转义字符
// @return is 是否转义
// @return escaped 转义后的内容
func checkEscape(b []byte) (is bool, escaped []byte) {
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

func cut(b []byte, start int, length int) []byte {
	if start+length >= len(b) {
		length = len(b) - start
	}
	return b[start : start+length]
}
