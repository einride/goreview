package multilinefunctions

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `review formatting of multi-line argument lists in function declarations`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "multilinefunction",
		Doc:      Doc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
		(*ast.CallExpr)(nil),
	}
	inspectResult.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.FuncType:
			analyzeFunctionDeclaration(pass, node.Params)
			analyzeFunctionDeclaration(pass, node.Results)
		case *ast.CallExpr:
			analyseFunctionCall(pass, node)
		}
	})
	return nil, nil
}

func analyseFunctionCall(pass *analysis.Pass, call *ast.CallExpr) {
	lastLine := pass.Fset.Position(call.Rparen).Line
	firstLine := pass.Fset.Position(call.Lparen).Line
	if firstLine == lastLine {
		return
	}
	if pass.Fset.Position(call.Args[0].Pos()).Line == firstLine {
		pass.Reportf(call.Lparen, "opening paren should be on a new line")
	}
	prevEnd := 0
	for _, a := range call.Args {
		if pass.Fset.Position(a.Pos()).Line == prevEnd {
			pass.Reportf(a.Pos(), "each argument should start on a new line")
		}
		prevEnd = pass.Fset.Position(a.End()).Line
	}
	if prevEnd == lastLine {
		pass.Reportf(call.Rparen, "closing paren should be on a new line")
	}
}

func analyzeFunctionDeclaration(pass *analysis.Pass, fields *ast.FieldList) {
	if fields == nil {
		return // no fields
	}
	openingLine := pass.Fset.Position(fields.Opening).Line
	closingLine := pass.Fset.Position(fields.Closing).Line
	if openingLine == closingLine {
		return // Not multiline
	}
	firstParamPos := fields.List[0].Type.Pos()
	firstParamLine := pass.Fset.Position(firstParamPos).Line
	if openingLine == firstParamLine {
		pass.Reportf(fields.Opening, "first field should not be on the same line as opening paren")
	}
	numParams := len(fields.List)
	lastParamPos := fields.List[numParams-1].Type.Pos()
	lastParamLine := pass.Fset.Position(lastParamPos).Line
	if closingLine == lastParamLine {
		pass.Reportf(fields.Closing, "last field should not be on the same line as closing paren")
	}

	prevParamLine := 0
	for _, result := range fields.List {
		if len(result.Names) > 1 {
			pass.Reportf(result.Names[0].NamePos, "multiline fields should declare the type for each name")
		}
		pos := result.Type.Pos()
		line := pass.Fset.Position(pos).Line
		if prevParamLine > 0 {
			if line == prevParamLine {
				pass.Reportf(pos, "multiline fields should declare one name per line")
			}
		}
		prevParamLine = line
	}
}
