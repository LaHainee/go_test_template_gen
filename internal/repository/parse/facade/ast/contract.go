package ast

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type src interface {
	Extend(filePath model.FilePath, file *ast.File) (func(file *model.File), error)
}
