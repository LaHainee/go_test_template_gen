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

	// Если среди входных аргументов есть указатели, то их нужно вынести в блок expectations
	for _, argument := range function.InputArguments {
		// Пропускаем аргументы без указателя
		if !argument.IsPointer() {
			continue
		}

		// Пропускаем аргументы, которые не содержат в своем названии out
		if !argument.NameContains("out") {
			continue
		}

		if argumentsAmount == 0 {
			arguments = append(arguments, fmt.Sprintf("got %s", argument.Dereference()))
		} else {
			arguments = append(arguments, fmt.Sprintf("got%d %s", argumentsAmount, argument.Dereference()))
		}

		argumentsAmount++
	}

	for _, argument := range function.OutputArguments {
		if argument.Is(model.ArgumentError) {
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
