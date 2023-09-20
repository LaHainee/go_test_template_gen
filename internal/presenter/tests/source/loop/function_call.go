package loop

import (
	"fmt"
	domain "github.com/LaHainee/go_test_template_gen/internal/domain/function"
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
			Rhs: s.getRhs(function),
		}
	}
}

func (s *FunctionCall) getRhs(function model.Function) string {
	functionType := domain.GetFunctionType(function)

	functionInput := s.getInput(function)

	if functionType == domain.TypeFunction {
		return fmt.Sprintf("%s(%s)", function.Name, functionInput)
	}
	return fmt.Sprintf("instance.%s(%s)", function.Name, functionInput)
}

func (s *FunctionCall) getInput(function model.Function) string {
	arguments := make([]string, 0)

	for _, argument := range function.InputArguments {
		if argument.Is(model.ArgumentContext) {
			arguments = append(arguments, "context.Background()")
			continue
		}

		template := "tc.%s"
		if argument.IsPointer() {
			template = "&tc.%s"
		}

		arguments = append(arguments, fmt.Sprintf(template, pointer.Val(argument.Name)))
	}

	return strings.Join(arguments, ", ")
}

func (s *FunctionCall) getOutput(function model.Function) string {
	arguments := make([]string, 0)

	var argumentsAmount int

	for _, argument := range function.OutputArguments {
		if argument.Is(model.ArgumentError) {
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
