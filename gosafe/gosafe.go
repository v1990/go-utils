package gosafe

import (
	"fmt"
	"log"
)

var recoverHandler = func(r interface{}) {
	log.Println("panic:", r)
}

// SetRecoverHandler ...
func SetRecoverHandler(f func(interface{})) {
	recoverHandler = f
}

// Go safe
func Go(f func()) {
	defer func() {
		if r := recover(); r != nil {
			recoverHandler(r)
		}
	}()

	f()
}

// Recover recover
func Recover() error {
	r := recover()
	if r == nil {
		return nil
	}

	switch r := r.(type) {
	case error:
		return r
	default:
		return fmt.Errorf("%v", r)
	}
}
