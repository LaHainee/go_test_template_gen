package receiver

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/structure"
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
	function.Receiver = s.getReceiver(funcDecl, astFile)

	if s.next != nil {
		return s.next.Extend(funcDecl, astFile, file, function)
	}
	return nil
}

func (s *Source) getReceiver(funcDecl *ast.FuncDecl, astFile *ast.File) *model.Structure {
	if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
		return nil
	}

	receiver := funcDecl.Recv.List[0]

	return structure.Get(receiver, astFile)
}
