package presenter

import "github.com/LaHainee/go_test_template_gen/internal/model"

type Presenter struct {
	structure structurePresenter
	loop      loopPresenter
}

func NewPresenter(sp structurePresenter, lp loopPresenter) *Presenter {
	return &Presenter{
		structure: sp,
		loop:      lp,
	}
}

func (p *Presenter) PresentLoop(function model.Function) string {
	return p.loop.Present(function)
}

func (p *Presenter) PresentStructure(function model.Function) string {
	return p.structure.Present(function)
}

func (p *Presenter) PresentImports(_ model.Function) []string {
	return []string{
		"\"context\"",
		"\"testing\"",
		"\"github.com/golang/mock/gomock\"",
		"\"github.com/stretchr/testify/assert\"",
	}
}
