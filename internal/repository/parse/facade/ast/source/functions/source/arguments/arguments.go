package arguments

import (
	"errors"
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/arguments"
)

type Source struct{}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Extend(funcDecl *ast.FuncDecl, _ *ast.File) (func(function *model.Function), error) {
	if funcDecl.Type == nil {
		return nil, errors.New("corrupted function type")
	}

	return func(function *model.Function) {
		function.InputArguments = arguments.Get(funcDecl.Type.Params)
		function.OutputArguments = arguments.Get(funcDecl.Type.Results)
	}, nil
}
