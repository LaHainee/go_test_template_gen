package presenter

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
	commonSourceLoop "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/common_source/loop"
	commonSourceStructure "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/common_source/structure"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/general"
	generalLoop "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/general/loop"
	generalStructure "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/general/structure"
)

type Factory struct {
	generalPresenter TestPresenter
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
	if function.HasInterfaceDependencies() {
		return f.getGeneralPresenter(), nil
	}

	return nil, model.ErrUnsupported
}

func (f *Factory) getGeneralPresenter() TestPresenter {
	if f.generalPresenter == nil {
		f.mustInitGeneralPresenter()
	}

	return f.generalPresenter
}

func (f *Factory) mustInitGeneralPresenter() {
	presenter := general.NewPresenter(
		generalStructure.NewPresenter(
			commonSourceStructure.NewName(),
			commonSourceStructure.NewArguments(),
			commonSourceStructure.NewPrepare(),
			commonSourceStructure.NewExpectation(),
		),
		generalLoop.NewPresenter(
			commonSourceLoop.NewMockCreate(),
			commonSourceLoop.NewMockPrepare(),
			commonSourceLoop.NewInstanceCreate(),
			commonSourceLoop.NewFunctionCall(),
			commonSourceLoop.NewExpectationCall(),
		),
	)

	f.generalPresenter = presenter
}
