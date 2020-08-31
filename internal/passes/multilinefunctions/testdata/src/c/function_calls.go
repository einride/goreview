package function_declaration_parameters

type A struct {
	a *A
}

func a(A, A) {}

func foo() {
	a(A{}, A{}) // ok
	a(A{
		a: nil,
	}, A{}) // ok
	a(A{
		a: nil,
	}, A{
		a: nil,
	}) // ok
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
	a(
		A{},
		A{}) // want `closing paren should be on a new line`
	a(A{
		a: nil,
	},
		A{}, // want `must either have all arguments on individual lines or no linebreaks before or after arguments`
	)
	a(A{},
		A{}, // want `must either have all arguments on individual lines or no linebreaks before or after arguments`
	)
	a(A{
		a: nil,
	}, A{},
	) // want `closing paren should be on the same line as last argument`
}
