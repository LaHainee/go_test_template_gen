package field

import (
	"fmt"
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

func GetNames(field *ast.Field) []string {
	names := make([]string, 0)

	if len(field.Names) == 0 {
		return []string{}
	}

	for _, name := range field.Names {
		if name == nil {
			continue
		}

		names = append(names, name.Name)
	}

	return names
}

func GetPackage(field *ast.Field) *string {
	switch t := field.Type.(type) {
	case *ast.SelectorExpr:
		return pointer.To(fmt.Sprintf("%s", t.X))
	}

	return nil
}

func GetType(field *ast.Field) string {
	switch t := field.Type.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		elemType := GetType(&ast.Field{Type: t.Elt})
		return "[]" + elemType
	case *ast.StarExpr:
		pointedType := GetType(&ast.Field{Type: t.X})
		return "*" + pointedType
	case *ast.MapType:
		keyType := GetType(&ast.Field{Type: t.Key})
		valueType := GetType(&ast.Field{Type: t.Value})
		return fmt.Sprintf("map[%s]%s", keyType, valueType)
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", t.X, t.Sel.Name)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return ""
	}
}
