package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func branchCount(fn *ast.FuncDecl) uint {
	var count uint
	ast.Inspect(fn, func (node ast.Node) bool {
		_, nfor := node.(*ast.ForStmt)
		if nfor {
			count += 1
			return true
		}
		_, nif := node.(*ast.IfStmt)
		if nif {
			count += 1
			return true
		}
		_, nrange := node.(*ast.RangeStmt)
		if nrange {
			count += 1
			return true
		}
		_, nswitch := node.(*ast.SwitchStmt)
                if nswitch {
                        count += 1
                        return true
                }
		_, ntypeswitch := node.(*ast.TypeSwitchStmt)
                if ntypeswitch {
                        count += 1
                        return true
                }
		_, nbranch := node.(*ast.BranchStmt)
                if nbranch {
                        count += 1
                        return true
                }
		return true
	})
	return count
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
