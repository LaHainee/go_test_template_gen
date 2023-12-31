package arguments

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/field"
	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

func Get(fieldList *ast.FieldList) []model.Argument {
	arguments := make([]model.Argument, 0)

	if fieldList == nil {
		return arguments
	}

	for _, f := range fieldList.List {
		fieldNames := field.GetNames(f)
		fieldType := field.GetType(f)
		fieldPackages := field.GetPackages(f, []string{})

		// Аргументы без имен
		if len(fieldNames) == 0 {
			arguments = append(arguments, model.Argument{
				Type:     fieldType,
				Packages: fieldPackages,
			})

			continue
		}

		for _, name := range fieldNames {
			arguments = append(arguments, model.Argument{
				Name:     pointer.To(name),
				Type:     fieldType,
				Packages: fieldPackages,
			})
		}

	}

	return arguments
}
