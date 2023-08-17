package structure

import (
	"go/ast"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/ast/types"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/field"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/interfaces"
)

func Get(f *ast.Field, file *ast.File) *model.Structure {
	structType := types.GetStruct(f.Type, file)
	if structType == nil {
		return nil
	}

	structureName := field.GetType(f)
	structureName = strings.TrimLeft(structureName, "*")

	return &model.Structure{
		Name:         structureName,
		Dependencies: getDependencies(structType, file),
	}
}

func getDependencies(structType *ast.StructType, file *ast.File) []model.Dependency {
	dependencies := make([]model.Dependency, 0)

	if structType.Fields == nil {
		return dependencies
	}

	for _, f := range structType.Fields.List {
		dependency := model.Dependency{
			Name:      f.Names[0].String(),
			Type:      field.GetType(f),
			Interface: interfaces.Get(f, file),
		}

		dependencies = append(dependencies, dependency)
	}

	return dependencies
}
