package releaselocks

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check that all locks are released when function ends`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "releaselocks",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			mutexes := mutexChecker{typeInfo: pass.TypesInfo}
			mutexes.Check(funcDecl.Body.List)
		}
	}
	return nil, nil
}

type mutexStatus int

const (
	locked mutexStatus = iota
	unlocked
)

type scope struct {
	mutexes map[string]mutexStatus
}

type mutexChecker struct {
	typeInfo *types.Info
	scopes   []scope
}

func (c *mutexChecker) Check(stmts []ast.Stmt) {
	c.scopes = append(c.scopes, scope{mutexes: make(map[string]mutexStatus)})
	for _, stmt := range stmts {
		switch v := stmt.(type) {
		case *ast.ExprStmt:
			if c.isMutexLock(stmt) {
				fmt.Println(stmt, "is lock call", getIdent(v.X))
				c.update(getIdent(v.X), locked)
			}
			if c.isMutexUnlock(stmt) {
				fmt.Println(stmt, "is unlock call", getIdent(v.X))
				c.update(getIdent(v.X), unlocked)
			}
		case *ast.DeferStmt:
			if c.isMutexDeferUnlock(stmt) {
				fmt.Println(stmt, "is defer unlock", getIdent(v.Call))
				c.update(getIdent(v.Call), unlocked)
			}
		case *ast.IfStmt:
			// TODO: Check for lock changes inside if statements
			// Approach:
			// - Return list of pairs with mutex changes
			// - Compare from different branches
			// - Only continue if all branches are equal
		case *ast.SwitchStmt:
			// TODO: Check for lock changes inside switch statements
			// Approach:
			// - Return list of pairs with mutex changes
			// - Compare from different branches
			// - Only continue if all branches are equal
		case *ast.ReturnStmt:
			leaks := c.lockedMutexes()
			if len(leaks) > 0 {
				fmt.Printf("Locked mutexes (%v) leaked in function return at %v\n", leaks, v.Pos())
			}
		}
	}
	leaks := c.lockedMutexes()
	if len(leaks) > 0 {
		fmt.Printf("Locked mutexes (%v) leaked at function end at %v\n", leaks, stmts[len(stmts)-1].End())
	}
}

func (c *mutexChecker) update(id string, val mutexStatus) {
	for i := len(c.scopes) - 1; i >= 0; i-- {
		_, ok := c.scopes[i].mutexes[id]
		if !ok {
			continue
		}
		c.scopes[i].mutexes[id] = val
		return
	}
	c.scopes[len(c.scopes)-1].mutexes[id] = val
}

func (c *mutexChecker) lockedMutexes() []string {
	var mutexes []string
	for i := len(c.scopes) - 1; i >= 0; i-- {
		for mu, v := range c.scopes[i].mutexes {
			if v != locked {
				continue
			}
			mutexes = append(mutexes, mu)
		}
	}
	return mutexes
}

func (c *mutexChecker) inspectSelectorExpr(s ast.Stmt) (*ast.SelectorExpr, types.Type, bool) {
	exprStmt, ok := s.(*ast.ExprStmt)
	if !ok {
		return nil, nil, false
	}
	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return nil, nil, false
	}
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil, false
	}
	return selExpr, c.typeInfo.TypeOf(selExpr.X), true
}

func (c *mutexChecker) isMutexLock(s ast.Stmt) bool {
	selExpr, t, ok := c.inspectSelectorExpr(s)
	if !ok {
		return false
	}
	if t.String() != "sync.Mutex" {
		return false
	}
	return selExpr.Sel.Name == "Lock"
}

func (c *mutexChecker) isMutexUnlock(s ast.Stmt) bool {
	selExpr, t, ok := c.inspectSelectorExpr(s)
	if !ok {
		return false
	}
	if t.String() != "sync.Mutex" {
		return false
	}
	return selExpr.Sel.Name == "Unlock"
}

func (c *mutexChecker) isMutexDeferUnlock(s ast.Stmt) bool {
	deferStmt, ok := s.(*ast.DeferStmt)
	if !ok {
		return false
	}
	deferExpr, ok := deferStmt.Call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	if c.typeInfo.TypeOf(deferExpr.X).String() != "sync.Mutex" {
		return false
	}
	return deferExpr.Sel.Name == "Unlock"
}

func getIdent(s ast.Expr) string {
	switch v := s.(type) {
	case *ast.CallExpr:
		return getIdent(v.Fun)
	case *ast.SelectorExpr:
		return getIdent(v.X)
	case *ast.Ident:
		return v.Name
	default:
		return "unknown ident"
	}
}
