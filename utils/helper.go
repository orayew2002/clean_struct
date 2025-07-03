package utils

import (
	"fmt"
	"go/ast"
	"reflect"
)

func BasicTypeFromName(name string) reflect.Type {
	switch name {
	case "int":
		return reflect.TypeOf(int(0))
	case "int8":
		return reflect.TypeOf(int8(0))
	case "int16":
		return reflect.TypeOf(int16(0))
	case "int32":
		return reflect.TypeOf(int32(0))
	case "int64":
		return reflect.TypeOf(int64(0))
	case "uint":
		return reflect.TypeOf(uint(0))
	case "uint8":
		return reflect.TypeOf(uint8(0))
	case "uint16":
		return reflect.TypeOf(uint16(0))
	case "uint32":
		return reflect.TypeOf(uint32(0))
	case "uint64":
		return reflect.TypeOf(uint64(0))
	case "string":
		return reflect.TypeOf("")
	case "bool":
		return reflect.TypeOf(true)
	case "float32":
		return reflect.TypeOf(float32(0))
	case "float64":
		return reflect.TypeOf(float64(0))
	default:
		return nil
	}
}

func SafeSize(t reflect.Type) uintptr {
	if t == nil {
		return 0
	}
	return t.Size()
}

func ExprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.ArrayType:
		return "[]" + ExprToString(e.Elt)
	case *ast.StarExpr:
		return "*" + ExprToString(e.X)
	case *ast.SelectorExpr:
		return ExprToString(e.X) + "." + e.Sel.Name
	default:
		return fmt.Sprintf("%T", expr)
	}
}
