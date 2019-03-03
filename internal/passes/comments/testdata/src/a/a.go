package a

import (
	"fmt"
	"log"
)

//go:generate echo go generate comments are OK!

func Imports() {
	// good comment
	//
	// also good
	log.Println("hello")
	// want "comments must start with '// '"
	//bad comment
	fmt.Println("world")
}
