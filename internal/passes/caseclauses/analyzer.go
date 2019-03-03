package caseclauses

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `review formatting of case clauses`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "caseclauses",
		Doc:      Doc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CaseClause)(nil),
	}
	inspectResult.Preorder(nodeFilter, func(n ast.Node) {
		caseClause := n.(*ast.CaseClause)
		if len(caseClause.List) < 2 {
			return // ok
		}
		isRow := true
		isCol := true
		prevLine := pass.Fset.Position(caseClause.List[0].Pos()).Line
		for _, e := range caseClause.List[1:] {
			currLine := pass.Fset.Position(e.Pos()).Line
			if currLine != prevLine {
				isRow = false
			}
			if currLine != prevLine+1 {
				isCol = false
			}
			if !isRow && !isCol {
				pass.Reportf(e.Pos(), "switch clauses must be on a single row or column")
				return
			}
			prevLine = currLine
		}
	})
	return nil, nil
}
