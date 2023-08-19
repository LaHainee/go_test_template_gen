package imports

import (
	"go/ast"
	"go/token"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	facade "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast"
)

type Source struct {
	next facade.Source
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) SetNext(next facade.Source) {
	s.next = next
}

func (s *Source) Extend(filePath model.FilePath, astFile *ast.File, file *model.File) error {
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.IMPORT {
			continue
		}

		file.Imports = model.NewImports().Append(getImports(genDecl)...)
		break
	}

	if s.next != nil {
		return s.next.Extend(filePath, astFile, file)
	}
	return nil
}

func getImports(genDecl *ast.GenDecl) []string {
	imports := make([]string, 0)

	for _, spec := range genDecl.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		imports = append(imports, getImport(importSpec))
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
