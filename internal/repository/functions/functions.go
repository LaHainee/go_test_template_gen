package functions

import (
	"errors"
	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/fs"
	"strings"
	"syscall"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUncovered(functions []model.Function, testPath model.FilePath) ([]model.Function, error) {
	file, err := parser.ParseFile(token.NewFileSet(), testPath.String(), nil, parser.AllErrors)
	if err != nil {
		var (
			pathErr *fs.PathError
			errNo   syscall.Errno
		)

		if errors.As(err, &pathErr) && errors.As(pathErr, &errNo) {
			return nil, model.ErrNotFound
		}

		return nil, err
	}
	err, ok := err.(scanner.ErrorList)
	if ok {
		return nil, err
	}

	uncovered := make([]model.Function, 0, len(functions))

	for _, function := range functions {
		if isCalled(function, file) {
			continue
		}

		uncovered = append(uncovered, function)
	}

	return uncovered, nil
}

func isCalled(function model.Function, file *ast.File) bool {
	for _, decl := range file.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			// Найденна функция не является тестом, пропускаем
			if !strings.Contains(t.Name.String(), "Test") {
				continue
			}

			var foundFunction bool    // найден ли вызов функции
			var foundConstructor bool // найден ли вызов конструктора

			var constructorName *string
			if function.Receiver != nil && function.Receiver.Constructor != nil {
				constructorName = pointer.To(function.Receiver.Constructor.Name)
			}

			if constructorName == nil {
				foundConstructor = true
			}

			ast.Inspect(t, func(node ast.Node) bool {
				if foundConstructor && foundFunction {
					return false
				}

				// TODO: подумать над улучшение качества поиска
				switch t := node.(type) {
				case *ast.CallExpr:
					switch tt := t.Fun.(type) {
					case *ast.Ident:
						if constructorName != nil && tt.Name == *constructorName {
							foundConstructor = true
						}

						if tt.Name == function.Name {
							foundFunction = true
						}
					case *ast.SelectorExpr:
						if tt.Sel == nil {
							return false
						}

						if constructorName != nil && tt.Sel.Name == *constructorName {
							foundConstructor = true
						}

						if tt.Sel.Name == function.Name {
							foundFunction = true
						}
					}
				}

				return true
			})

			if foundFunction && foundConstructor {
				return true
			}
		}
	}

	return false
}
