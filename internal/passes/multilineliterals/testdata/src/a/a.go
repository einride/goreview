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
	_ = []A{}

	// ok
	_ = []A{{}}

	// ok
	_ = []A{{}, {}}

	// ok
	_ = []A{
		{}, {},
	}

	// ok
	_ = []A{
		{},
		{},
	}

	// not ok
	_ = []A{
		// empty
	} // want "put closing brace on same line as opening brace"

	// not ok
	_ = []A{{}, // want "line break after opening brace"
		{}} // want "line break before closing brace"

	// not ok
	_ = []A{
		{}, {}, // want "line break after each element"
		{},
		{},
		{}, {}, // want "line break after each element"
	}
}
