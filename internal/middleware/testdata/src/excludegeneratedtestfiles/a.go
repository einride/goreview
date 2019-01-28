package excludegeneratedtestfiles

import (
	"fmt"

	"log" // want "remove blank line above import"
)

func A() {
	log.Println("hello")
	fmt.Printf("world")
}
