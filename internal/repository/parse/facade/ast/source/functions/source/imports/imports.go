package imports

import (
	"errors"
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/functions"
)

type Source struct {
	next functions.SourceFunction
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) SetNext(next functions.SourceFunction) {
	s.next = next
}

func (s *Source) Extend(funcDecl *ast.FuncDecl, astFile *ast.File, file *model.File, function *model.Function) error {
	imports := make([]string, 0)

	imports = append(imports, s.importsForFunctionReceiverFields(file, function)...)
	imports = append(imports, s.importsForFunctionArgs(file, function)...)

	function.NeededImports = model.NewImports().Append(imports...)

	if s.next != nil {
		return s.next.Extend(funcDecl, astFile, file, function)
	}
	return nil
}

func (s *Source) importsForFunctionReceiverFields(file *model.File, function *model.Function) []string {
	imports := make([]string, 0)

	if function.Receiver == nil {
		return imports
	}

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Package == nil {
			continue
		}

		importPath, err := file.Imports.Search(*dependency.Package)
		if errors.Is(err, model.ErrNotFound) {
			continue
		}

		imports = append(imports, importPath)
	}

	return imports
}

func (s *Source) importsForFunctionArgs(file *model.File, function *model.Function) []string {
	imports := make([]string, 0)

	// Для входных аргументов
	for _, argument := range function.InputArguments {
		if argument.Package == nil {
			continue
		}

		importPath, err := file.Imports.Search(*argument.Package)
		if errors.Is(err, model.ErrNotFound) {
			continue
		}

		imports = append(imports, importPath)
	}

	// Аналогично для выходных
	for _, argument := range function.OutputArguments {
		if argument.Package == nil {
			continue
		}

		importPath, err := file.Imports.Search(*argument.Package)
		if errors.Is(err, model.ErrNotFound) {
			continue
		}

		imports = append(imports, importPath)
	}

	return imports
}
