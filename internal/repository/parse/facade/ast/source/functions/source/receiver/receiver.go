package receiver

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/common/structure"
)

type Source struct{}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Extend(funcDecl *ast.FuncDecl, file *ast.File) (func(function *model.Function), error) {
	return func(function *model.Function) {
		function.Receiver = s.getReceiver(funcDecl, file)
	}, nil
}

func (s *Source) getReceiver(funcDecl *ast.FuncDecl, file *ast.File) *model.Structure {
	if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
		return nil
	}

	receiver := funcDecl.Recv.List[0]

	return structure.Get(receiver, file)
}
