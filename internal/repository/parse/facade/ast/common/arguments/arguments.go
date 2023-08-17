package arguments

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/field"
)

func Get(fieldList *ast.FieldList) []model.Argument {
	arguments := make([]model.Argument, 0)

	if fieldList == nil {
		return arguments
	}

	for _, f := range fieldList.List {
		arguments = append(arguments, model.Argument{
			Name: field.GetName(f),
			Type: field.GetType(f),
		})
	}

	return arguments
}
