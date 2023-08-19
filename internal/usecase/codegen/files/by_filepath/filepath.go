package by_filepath

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Source struct {
	file fileParser
}

func NewSource(fp fileParser) *Source {
	return &Source{
		file: fp,
	}
}

func (s *Source) Get(path model.FilePath) ([]model.File, error) {
	// Приходится парсить целиком директорию, поскольку объявления могут лежать в разных
	files, err := s.file.ParseDirectory(model.FilePath(path.DirectoryPath()))
	if err != nil {
		return nil, err
	}

	needed := make([]model.File, 0)

	for _, file := range files {
		if file.Path == path {
			needed = append(needed, file)
		}

		if file.Path == path.ToTest() {
			needed = append(needed, file)
		}
	}

	return needed, nil
}
