package labels

import (
	"fmt"
	"go/ast"

	"go.einride.tech/review/internal/lettercase"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `review labels`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "labels",
		Doc:      Doc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, fmt.Errorf("labels inspector type cast failed")
	}
	nodeFilter := []ast.Node{
		(*ast.LabeledStmt)(nil),
	}
	inspectResult.Preorder(nodeFilter, func(n ast.Node) {
		labeledStmt, ok := n.(*ast.LabeledStmt)
		if !ok {
			return
		}
		if !lettercase.IsCamel(labeledStmt.Label.Name) {
			pass.Reportf(labeledStmt.Pos(), "labels must use CamelCase")
		}
	})
	return nil, nil
}
