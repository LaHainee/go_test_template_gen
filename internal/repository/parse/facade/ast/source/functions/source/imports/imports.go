package imports

import (
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
		if len(dependency.Packages) == 0 {
			continue
		}

		imports = append(imports, getUsedImports(file.Imports, dependency.Packages)...)
	}

	return imports
}

func (s *Source) importsForFunctionArgs(file *model.File, function *model.Function) []string {
	imports := make([]string, 0)

	// Для входных аргументов
	for _, argument := range function.InputArguments {
		if len(argument.Packages) == 0 {
			continue
		}
		imports = append(imports, getUsedImports(file.Imports, argument.Packages)...)
	}

	// Аналогично для выходных
	for _, argument := range function.OutputArguments {
		if argument.Packages == nil {
			continue
		}
		imports = append(imports, getUsedImports(file.Imports, argument.Packages)...)
	}

	return imports
}

func getUsedImports(fileImports model.Imports, packages []string) []string {
	imports := make([]string, 0)

	for _, pkg := range packages {
		importPath, err := fileImports.Search(pkg)
		if err != nil {
			continue
		}

		imports = append(imports, importPath)
	}

	return imports
}
