package imports

import (
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/test/create"
)

type Source struct {
	filesystem filesystem
	next       create.Source
}

func NewSource(fs filesystem) *Source {
	return &Source{
		filesystem: fs,
	}
}

func (s *Source) SetNext(next create.Source) {
	s.next = next
}

func (s *Source) Add(rows []string, testFile model.TestFile) ([]string, error) {
	var err error

	rows, err = s.add(rows, testFile)
	if err != nil {
		return nil, err
	}

	if s.next == nil {
		return rows, nil
	}
	return s.next.Add(rows, testFile)
}

func (s *Source) add(rows []string, testFile model.TestFile) ([]string, error) {
	for num, row := range rows {
		if strings.Contains(row, "import") {
			return s.update(rows, num, testFile)
		}
	}

	return s.insert(rows, testFile)
}
