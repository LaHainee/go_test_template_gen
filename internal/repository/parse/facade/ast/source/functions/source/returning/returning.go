package returning

import (
	"errors"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/functions"
	"go/ast"
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
	function.Return = s.getReturn(funcDecl)

	if s.next != nil {
		return s.next.Extend(funcDecl, astFile, file, function)
	}
	return nil
}

func (s *Source) getReturn(funcDecl *ast.FuncDecl) *model.Return {
	if funcDecl.Body == nil {
		return nil
	}

	// У функции может быть несколько return statements, считаем, что корректный – самый последний
	returnStatement, err := s.getLastReturnStatement(funcDecl.Body.List)
	if errors.Is(err, model.ErrNotFound) {
		return nil
	}

	return &model.Return{
		Structure: s.getReturnStructure(returnStatement),
	}
}

func (s *Source) getLastReturnStatement(statements []ast.Stmt) (*ast.ReturnStmt, error) {
	for i := len(statements) - 1; i >= 0; i-- {
		returnStatement, ok := statements[i].(*ast.ReturnStmt)
		if !ok {
			continue
		}

		return returnStatement, nil
	}

	return nil, model.ErrNotFound
}

func (s *Source) getReturnStructure(returnStatement *ast.ReturnStmt) *model.ReturnStructure {
	if len(returnStatement.Results) == 0 {
		return nil
	}

	// Среди аргументов, котоыре возвращает функция ищем *ast.UnaryExpr – это структура
	var unaryExpr *ast.UnaryExpr
	for _, res := range returnStatement.Results {
		converted, ok := res.(*ast.UnaryExpr)
		if !ok {
			continue
		}

		unaryExpr = converted
		break
	}

	if unaryExpr == nil {
		return nil
	}

	compositeLit, ok := unaryExpr.X.(*ast.CompositeLit)
	if !ok {
		return nil
	}

	if len(compositeLit.Elts) == 0 {
		return nil
	}

	argumentsBindings := make(map[string]string)

	for _, elt := range compositeLit.Elts {
		keyValueExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		key, ok := keyValueExpr.Key.(*ast.Ident)
		if !ok {
			continue
		}

		value, ok := keyValueExpr.Value.(*ast.Ident)
		if !ok {
			continue
		}

		argumentsBindings[value.String()] = key.String()
	}

	return &model.ReturnStructure{
		ArgumentBindings: argumentsBindings,
	}
}
