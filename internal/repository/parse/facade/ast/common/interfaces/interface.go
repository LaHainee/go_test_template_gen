package interfaces

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/arguments"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/ast/types"
)

func Get(field *ast.Field, file *ast.File) *model.Interface {
	interfaceType := types.GetInterface(field.Type, file)
	if interfaceType == nil {
		return nil
	}

	if interfaceType.Methods == nil {
		return &model.Interface{}
	}

	functions := make([]model.Function, 0)

	for _, method := range interfaceType.Methods.List {
		funcType, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		function := model.Function{
			Name:            method.Names[0].String(),
			InputArguments:  arguments.Get(funcType.Params),
			OutputArguments: arguments.Get(funcType.Results),
		}

		functions = append(functions, function)
	}

	return &model.Interface{
		Functions: functions,
	}
}
