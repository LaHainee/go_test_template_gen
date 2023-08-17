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
	files := make([]model.File, 0)

	sourceFile, err := s.file.Parse(path)
	if err != nil {
		return nil, err
	}

	files = append(files, sourceFile)

	testFile, err := s.file.Parse(path.ToTest())
	if err == nil {
		files = append(files, testFile)
	}

	return files, nil
}
