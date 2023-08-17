package loop

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

type FunctionCall struct{}

func NewFunctionCall() *FunctionCall {
	return &FunctionCall{}
}

func (s *FunctionCall) Extend(function model.Function) func(loop *test.Loop) {
	return func(loop *test.Loop) {
		loop.FunctionCall = test.Statement{
			Lhs: s.getOutput(function),
			Rhs: fmt.Sprintf("instance.%s(%s)", function.Name, s.getInput(function)),
		}
	}
}

func (s *FunctionCall) getInput(function model.Function) string {
	arguments := make([]string, 0)

	for _, argument := range function.InputArguments {
		if argument.Type == model.ArgumentContext {
			arguments = append(arguments, "context.Background()")
			continue
		}

		arguments = append(arguments, fmt.Sprintf("tc.%s", pointer.Val(argument.Name)))
	}

	return strings.Join(arguments, ", ")
}

func (s *FunctionCall) getOutput(function model.Function) string {
	arguments := make([]string, 0)

	var argumentsAmount int

	for _, argument := range function.OutputArguments {
		if argument.Type == model.ArgumentError {
			arguments = append(arguments, "err")
			continue
		}

		if argumentsAmount == 0 {
			arguments = append(arguments, "out")
		} else {
			arguments = append(arguments, fmt.Sprintf("out%d", argumentsAmount))
		}

		argumentsAmount++
	}

	return strings.Join(arguments, ", ")
}
