package by_filepath

import "github.com/LaHainee/go_test_template_gen/internal/model"

type fileParser interface {
	Parse(filepath model.FilePath) (model.File, error)
}
