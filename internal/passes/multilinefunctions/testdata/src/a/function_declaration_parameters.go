package function_declaration_parameters

type A struct{}

// ok
func a(a, b, c int) {}

// ok
func (A) a(a, b, c int) {}

// ok
func b(
	a int,
	b int,
	c int,
) {
}

// ok
func (A) b(
	a int,
	b int,
	c int,
) {
}

func bad_1(
	a, // want `multiline fields should declare the type for each name`
	b int,
) {
}

func (A) bad_1(
	a, // want `multiline fields should declare the type for each name`
	b int,
) {
}

func bad_2(
	a, b, c int, // want `multiline fields should declare the type for each name`
) {
}

func (A) bad_2(
	a, b, c int, // want `multiline fields should declare the type for each name`
) {
}

func bad_3(
	a int, b int, // want `multiline fields should declare one name per line`
) {
}

func (A) bad_3(
	a int, b int, // want `multiline fields should declare one name per line`
) {
}

func bad_4(a int, // want `first field should not be on the same line as opening paren`
	b int,
	c int,
) {
}

func (A) bad_4(a int, // want `first field should not be on the same line as opening paren`
	b int,
	c int,
) {
}

func bad_5(
	a int,
	b int,
	c int) { // want `last field should not be on the same line as closing paren`
}

func (A) bad_5(
	a int,
	b int,
	c int) { // want `last field should not be on the same line as closing paren`
}
