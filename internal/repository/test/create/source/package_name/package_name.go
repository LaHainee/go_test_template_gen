package package_name

import (
	"fmt"

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
	isNewFile := len(rows) == 0 || rows[0] == ""

	if !isNewFile {
		if s.next == nil {
			return rows, nil
		}
		return s.next.Add(rows, testFile)
	}

	rows = make([]string, 0)
	rows = append(rows, fmt.Sprintf("package %s", testFile.Package.Name))
	rows = append(rows, "")

	if s.next == nil {
		return rows, nil
	}
	return s.next.Add(rows, testFile)
}
