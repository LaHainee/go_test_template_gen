package imports

import "github.com/LaHainee/go_test_template_gen/internal/model"

type WithMock struct{}

func NewWithMock() *WithMock {
	return &WithMock{}
}

func (i *WithMock) Get(function model.Function) []string {
	importsList := append(defaultImportsList, "\"go.uber.org/mock/gomock\"")

	if function.Receiver == nil {
		return importsList
	}

	for _, dep := range function.Receiver.Dependencies {
		if dep.IsLogger() {
			importsList = append(importsList, "\"go.avito.ru/gl/logger/v3\"")
		}
	}

	return importsList
}
