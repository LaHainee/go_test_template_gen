package loop

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type InstanceCreate struct{}

func NewInstanceCreate() *InstanceCreate {
	return &InstanceCreate{}
}

func (s *InstanceCreate) Extend(function model.Function) func(loop *test.Loop) {
	if function.Receiver == nil {
		return func(loop *test.Loop) {
			loop.InstanceCreate = nil
		}
	}

	arguments := make([]string, 0)

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil {
			arguments = append(arguments, fmt.Sprintf("tc.%s", dependency.Name))
			continue
		}

		if dependency.IsLogger() {
			arguments = append(arguments, "log")
			continue
		}

		arguments = append(arguments, getMockName(dependency))
	}

	return func(loop *test.Loop) {
		loop.InstanceCreate = &test.Statement{
			Lhs: "instance",
			Rhs: fmt.Sprintf("New%s(%s)", function.Receiver.Name, strings.Join(arguments, ", ")),
		}
	}
}
