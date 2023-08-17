package create

import "github.com/LaHainee/go_test_template_gen/internal/model"

type Source interface {
	Add(rows []string, testFile model.TestFile) ([]string, error)
	SetNext(source Source)
}
