package loop

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	presenter "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type MockPrepare struct{}

func NewMockPrepare() *MockPrepare {
	return &MockPrepare{}
}

func (s *MockPrepare) Extend(function model.Function) func(loop *presenter.Loop) {
	rows := s.getPrepareFunctionCall(function)
	rows = append(rows, s.getMocksWithExpect(function)...)

	return func(loop *presenter.Loop) {
		loop.MockPrepare = rows
	}
}

func (s *MockPrepare) getPrepareFunctionCall(function model.Function) []string {
	mocks := make([]string, 0)
	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil || dependency.IsMetric() || dependency.IsGrace() || dependency.IsLogger() {
			continue
		}

		mocks = append(mocks, getMockName(dependency))
	}

	rows := make([]string, 0)
	rows = append(rows, "if tc.prepare != nil {")
	rows = append(rows, fmt.Sprintf("\ttc.prepare(%s)", strings.Join(mocks, ", ")))
	rows = append(rows, "}")

	return rows
}

func (s *MockPrepare) getMocksWithExpect(function model.Function) []string {
	rows := make([]string, 0)

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil {
			continue
		}

		if !dependency.IsGrace() && !dependency.IsMetric() {
			continue
		}

		for _, f := range dependency.Interface.Functions {
			rows = append(rows, s.getMockWithExpect(dependency, f))
		}
	}

	return rows
}

func (s *MockPrepare) getMockWithExpect(dependency model.Dependency, function model.Function) string {
	gomockAnyCalls := make([]string, 0)
	for i := 0; i < len(function.InputArguments); i++ {
		gomockAnyCalls = append(gomockAnyCalls, "gomock.Any()")
	}

	return fmt.Sprintf(
		"%s.EXPECT().%s(%s).AnyTimes()",
		getMockName(dependency),
		function.Name,
		strings.Join(gomockAnyCalls, ", "),
	)
}
