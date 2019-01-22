package a

// no imports, no problems

func NoImports() {
	c := make(chan int, 1)
	c <- 42
	close(c)
}
