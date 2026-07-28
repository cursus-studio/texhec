package deps

import (
	"fmt"
	"go/ast"
	"go/types"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Meta struct {
	ImportPath  string
	CalledFuncs []string
}

type ASTMeta struct {
	Deps           []Meta
	ThirdPartyDeps []Meta
}

func (m *Meta) String() string {
	if len(m.CalledFuncs) == 0 {
		return fmt.Sprintf("- `%s`\n", m.ImportPath)
	}
	res := &strings.Builder{}
	fmt.Fprintf(res, "`%s`:\n", m.ImportPath)
	for _, fn := range m.CalledFuncs {
		fmt.Fprintf(res, "  - `%s.%s`\n", m.ImportPath, fn)
	}
	return res.String()
}

func (d *ASTMeta) String() string {
	res := &strings.Builder{}
	if len(d.Deps) == 0 && len(d.ThirdPartyDeps) == 0 {
		return ""
	}

	res.WriteString("## Dependencies\n")
	for _, dep := range d.Deps {
		fmt.Fprintf(res, "%v\n", dep.String())
	}

	if len(d.ThirdPartyDeps) != 0 {
		res.WriteString("### Third Party\n")
	}
	for _, dep := range d.ThirdPartyDeps {
		// do not show called funcs for 3rd party packages
		dep.CalledFuncs = nil
		fmt.Fprintf(res, "%v", dep.String())
	}
	return res.String()
}

func (d *ASTMeta) ThirdPartyString() string {
	res := &strings.Builder{}
	if len(d.Deps) == 0 && len(d.ThirdPartyDeps) == 0 {
		return ""
	}

	res.WriteString("## Dependencies\n")
	if len(d.ThirdPartyDeps) != 0 {
		res.WriteString("### Third Party\n")
	}
	for _, dep := range d.ThirdPartyDeps {
		// do not show called funcs for 3rd party packages
		dep.CalledFuncs = nil
		fmt.Fprintf(res, "%v", dep.String())
	}
	return res.String()
}

func NewAST(path string) (ASTMeta, error) {
	meta := ASTMeta{}

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedModule,
		Dir:  path,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return meta, err
	}

	callsMap := make(map[string]map[string]bool)
	importedPkgs := make(map[string]*types.Package)

	pkgModules := make(map[string]*packages.Module)

	for _, pkg := range pkgs {
		for impPath, impPkg := range pkg.Imports {
			if impPkg.Module != nil {
				pkgModules[impPath] = impPkg.Module
			}
		}

		for _, imp := range pkg.Types.Imports() {
			importedPkgs[imp.Path()] = imp
		}

		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				sel, ok := n.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				obj, ok := pkg.TypesInfo.Uses[sel.Sel]
				if !ok || obj == nil || obj.Pkg() == nil {
					return true
				}

				targetPkg := obj.Pkg()

				if targetPkg.Path() == pkg.Types.Path() {
					return true
				}

				if _, exists := callsMap[targetPkg.Path()]; !exists {
					callsMap[targetPkg.Path()] = make(map[string]bool)
				}

				callsMap[targetPkg.Path()][obj.Name()] = true
				return true
			})
		}
	}

	var stdPaths []string
	var thirdPaths []string

	allDeps := make(map[string]Meta)

	for path := range importedPkgs {
		allDeps[path] = Meta{ImportPath: path, CalledFuncs: []string{}}
	}
	for path, funcs := range callsMap {
		var funcList []string
		for fn := range funcs {
			funcList = append(funcList, fn)
		}
		sort.Strings(funcList)
		allDeps[path] = Meta{ImportPath: path, CalledFuncs: funcList}
	}

	for path := range allDeps {
		if strings.Contains(path, "/internal/") || strings.Contains(path, "/internal") {
			continue
		}

		mod, hasMod := pkgModules[path]

		if !hasMod {
			continue
		}

		if mod.Main {
			stdPaths = append(stdPaths, path)
		} else {
			thirdPaths = append(thirdPaths, path)
		}
	}

	sort.Strings(stdPaths)
	sort.Strings(thirdPaths)

	for _, path := range stdPaths {
		meta.Deps = append(meta.Deps, allDeps[path])
	}
	for _, path := range thirdPaths {
		meta.ThirdPartyDeps = append(meta.ThirdPartyDeps, allDeps[path])
	}

	return meta, nil
}
