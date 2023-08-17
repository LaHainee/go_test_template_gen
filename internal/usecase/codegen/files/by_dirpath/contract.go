package by_dirpath

import "github.com/LaHainee/go_test_template_gen/internal/model"

type fileParser interface {
	ParseDirectory(directoryPath model.FilePath) ([]model.File, error)
}
