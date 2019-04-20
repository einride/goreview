package function_declaration_parameters

type A struct {
	a *A
}

func a(A, A) {}

func foo() {
	a(A{}, A{}) // ok
	a(
		A{},
		A{},
	) // ok
	a(
		A{
			a: nil,
		}, A{}, // want `each argument should start on a new line`
	)
	a(
		A{}, A{}, // want `each argument should start on a new line`
	)
	a(A{}, // want `opening paren should be on a new line`
		A{},
	)
	a(
		A{},
		A{}) // want `closing paren should be on a new line`
}
