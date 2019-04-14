package a

import "github.com/pkg/errors" // want `use "golang.org/x/xerrors" instead of "github.com/pkg/errors"`

func IllegalImports() {
	_ = errors.New("hey")
}
