package files

import "github.com/LaHainee/go_test_template_gen/internal/model"

type source interface {
	Get(path model.FilePath) ([]model.File, error)
}

type functionsRepository interface {
	GetUncovered(functions []model.Function, testPath model.FilePath) ([]model.Function, error)
}
