package function_declaration_parameters

type A struct{}

// ok
func a() (a, b, c int) {
	return 0, 0, 0
}

// ok
func (A) a() (a, b, c int) {
	return 0, 0, 0
}

// ok
func b() (
	a int,
	b int,
	c int,
) {
	return 0, 0, 0
}

// ok
func (A) b() (
	a int,
	b int,
	c int,
) {
	return 0, 0, 0
}

func bad_1() (
	a, // want `multiline fields should declare the type for each name`
	b int,
) {
	return 0, 0
}

func (A) bad_1() (
	a, // want `multiline fields should declare the type for each name`
	b int,
) {
	return 0, 0
}

func bad_2() (
	a, b, c int, // want `multiline fields should declare the type for each name`
) {
	return 0, 0, 0
}

func (A) bad_2() (
	a, b, c int, // want `multiline fields should declare the type for each name`
) {
	return 0, 0, 0
}

func bad_3() (
	a int, b int, // want `multiline fields should declare one name per line`
) {
	return 0, 0
}

func (A) bad_3() (
	a int, b int, // want `multiline fields should declare one name per line`
) {
	return 0, 0
}

func bad_4() (a int, // want `first field should not be on the same line as opening paren`
	b int,
	c int,
) {
	return 0, 0, 0
}

func (A) bad_4() (a int, // want `first field should not be on the same line as opening paren`
	b int,
	c int,
) {
	return 0, 0, 0
}

func bad_5() (
	a int,
	b int,
	c int) { // want `last field should not be on the same line as closing paren`
	return 0, 0, 0
}

func (A) bad_5() (
	a int,
	b int,
	c int) { // want `last field should not be on the same line as closing paren`
	return 0, 0, 0
}

func bad_6() (
	int, int, // want `multiline fields should declare one name per line`
) {
	return 0, 0
}

func (A) bad_6() (
	int, int, // want `multiline fields should declare one name per line`
) {
	return 0, 0
}

func bad_7() (int, // want `first field should not be on the same line as opening paren`
	int,
	int,
) {
	return 0, 0, 0
}

func (A) bad_7() (int, // want `first field should not be on the same line as opening paren`
	int,
	int,
) {
	return 0, 0, 0
}

func bad_8() (
	int,
	int,
	int) { // want `last field should not be on the same line as closing paren`
	return 0, 0, 0
}

func (A) bad_8() (
	int,
	int,
	int) { // want `last field should not be on the same line as closing paren`
	return 0, 0, 0
}
