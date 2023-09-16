package loop

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
	presenter "github.com/LaHainee/go_test_template_gen/internal/presenter/tests/model"
)

type Presenter struct {
	sources []src
}

func NewPresenter(sources ...src) *Presenter {
	return &Presenter{
		sources: sources,
	}
}

func (s *Presenter) Present(function model.Function) string {
	var loop presenter.Loop

	for _, source := range s.sources {
		apply := source.Extend(function)

		apply(&loop)
	}

	return loop.Present()
}
