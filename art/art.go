package art

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"sort"
	"strings"

	"github.com/orayew2002/art/utils"
)

func CleanStruct(filePath string) {
	fileBody, err := os.ReadFile(filePath)
	plog("can't read file", err)

	fset := token.NewFileSet()
	astTree, err := parser.ParseFile(fset, "", fileBody, parser.AllErrors|parser.ParseComments)
	plog("can't parse file ast", err)

	ast.Inspect(astTree, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		sizeOfType := func(expr ast.Expr) uintptr {
			typeName := exprToString(fset, expr)
			cleanName := strings.TrimPrefix(typeName, "*")
			t := utils.BasicTypeFromName(cleanName)
			if t == nil {
				return 0
			}
			return t.Size()
		}

		sort.SliceStable(st.Fields.List, func(i, j int) bool {
			return sizeOfType(st.Fields.List[i].Type) > sizeOfType(st.Fields.List[j].Type)
		})

		newFields := make([]*ast.Field, len(st.Fields.List))
		for i, f := range st.Fields.List {

			newNames := make([]*ast.Ident, len(f.Names))
			for j, name := range f.Names {
				newNames[j] = ast.NewIdent(name.Name)
			}

			newFields[i] = &ast.Field{
				Doc:     f.Doc,
				Names:   newNames,
				Type:    f.Type,
				Tag:     f.Tag,
				Comment: f.Comment,
			}
		}
		st.Fields.List = newFields

		return false
	})

	var buf bytes.Buffer
	plog("error formatting code", format.Node(&buf, fset, astTree))
	plog("error remove old file", os.Remove(filePath))
	plog("error writing file", os.WriteFile(filePath, buf.Bytes(), 0644))
}

func plog(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s: %+v", msg, err))
	}
}

func exprToString(fset *token.FileSet, expr ast.Expr) string {
	if expr == nil {
		return ""
	}
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return ""
	}
	return buf.String()
}
