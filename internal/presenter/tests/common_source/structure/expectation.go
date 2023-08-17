package structure

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type Expectation struct{}

func NewExpectation() *Expectation {
	return &Expectation{}
}

func (s *Expectation) Extend(function model.Function) func(structure *test.Structure) {
	arguments := make([]string, 0)
	arguments = append(arguments, "t assert.TestingT")

	var argumentsAmount int

	for _, argument := range function.OutputArguments {
		if argument.Type == model.ArgumentError {
			arguments = append(arguments, "err error")
			continue
		}

		if argumentsAmount == 0 {
			arguments = append(arguments, fmt.Sprintf("got %s", argument.Type))
		} else {
			arguments = append(arguments, fmt.Sprintf("got%d %s", argumentsAmount, argument.Type))
		}

		argumentsAmount++
	}

	return func(structure *test.Structure) {
		structure.Expectation = test.Statement{
			Lhs: "expectations",
			Rhs: fmt.Sprintf("func(%s)", strings.Join(arguments, ", ")),
		}
	}
}
