package imports

import "github.com/LaHainee/go_test_template_gen/internal/model"

type WithOutMock struct{}

func NewWithOutMock() *WithOutMock {
	return &WithOutMock{}
}

func (i *WithOutMock) Get(_ model.Function) []string {
	return defaultImportsList
}
