package types

import "go/ast"

func GetStruct(typ ast.Expr, file *ast.File) *ast.StructType {
	typeSpec := getTypeSpec(typ, file)
	if typeSpec == nil {
		return nil
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil
	}

	return structType
}

func GetInterface(typ ast.Expr, file *ast.File) *ast.InterfaceType {
	typeSpec := getTypeSpec(typ, file)
	if typeSpec == nil {
		return nil
	}

	interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
	if !ok {
		return nil
	}

	return interfaceType
}

func getTypeSpec(typ ast.Expr, file *ast.File) *ast.TypeSpec {
	switch t := typ.(type) {
	case *ast.Ident:
		return lookup(t.Name, file)
	case *ast.StarExpr:
		ident, ok := t.X.(*ast.Ident)
		if !ok {
			return nil
		}

		return lookup(ident.Name, file)
	}

	return nil
}

func lookup(name string, file *ast.File) *ast.TypeSpec {
	object := file.Scope.Lookup(name)

	if object == nil || object.Kind != ast.Typ {
		return nil
	}

	typeSpec, ok := object.Decl.(*ast.TypeSpec)
	if !ok {
		return nil
	}

	return typeSpec
}
