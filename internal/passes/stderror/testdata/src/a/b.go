package a

import (
	"golang.org/x/xerrors" // want `use package "errors" and "fmt.Errorf" instead of "golang.org/x/xerrors"`
)

func IllegalImports() {
	_ = xerrors.New("hey")
}
