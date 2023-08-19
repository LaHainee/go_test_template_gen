package functions

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type SourceFunction interface {
	Extend(funcDecl *ast.FuncDecl, astFile *ast.File, file *model.File, function *model.Function) error
	SetNext(next SourceFunction)
}
