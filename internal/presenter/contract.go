package presenter

import "github.com/LaHainee/go_test_template_gen/internal/model"

type factory interface {
	CreateTestPresenter(function model.Function) (TestPresenter, error)
}
