package a

import (
	"golang.org/x/xerrors" // want `use "github.com/pkg/errors" instead of "golang.org/x/xerrors"`
)

func IllegalImports() {
	_ = xerrors.New("hey")
}
