package functions

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	facade "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast"
)

type Source struct {
	source SourceFunction
	next   facade.Source
}

func NewSource(source SourceFunction) *Source {
	return &Source{
		source: source,
	}
}

func (s *Source) SetNext(next facade.Source) {
	s.next = next
}

func (s *Source) Extend(filePath model.FilePath, astFile *ast.File, file *model.File) error {
	functions, err := s.getFunctions(astFile, file)
	if err != nil {
		return err
	}

	file.Functions = functions

	if s.next != nil {
		return s.next.Extend(filePath, astFile, file)
	}
	return nil
}

func (s *Source) getFunctions(astFile *ast.File, file *model.File) ([]model.Function, error) {
	functions := make([]model.Function, 0)

	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		function := model.Function{
			Name: funcDecl.Name.Name,
		}

		err := s.source.Extend(funcDecl, astFile, file, &function)
		if err != nil {
			return nil, err
		}

		functions = append(functions, function)
	}

	return functions, nil
}
