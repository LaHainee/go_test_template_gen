package by_dirpath

import "github.com/LaHainee/go_test_template_gen/internal/model"

type Source struct {
	file fileParser
}

func NewSource(fp fileParser) *Source {
	return &Source{
		file: fp,
	}
}

func (s *Source) Get(path model.FilePath) ([]model.File, error) {
	return s.file.ParseDirectory(path)
}
