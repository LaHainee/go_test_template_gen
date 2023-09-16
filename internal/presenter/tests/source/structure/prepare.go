package structure

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type Prepare struct{}

func NewPrepare() *Prepare {
	return &Prepare{}
}

func (s *Prepare) Extend(function model.Function) func(structure *test.Structure) {
	mocks := make([]string, 0)

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil || dependency.IsGrace() || dependency.IsMetric() || dependency.IsLogger() {
			continue
		}

		mocks = append(mocks, fmt.Sprintf("%s *Mock%s", dependency.Name, dependency.Type))
	}

	if len(mocks) == 0 {
		return func(structure *test.Structure) {
			structure.Prepare = nil
		}
	}

	return func(structure *test.Structure) {
		structure.Prepare = &test.Statement{
			Lhs: "prepare",
			Rhs: fmt.Sprintf("func(%s)", strings.Join(mocks, ", ")),
		}
	}
}
