package caplit

import (
	"go/ast"
	"fmt"
	"unicode"
	"go/token"
	"path/filepath"
	"strings"
	"go/parser"
	"text/template"
	"os"
	"os/exec"
)

type CapnpStructParams struct {
	Name     string
	Template string
}

type CapnpStruct struct {
	Name string
	Path string
	Keys []CapnpStructParams
}

type CapnpFuncDecl struct {
	FuncDecl *ast.FuncDecl
	Path     string
}

// return the name of param type from FunDecl
func paramType(fn *ast.FuncDecl) string {
	if fn.Type.Params == nil || len(fn.Type.Params.List) == 0 {
		return ""
	}

	switch e := fn.Type.Params.List[0].Type.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		switch es := e.X.(type) {
		case *ast.Ident:
			return es.Name
		case *ast.SelectorExpr:
			return es.X.(*ast.Ident).Name + "." + es.Sel.String()
		default:
			panic(fmt.Sprintf("unknown param type %T", e.X))
		}
	case *ast.SelectorExpr:
		switch sel := e.Sel; es := e.X.(type) {
		case *ast.Ident:
			// Name is not neccesary ( shared.MagicStat -> MagicStat )
			return sel.Name
		case *ast.SelectorExpr:
			panic("do not enter this phase")
			return "[]" + es.X.(*ast.Ident).Name + "." + es.Sel.String()
		default:
			panic(fmt.Sprintf("unknown param type %T", e.X))
		}
	case *ast.ArrayType:
		switch es := e.Elt.(type) {
		case *ast.Ident:
			return "[]" + es.Name
		default:
			panic(fmt.Sprintf("unknown param type %T", es))
		}
	}
	panic(fmt.Sprintf("unknown method receiver AST node type %T", fn.Type.Params.List[0].Type))
}

// return the name of type of receiver from FunDecl
func receiverType(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}

	switch e := fn.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return e.X.(*ast.Ident).Name
	}
	fmt.Println("fn ", fn)
	panic(fmt.Sprintf("unknown method receiver AST node type %T", fn.Recv.List[0].Type))
}

// return the return type from FunDecl
func returnType(fn *ast.FuncDecl) string {
	if fn.Type.Results == nil || len(fn.Type.Results.List) == 0 {
		return ""
	}

	switch e := fn.Type.Results.List[0].Type.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return e.X.(*ast.Ident).Name
	}
	panic(fmt.Sprintf("unknown method receiver AST node type %T", fn.Recv.List[0].Type))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseCapnpFuncDecl(capnpDir string, filter func(*ast.FuncDecl) bool) []CapnpFuncDecl {
	funcDecls := make([]CapnpFuncDecl, 0)
	fset := token.NewFileSet()
	matches, err := filepath.Glob(capnpDir)
	check(err)
	for _, path := range matches {
		if strings.Contains(path, "generated") {
			continue
		}

		f, err := parser.ParseFile(fset, path, nil, 0)
		check(err)

		for _, d := range f.Decls {
			switch t := d.(type) {
			case *ast.FuncDecl:
				if filter(t) {
					funcDecls = append(funcDecls, CapnpFuncDecl{t, strings.TrimSuffix(path, ".go")})
				}
			}
		}
	}
	return funcDecls
}

func CapnpStructs(capnpDir string) []CapnpStruct {
	funcDecls := parseCapnpFuncDecl(capnpDir+"/*.capnp.go", func(t *ast.FuncDecl) bool { return strings.HasPrefix(t.Name.Name, "NewRoot") })
	capnpStructs := make([]CapnpStruct, 0)
	for _, funcDecl := range funcDecls {
		capnpStruct := CapnpStruct{
			Name: returnType(funcDecl.FuncDecl),
			Path: funcDecl.Path,
		}
		capnpStructs = append(capnpStructs, capnpStruct)
	}

	return capnpStructs
}

func firstToLower(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func elemInList(elem string, l []string) bool {
	for _, s := range l {
		if s == elem {
			return true
		}
	}
	return false
}

func generate(tmplStr, filepath string, params map[string]interface{}) {
	tmpl, err := template.New("capnp").Parse(tmplStr)
	check(err)

	// generate
	{
		file, err := os.Create(filepath)
		check(err)
		defer file.Close()
		err = tmpl.Execute(file, params)
		check(err)
	}

	// do formatting
	cmd := exec.Command("gofmt", "-w", filepath)
	err = cmd.Run()
	check(err)
}

func CapnpEnums(capnpDir string) []CapnpStruct {
	funcDecls := parseCapnpFuncDecl(capnpDir+"/*.capnp.go", func(t *ast.FuncDecl) bool { return strings.HasSuffix(t.Name.Name, "FromString") })
	capnpStructs := make([]CapnpStruct, 0)
	for _, funcDecl := range funcDecls {
		capnpStruct := CapnpStruct{
			Name: returnType(funcDecl.FuncDecl),
			Path: funcDecl.Path,
		}
		capnpStructs = append(capnpStructs, capnpStruct)
	}

	return capnpStructs
}

// return list of name of Enums in targetDir
func GetEnumList(targetDir string) []string {
	enums := CapnpEnums(targetDir)
	ret_val := make([]string, len(enums))
	for i, enumStruct := range enums {
		ret_val[i] = enumStruct.Name
	}
	return ret_val
}
