package structure

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type Name struct{}

func NewName() *Name {
	return &Name{}
}

func (s *Name) Extend(_ model.Function) func(structure *test.Structure) {
	return func(structure *test.Structure) {
		structure.TestName = test.Statement{
			Lhs: "name",
			Rhs: "string",
		}
	}
}
