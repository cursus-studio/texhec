package types

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type (
	Meta struct {
		Name,
		Type,
		Comment string
	}
	StructMeta struct {
		Meta
		Properties []Meta
		Methods    []Meta
	}
	InterfaceMeta struct {
		Meta
		Methods []Meta
	}
	ASTMeta struct {
		Vars       []Meta
		Funcs      []Meta
		Structs    []StructMeta
		Interfaces []InterfaceMeta
	}
)

func (m *Meta) String() string { return fmt.Sprintf("%v\nType: `%v`\n%v", m.Name, m.Type, m.Comment) }
func (s *ASTMeta) String() string {
	res := &strings.Builder{}
	if len(s.Interfaces) != 0 || len(s.Structs) != 0 {
		res.WriteString("## Types\n")
	}
	for _, interfaceMeta := range s.Interfaces {
		fmt.Fprintf(res, "### type %v\n", interfaceMeta.String())
		for _, interfaceMethodMeta := range interfaceMeta.Methods {
			fmt.Fprintf(res, "#### method %v %v\n", interfaceMeta.Name, interfaceMethodMeta.String())
		}
	}
	for _, structMeta := range s.Structs {
		fmt.Fprintf(res, "### type %v\n", structMeta.String())
		for _, interfacePropertyMeta := range structMeta.Properties {
			fmt.Fprintf(res, "#### property %v %v\n", structMeta.Name, interfacePropertyMeta.String())
		}
		for _, interfaceMethodMeta := range structMeta.Methods {
			fmt.Fprintf(res, "#### method %v %v\n", structMeta.Name, interfaceMethodMeta.String())
		}
	}

	if len(s.Vars) != 0 {
		res.WriteString("## Variables\n")
	}
	for _, varMeta := range s.Vars {
		fmt.Fprintf(res, "### var %v\n", varMeta.String())
	}

	if len(s.Funcs) != 0 {
		res.WriteString("## Functions\n")
	}
	for _, funcMeta := range s.Funcs {
		fmt.Fprintf(res, "### func %v\n", funcMeta.String())
	}

	return res.String()
}

//

func fetchInterfaceMethodComment(pkg *packages.Package, methodName string) string {
	for _, file := range pkg.Syntax {
		var comment string
		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			iface, ok := ts.Type.(*ast.InterfaceType)
			if !ok || iface.Methods == nil {
				return true
			}
			for _, field := range iface.Methods.List {
				for _, name := range field.Names {
					if name.Name != methodName || field.Doc == nil {
						continue
					}
					comment = field.Doc.Text()
					return false
				}
			}
			return true
		})
		if comment != "" {
			return comment
		}
	}
	return ""
}
func fetchStructMethodComment(pkg *packages.Package, methodName string) string {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			fd, ok := decl.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || fd.Name.Name != methodName || fd.Doc == nil {
				continue
			}
			return fd.Doc.Text()
		}
	}
	return ""
}
func fetchStructPropertyComment(pkg *packages.Package, propertyName string) string {
	var fComment string
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}
			for _, field := range st.Fields.List {
				for _, name := range field.Names {
					if name.Name == propertyName {
						if field.Doc != nil {
							fComment = field.Doc.Text()
						} else if field.Comment != nil {
							fComment = field.Comment.Text()
						}
						return false
					}
				}
			}
			return true
		})
	}
	return fComment
}

func NewAST(path string) (ASTMeta, error) {
	astMeta := ASTMeta{}
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  path,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return astMeta, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				gd, ok := n.(*ast.GenDecl)
				if !ok {
					return true
				}

				var comment string
				if gd.Doc != nil {
					comment = gd.Doc.Text()
				}

				for _, spec := range gd.Specs {
					switch s := spec.(type) {
					case *ast.ValueSpec:
						for _, name := range s.Names {
							if !name.IsExported() {
								continue
							}
							obj := pkg.TypesInfo.Defs[name]
							if obj == nil {
								continue
							}

							var specComment string
							if s.Doc != nil {
								specComment = s.Doc.Text()
							} else {
								specComment = comment
							}

							astMeta.Vars = append(astMeta.Vars, Meta{
								Name:    obj.Name(),
								Type:    obj.Type().String(),
								Comment: specComment,
							})
						}

					case *ast.TypeSpec:
						if !s.Name.IsExported() {
							continue
						}
						obj := pkg.TypesInfo.Defs[s.Name]
						if obj == nil {
							continue
						}

						var specComment string
						if s.Doc != nil {
							specComment = s.Doc.Text()
						} else {
							specComment = comment
						}

						t := obj.Type()

						if iface, ok := t.Underlying().(*types.Interface); ok {
							iMeta := InterfaceMeta{
								Meta: Meta{
									Name:    obj.Name(),
									Type:    t.String(),
									Comment: specComment,
								},
							}

							for m := range iface.Methods() {
								if !m.Exported() {
									continue
								}
								mComment := fetchInterfaceMethodComment(pkg, m.Name())
								iMeta.Methods = append(iMeta.Methods, Meta{
									Name:    m.Name(),
									Type:    m.Type().String(),
									Comment: mComment,
								})
							}

							if obj.Name() == "Service" {
								astMeta.Interfaces = append([]InterfaceMeta{iMeta}, astMeta.Interfaces...)
							} else {
								astMeta.Interfaces = append(astMeta.Interfaces, iMeta)
							}

						} else if named, ok := t.(*types.Named); ok {
							sMeta := StructMeta{
								Meta: Meta{
									Name:    obj.Name(),
									Type:    t.String(),
									Comment: specComment,
								},
							}

							if strct, ok := named.Underlying().(*types.Struct); ok {
								for f := range strct.Fields() {
									if !f.Exported() {
										continue
									}
									sMeta.Properties = append(sMeta.Properties, Meta{
										Name:    f.Name(),
										Type:    f.Type().String(),
										Comment: fetchStructPropertyComment(pkg, f.Name()),
									})
								}
							}

							for m := range named.Methods() {
								if !m.Exported() {
									continue
								}
								mComment := fetchStructMethodComment(pkg, m.Name())
								sMeta.Methods = append(sMeta.Methods, Meta{
									Name:    m.Name(),
									Type:    m.Type().String(),
									Comment: mComment,
								})
							}

							astMeta.Structs = append(astMeta.Structs, sMeta)
						}
					}
				}
				return true
			})

			for _, decl := range file.Decls {
				fd, ok := decl.(*ast.FuncDecl)
				if !ok || !fd.Name.IsExported() || fd.Recv != nil {
					continue
				}
				obj := pkg.TypesInfo.Defs[fd.Name]
				if obj == nil {
					continue
				}

				var funcComment string
				if fd.Doc != nil {
					funcComment = fd.Doc.Text()
				}

				astMeta.Funcs = append(astMeta.Funcs, Meta{
					Name:    obj.Name(),
					Type:    obj.Type().String(),
					Comment: funcComment,
				})
			}
		}
	}
	return astMeta, nil
}
