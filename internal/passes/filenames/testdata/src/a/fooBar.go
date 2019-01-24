package a // want "file names must be lowercase"

func Test() {
	c := make(chan int, 1)
	c <- 42
	close(c)
}
