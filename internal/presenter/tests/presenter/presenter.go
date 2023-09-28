package presenter

import "github.com/LaHainee/go_test_template_gen/internal/model"

type Presenter struct {
	structure structurePresenter
	loop      loopPresenter
	imports   importsGetter
}

func NewPresenter(sp structurePresenter, lp loopPresenter, ig importsGetter) *Presenter {
	return &Presenter{
		structure: sp,
		loop:      lp,
		imports:   ig,
	}
}

func (p *Presenter) PresentLoop(function model.Function) string {
	return p.loop.Present(function)
}

func (p *Presenter) PresentStructure(function model.Function) string {
	return p.structure.Present(function)
}

func (p *Presenter) PresentImports(function model.Function) []string {
	return p.imports.Get(function)
}
