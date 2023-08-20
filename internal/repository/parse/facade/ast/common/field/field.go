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
	case *ast.FuncType:
		params := ""
		results := ""

		if t.Params != nil && len(t.Params.List) > 0 {
			for i, param := range t.Params.List {
				paramType := GetType(&ast.Field{Type: param.Type})
				if i > 0 {
					params += ", "
				}
				params += paramType
			}
		}

		if t.Results != nil && len(t.Results.List) > 0 {
			for i, result := range t.Results.List {
				resultType := GetType(&ast.Field{Type: result.Type})
				if i > 0 {
					results += ", "
				}
				results += resultType
			}
		}

		if len(results) == 0 {
			return fmt.Sprintf("func(%s)", params)
		}
		return fmt.Sprintf("func(%s) (%s)", params, results)
	default:
		return ""
	}
}
