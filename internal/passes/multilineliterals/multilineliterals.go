package multilineliterals

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `review formatting of multi-line composite literals`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "multilineliterals",
		Doc:      Doc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspectResult.Preorder(nodeFilter, func(n ast.Node) {
		lit := n.(*ast.CompositeLit)
		lBrace := pass.Fset.Position(lit.Lbrace)
		rBrace := pass.Fset.Position(lit.Rbrace)
		if lBrace.Line == rBrace.Line {
			return // one-liner always OK
		}
		if len(lit.Elts) == 0 {
			if lBrace.Line != rBrace.Line {
				pass.Reportf(lit.Rbrace, "put closing brace on same line as opening brace")
			}
			return // no elements
		}
		firstElt := pass.Fset.Position(lit.Elts[0].Pos())
		if firstElt.Line == lBrace.Line {
			pass.Reportf(lit.Lbrace, "line break after opening brace")
		}
		lastElt := pass.Fset.Position(lit.Elts[len(lit.Elts)-1].Pos())
		if lastElt.Line == rBrace.Line {
			pass.Reportf(lit.Rbrace, "line break before closing brace")
		}
		// check if all elements are on same line
		areAllOnSameLine := true
		for i, prevElt := 1, firstElt; i < len(lit.Elts); i++ {
			elt := pass.Fset.Position(lit.Elts[i].Pos())
			if elt.Line != prevElt.Line {
				areAllOnSameLine = false
				break
			}
			prevElt = elt
		}
		// check each element
		for i, prevElt := 1, firstElt; i < len(lit.Elts); i++ {
			elt := pass.Fset.Position(lit.Elts[i].Pos())
			if elt.Line == prevElt.Line && !areAllOnSameLine {
				pass.Reportf(lit.Elts[i].Pos(), "line break after each element")
			}
			prevElt = elt
		}
	})
	return nil, nil
}
