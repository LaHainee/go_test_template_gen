package instance_create

import (
	"fmt"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/source/loop/instance_create/arguments"
)

type InstanceCreate struct{}

func NewInstanceCreate() *InstanceCreate {
	return &InstanceCreate{}
}

func (s *InstanceCreate) Extend(function model.Function) func(loop *test.Loop) {
	if function.Receiver == nil {
		return func(loop *test.Loop) { loop.InstanceCreate = nil }
	}

	argsGetter := createArgumentsGetter(function)

	return func(loop *test.Loop) {
		loop.InstanceCreate = &test.Statement{
			Lhs: "instance",
			Rhs: fmt.Sprintf("%s(%s)", s.getConstructorName(function), argsGetter.Get(function)),
		}
	}
}

func (s *InstanceCreate) getConstructorName(function model.Function) string {
	if function.Receiver.Constructor != nil {
		return function.Receiver.Constructor.Name
	}

	return fmt.Sprintf("New%s", function.Receiver.Name)
}

type argumentsGetter interface {
	Get(function model.Function) string
}

func createArgumentsGetter(function model.Function) argumentsGetter {
	if function.Receiver.Constructor != nil &&
		function.Receiver.Constructor.Return != nil &&
		function.Receiver.Constructor.Return.Structure != nil {
		return arguments.NewConstructor()
	}

	// Деградация
	return arguments.NewStructure()
}
