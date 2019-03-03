package a

func A() {
	var value int

	// ok
	switch value {
	case 0, 1, 2:
	case 3,
		4,
		5:
	}

	// not ok
	switch value {
	case 0, 1,
		2: // want "switch clauses must be on a single row or column"
	case 3,
		4, 5: // want "switch clauses must be on a single row or column"
	}
}
