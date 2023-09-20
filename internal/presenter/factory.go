package presenter

import (
	domain "github.com/LaHainee/go_test_template_gen/internal/domain/function"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/presenter"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/presenter/loop"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/presenter/structure"
	sourceLoop "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/source/loop"
	sourceStructure "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/source/structure"
)

type Factory struct {
	presenterMethodWithMocks    TestPresenter
	presenterMethodWithoutMocks TestPresenter
	presenterFunction           TestPresenter
}

func NewFactory() *Factory {
	return &Factory{}
}

type TestPresenter interface {
	PresentLoop(function model.Function) string
	PresentStructure(function model.Function) string
	PresentImports(function model.Function) []string
}

func (f *Factory) CreateTestPresenter(function model.Function) (TestPresenter, error) {
	functionType := domain.GetFunctionType(function)

	switch functionType {
	case domain.TypeMethodWithInterfaceDeps:
		return f.getPresenterMethodWithMocks(), nil
	case domain.TypeMethodWithoutInterfaceDeps:
		return f.getPresenterMethodWithoutMocks(), nil
	case domain.TypeFunction:
		return f.getPresenterFunction(), nil
	}

	return nil, model.ErrUnsupported
}

func (f *Factory) getPresenterMethodWithoutMocks() TestPresenter {
	if f.presenterMethodWithoutMocks == nil {
		f.mustInitPresenterMethodWithoutMocks()
	}

	return f.presenterMethodWithoutMocks
}

func (f *Factory) getPresenterMethodWithMocks() TestPresenter {
	if f.presenterMethodWithMocks == nil {
		f.mustInitPresenterMethodWithMocks()
	}

	return f.presenterMethodWithMocks
}

func (f *Factory) getPresenterFunction() TestPresenter {
	if f.presenterFunction == nil {
		f.mustInitPresenterFunction()
	}

	return f.presenterFunction
}

func (f *Factory) mustInitPresenterFunction() {
	f.presenterFunction = presenter.NewPresenter(
		structure.NewPresenter(
			sourceStructure.NewName(),
			sourceStructure.NewArguments(),
			sourceStructure.NewExpectation(),
		),
		loop.NewPresenter(
			sourceLoop.NewFunctionCall(),
			sourceLoop.NewExpectationCall(),
		),
	)
}

func (f *Factory) mustInitPresenterMethodWithoutMocks() {
	f.presenterMethodWithoutMocks = presenter.NewPresenter(
		structure.NewPresenter(
			sourceStructure.NewName(),
			sourceStructure.NewArguments(),
			sourceStructure.NewPrepare(),
			sourceStructure.NewExpectation(),
		),
		loop.NewPresenter(
			sourceLoop.NewInstanceCreate(),
			sourceLoop.NewFunctionCall(),
			sourceLoop.NewExpectationCall(),
		),
	)
}

func (f *Factory) mustInitPresenterMethodWithMocks() {
	f.presenterMethodWithMocks = presenter.NewPresenter(
		structure.NewPresenter(
			sourceStructure.NewName(),
			sourceStructure.NewArguments(),
			sourceStructure.NewPrepare(),
			sourceStructure.NewExpectation(),
		),
		loop.NewPresenter(
			sourceLoop.NewMockCreate(),
			sourceLoop.NewMockPrepare(),
			sourceLoop.NewInstanceCreate(),
			sourceLoop.NewFunctionCall(),
			sourceLoop.NewExpectationCall(),
		),
	)
}
