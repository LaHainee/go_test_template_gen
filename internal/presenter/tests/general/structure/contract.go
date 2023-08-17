package structure

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
	presenter "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type src interface {
	Extend(function model.Function) func(structure *presenter.Structure)
}
