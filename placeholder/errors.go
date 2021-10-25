package placeholder

import (
	"errors"
	"fmt"
)

var (
	ErrFormat = errors.New("format err") // 格式错误，一般是缺少左/右通配符
)

type MissKeysError []string

func (e MissKeysError) Error() string {
	return fmt.Sprintf("miss keys: %v", e.Keys())
}
func (e MissKeysError) Keys() []string { return []string(e) }

func IsMissKeysErr(err error) (MissKeysError, bool) {
	switch e := err.(type) {
	case MissKeysError:
		if len(e) == 0 {
			return nil, false
		}
		return e, true
	}
	return nil, false
}

type ParseError struct {
	pos int
	txt []byte
	msg string
}

func newParseErr(pos int, content []byte) ParseError {
	return ParseError{
		pos: pos,
		txt: cut(content, pos, 20),
	}
}
func (p ParseError) Error() string {
	return fmt.Sprintf("parse err(pos: %d): %s ", p.pos, p.txt)
}
