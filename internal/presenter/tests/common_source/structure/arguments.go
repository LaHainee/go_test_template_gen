package structure

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

type Arguments struct{}

func NewArguments() *Arguments {
	return &Arguments{}
}

func (s *Arguments) Extend(function model.Function) func(structure *test.Structure) {
	if len(function.InputArguments) == 0 {
		return func(structure *test.Structure) {
			structure.FunctionArguments = []test.Statement{}
		}
	}

	declarations := make([]test.Statement, 0)
	declarations = append(declarations, s.getInputArgumentsDeclarations(function)...)
	declarations = append(declarations, s.getNonInterfaceDependenciesDeclarations(function)...)

	return func(structure *test.Structure) {
		structure.FunctionArguments = declarations
	}
}

func (s *Arguments) getNonInterfaceDependenciesDeclarations(function model.Function) []test.Statement {
	if function.Receiver == nil {
		return []test.Statement{}
	}

	statements := make([]test.Statement, 0, len(function.Receiver.Dependencies))

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface != nil {
			continue
		}

		statements = append(statements, test.Statement{
			Lhs: dependency.Name,
			Rhs: dependency.Type,
		})
	}

	return statements
}

func (s *Arguments) getInputArgumentsDeclarations(function model.Function) []test.Statement {
	statements := make([]test.Statement, 0, len(function.InputArguments))

	for _, argument := range function.InputArguments {
		if argument.Name == nil {
			continue
		}

		if argument.Type == model.ArgumentContext {
			continue
		}

		statements = append(statements, test.Statement{
			Lhs: pointer.Val(argument.Name),
			Rhs: argument.Type,
		})
	}

	return statements
}
