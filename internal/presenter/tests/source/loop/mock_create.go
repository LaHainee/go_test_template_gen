package loop

import (
	"fmt"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	test "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type MockCreate struct{}

func NewMockCreate() *MockCreate {
	return &MockCreate{}
}

func (s *MockCreate) Extend(function model.Function) func(loop *test.Loop) {
	statements := make([]test.Statement, 0)
	statements = append(statements, test.Statement{
		Lhs: "ctrl",
		Rhs: "gomock.NewController(t)",
	})

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil {
			continue
		}

		if dependency.IsLogger() {
			statements = append(statements, s.getLoggerDeclaration())
			continue
		}

		statements = append(statements, test.Statement{
			Lhs: getMockName(dependency),
			Rhs: fmt.Sprintf("NewMock%s(ctrl)", dependency.Type),
		})
	}

	return func(loop *test.Loop) {
		loop.MockCreate = statements
	}
}

func (s *MockCreate) getLoggerDeclaration() test.Statement {
	return test.Statement{
		Lhs: "log, _",
		Rhs: "logger.New(logger.WithEnabled(false))",
	}
}
