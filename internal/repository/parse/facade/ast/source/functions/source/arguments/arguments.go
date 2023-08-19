package arguments

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/arguments"
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
	if funcDecl.Type == nil {
		return nil
	}

	function.InputArguments = arguments.Get(funcDecl.Type.Params)
	function.OutputArguments = arguments.Get(funcDecl.Type.Results)

	if s.next != nil {
		return s.next.Extend(funcDecl, astFile, file, function)
	}
	return nil
}
