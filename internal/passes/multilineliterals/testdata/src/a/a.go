package a

type A struct {
	B int
	C int
}

func Structs() {
	// ok
	_ = &A{}

	// ok
	_ = &A{B: 1, C: 2}

	// ok
	_ = &A{
		B: 1,
		C: 2,
	}

	// ok
	_ = &A{
		B: 1, C: 2,
	}

	// not ok
	_ = &A{
		B: 1,
		C: 2} // want "line break before closing brace"

	// not ok
	_ = &A{B: 1, // want "line break after opening brace"
		C: 2} // want "line break before closing brace"
}

func Slices() {
	// ok
	_ = []int{}

	// ok
	_ = []int{1}

	// ok
	_ = []int{1, 2}

	// ok
	_ = []int{
		1, 2,
	}

	// ok
	_ = []int{
		1,
		2,
	}

	// not ok
	_ = []int{
		// empty
	} // want "put closing brace on same line as opening brace"

	// not ok
	_ = []int{1, // want "line break after opening brace"
		2} // want "line break before closing brace"

	// not ok
	_ = []int{
		1, 2, // want "line break after each element"
		3,
		4,
		5, 6, // want "line break after each element"
	}
}
