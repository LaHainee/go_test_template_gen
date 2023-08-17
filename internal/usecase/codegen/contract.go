package codegen

import "github.com/LaHainee/go_test_template_gen/internal/model"

type filesGetter interface {
	Get(path string) ([]model.File, error)
}

type presenter interface {
	Present(files []model.File) []model.TestFile
}

type testRepository interface {
	Create(files []model.TestFile) error
}
