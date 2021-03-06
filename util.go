package caplit

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type CapnpStructParams struct {
	Name     string
	Template string
}

type CapnpStruct struct {
	Name   string
	Path   string
	Keys   []CapnpStructParams
	Parent string
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
		switch sel := e.Sel; e.X.(type) {
		case *ast.Ident:
			// Name is not neccesary ( shared.MagicStat -> MagicStat )
			return sel.Name
		case *ast.SelectorExpr:
			panic("do not enter this phase")
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
	case *ast.ArrayType:
		switch es := e.Elt.(type) {
		case *ast.Ident:
			return "[]" + es.Name
		default:
			panic(fmt.Sprintf("unknown method receiver AST node type %T", es))

		}
	case *ast.SelectorExpr:
		switch sel := e.Sel; e.X.(type) {
		case *ast.Ident:
			return sel.Name
		default:
			panic(fmt.Sprintf("unknown method receiver AST node type %T", e.X))
		}
	}
	panic(fmt.Sprintf("unknown method receiver AST node type %T", fn.Type.Results.List[0].Type))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type pathSourcePair struct {
	path string
	file *ast.File
}
type parsedSources []pathSourcePair

func parseCapnpSources(filePattern string) parsedSources {
	parsedFiles := make(parsedSources, 0)
	fset := token.NewFileSet()
	matches, err := filepath.Glob(filePattern)
	check(err)
	for _, path := range matches {
		if strings.Contains(path, "generated") {
			continue
		}

		parsed, err := parser.ParseFile(fset, path, nil, 0)
		check(err)

		parsedFiles = append(parsedFiles, pathSourcePair{path: path, file: parsed})
	}
	return parsedFiles
}

func (files parsedSources) FilterFuncDecl(filter func(*ast.FuncDecl) bool) []CapnpFuncDecl {
	funcDecls := make([]CapnpFuncDecl, 0)
	for _, pair := range files {
		for _, d := range pair.file.Decls {
			switch t := d.(type) {
			case *ast.FuncDecl:
				if filter(t) {
					funcDecls = append(funcDecls, CapnpFuncDecl{t, strings.TrimSuffix(pair.path, ".go")})
				}
			}
		}
	}
	return funcDecls
}

func CapnpStructs(capnpDir string) []CapnpStruct {
	funcDecls := parseCapnpSources(capnpDir + "/*.capnp.go").FilterFuncDecl(func(t *ast.FuncDecl) bool { return strings.HasPrefix(t.Name.Name, "NewRoot") })
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
	funcDecls := parseCapnpSources(capnpDir + "/*.capnp.go").FilterFuncDecl(func(t *ast.FuncDecl) bool { return strings.HasSuffix(t.Name.Name, "FromString") })
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
