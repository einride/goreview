package releaselocks_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"go.einride.tech/review/internal/passes/releaselocks"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, releaselocks.Analyzer(), "a")
}

func TestIsMutex(t *testing.T) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testdata/src/a/clean.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal("failed to parse file", err)
	}
	for _, decl := range node.Decls {
		fDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		for _, stmt := range fDecl.Body.List {
			fmt.Println(tostringStmt(stmt))
		}
	}
}

func tostringStmt(s ast.Stmt) string {
	switch v := s.(type) {
	case *ast.DeclStmt:
		return fmt.Sprintf("&DeclStmt{Decl: %s}", tostringDecl(v.Decl))
	case *ast.ExprStmt:
		return fmt.Sprintf("&ExprStmt{X: %s}", tostringExpr(v.X))
	case *ast.DeferStmt:
		return fmt.Sprintf("&DeferStmt{Call: %s}", tostringExpr(v.Call))
	default:
		return fmt.Sprintf("%#v (unimplemented stmt)", s)
	}
}

func tostringDecl(s ast.Decl) string {
	switch v := s.(type) {
	case *ast.FuncDecl:
		return fmt.Sprintf("&FuncDecl{Name: %s, Recv: %v}", v.Name.Name, v.Recv)
	case *ast.GenDecl:
		return fmt.Sprintf("&GenDecl{Tok: %s,Spec: %s}", v.Tok.String(), v.Specs)
	default:
		return fmt.Sprintf("%#v (unimplemented decl)", s)
	}
}

func tostringExpr(s ast.Expr) string {
	switch v := s.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("&SelectorExpr{Sel: %s, X: %s}", v.Sel.Name, tostringExpr(v.X))
	case *ast.CallExpr:
		return fmt.Sprintf("&CallExpr{Fun: %s}", tostringExpr(v.Fun))
	case *ast.Ident:
		return fmt.Sprintf("&Ident{Name: %s}", v.Name)
	default:
		return fmt.Sprintf("%#v (unimplemented expr)", s)
	}
}
