package comments

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check that comments follow conventions`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "comments",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, cg := range f.Comments {
			for i, g := range cg.List {
				if g.Text == "//" {
					// ignore empty comments
					continue
				}
				switch {
				case strings.HasPrefix(g.Text, "//go:generate"), // ignore go:generate comments
					strings.HasPrefix(g.Text, "//go:embed"), // ignore go:embed comments
					strings.HasPrefix(g.Text, "//nolint"):   // ignore nolint comments
					continue
				}
				if !strings.HasPrefix(g.Text, "// ") {
					pos := g.Slash
					// special case for tests, since tests use comments for assertions
					// expect the assertion to be located directly above the failing comment
					if i > 0 && cg.List[i-1].Text == `// want "comments must start with '// '"` {
						pos = cg.List[i-1].Slash
					}
					pass.Reportf(pos, "comments must start with '// '")
				}
			}
		}
	}
	return nil, nil
}
