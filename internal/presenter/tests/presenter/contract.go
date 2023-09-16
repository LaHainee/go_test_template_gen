package presenter

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type structurePresenter interface {
	Present(function model.Function) string
}

type loopPresenter interface {
	Present(function model.Function) string
}
