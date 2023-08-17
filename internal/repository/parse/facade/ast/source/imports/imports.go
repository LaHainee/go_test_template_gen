package imports

import (
	"go/ast"
	"go/token"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Source struct{}

func NewSource() *Source {
	return &Source{}
}

//nolint:unparam
func (s *Source) Extend(_ model.FilePath, file *ast.File) (func(file *model.File), error) {
	var imports model.Imports

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.IMPORT {
			continue
		}

		imports = getImports(genDecl)
		break
	}

	return func(file *model.File) {
		file.Imports = imports
	}, nil
}

func getImports(genDecl *ast.GenDecl) model.Imports {
	imports := model.NewImports()

	for _, spec := range genDecl.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		imports = imports.Append(getImport(importSpec))
	}

	return imports
}

func getImport(importSpec *ast.ImportSpec) string {
	// У импорта отсутствует алиас
	if importSpec.Name == nil {
		return importSpec.Path.Value
	}

	return importSpec.Name.Name + " " + importSpec.Path.Value
}
