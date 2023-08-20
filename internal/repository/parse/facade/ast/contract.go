package ast

import (
	"go/ast"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Source interface {
	Extend(filePath model.FilePath, astFile *ast.File, file *model.File) error
	SetNext(src Source)
}
