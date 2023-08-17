package functions

import (
	"errors"
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Source struct {
	sources []src
}

func NewSource(sources ...src) *Source {
	return &Source{
		sources: sources,
	}
}

func (s *Source) Extend(_ model.FilePath, file *ast.File) (func(file *model.File), error) {
	functions := make([]model.Function, 0)

	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		function := model.Function{
			Name: funcDecl.Name.Name,
		}

		for _, source := range s.sources {
			apply, err := source.Extend(funcDecl, file)
			if err != nil {
				return nil, err
			}

			apply(&function)
		}

		functions = append(functions, function)
	}

	s.setConstructorsName(functions)

	return func(file *model.File) {
		file.Functions = functions
	}, nil
}

func (s *Source) setConstructorsName(functions []model.Function) {
	for i, function := range functions {
		if function.Receiver == nil {
			continue
		}

		constructor, err := model.Functions(functions).LookupByOutputArgument(function.Receiver.Name)
		if errors.Is(err, model.ErrNotFound) {
			continue
		}

		functions[i].Receiver.ConstructorName = &constructor.Name
	}
}
