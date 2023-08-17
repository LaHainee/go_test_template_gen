package loop

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	presenter "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type ExpectationCall struct{}

func NewExpectationCall() *ExpectationCall {
	return &ExpectationCall{}
}

func (s *ExpectationCall) Extend(function model.Function) func(loop *presenter.Loop) {
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

	return func(loop *presenter.Loop) {
		loop.ExpectationCall = fmt.Sprintf("tc.expectations(t, %s)", strings.Join(arguments, ", "))
	}
}
