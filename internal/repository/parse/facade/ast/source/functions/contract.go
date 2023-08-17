package functions

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type src interface {
	Extend(funcDecl *ast.FuncDecl, file *ast.File) (func(function *model.Function), error)
}
