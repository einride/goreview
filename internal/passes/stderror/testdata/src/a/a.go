package a

import "errors"

func Imports() {
	_ = errors.New("hey") // ok
}
