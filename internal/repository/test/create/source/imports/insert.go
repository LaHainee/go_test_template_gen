package imports

import "github.com/LaHainee/go_test_template_gen/internal/model"

func (s *Source) insert(rows []string, testFile model.TestFile) ([]string, error) {
	rows = append(rows, "import (")
	rows = append(rows, testFile.Imports.PresentReformatted()...)
	rows = append(rows, ")")
	rows = append(rows, "")

	return rows, nil
}
