package a

//go:cgo_

import (
	"fmt"
	"log"
)

//go:generate echo go generate comments are OK!
//go:embed echo go embed commentts are OK!
//nolint echo nolint comments are OK!
//nolint:specifilinter echo specific linter comments are OK!

func Imports() {
	// good comment
	//
	// also good
	//
	//	still a valid comment (with a tab)
	log.Println("hello")
	// want "comments must start with '// '"
	//bad comment
	fmt.Println("world")
}
