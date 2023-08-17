package test

import (
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/test/create"
)

type Source struct {
	next create.Source
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) SetNext(next create.Source) {
	s.next = next
}

func (s *Source) Add(rows []string, testFile model.TestFile) ([]string, error) {
	for _, test := range testFile.Tests {
		rows = s.add(rows, test)
	}

	if s.next == nil {
		return rows, nil
	}
	return s.next.Add(rows, testFile)
}

func (s *Source) add(rows []string, test model.Test) []string {
	testRows := strings.Split(test.Source, "\n")

	return append(rows, testRows...)
}
