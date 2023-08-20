package ast

import (
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Facade struct {
	source Source
}

func NewFacade(source Source) *Facade {
	return &Facade{
		source: source,
	}
}

func (f *Facade) Parse(filePath model.FilePath) (model.File, error) {
	astFile, err := getAstFile(filePath)
	if err != nil {
		return model.File{}, err
	}

	var file model.File

	err = f.source.Extend(filePath, astFile, &file)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func getAstFile(filePath model.FilePath) (*ast.File, error) {
	file, err := parser.ParseFile(token.NewFileSet(), filePath.String(), nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	err, ok := err.(scanner.ErrorList)
	if ok {
		return nil, err
	}

	err = expandAstScopeToDirectory(filePath, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
expandAstScopeToDirectory - расширить *ast.File.Scope до директории в котрой находится файл

Исходный скоуп покрывает только рассматриваемый файл. Для того, чтобы работать с объявлениями
в соседних файлах необходимо расширить скоуп до всей директории, т.е. покрыть все файлы в директории
*/
func expandAstScopeToDirectory(filePath model.FilePath, source *ast.File) error {
	filePaths, err := filePath.DirectoryFilePaths()
	if err != nil {
		return err
	}

	scope := ast.NewScope(nil)

	for _, filePath = range filePaths {
		if !filePath.IsGolangSource() {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), string(filePath), nil, parser.AllErrors)
		if err != nil {
			return err
		}
		err, ok := err.(scanner.ErrorList)
		if ok {
			return err
		}

		if file.Scope == nil {
			continue
		}

		for _, object := range file.Scope.Objects {
			scope.Insert(object)
		}
	}

	source.Scope = scope

	return nil
}
