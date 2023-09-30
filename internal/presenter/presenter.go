package presenter

import (
	"errors"
	"fmt"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

const template = `func %s(t *testing.T) {
	t.Parallel()
	
	tests := []struct {
		%s
	}{
		%s
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			%s
		})
	}
}
`

var errEmpty = errors.New("empty")

type Presenter struct {
	factory factory
}

func NewPresenter(f factory) *Presenter {
	return &Presenter{
		factory: f,
	}
}

func (p *Presenter) Present(files []model.File) []model.TestFile {
	testFiles := make([]model.TestFile, 0, len(files))

	for _, file := range files {
		testFile, err := p.presentFile(file)
		if errors.Is(err, errEmpty) {
			continue
		}

		testFiles = append(testFiles, testFile)
	}

	return testFiles
}

func (p *Presenter) presentFile(file model.File) (model.TestFile, error) {
	imports := model.NewImports().SetProjectModuleName(file.Package.ProjectModuleName)
	imports = imports.Append(fmt.Sprintf(". \"%s\"", file.Package.Path)) // селф-импорт

	tests := make([]model.Test, 0, len(file.Functions))
	for _, function := range file.Functions {
		presenter, err := p.factory.CreateTestPresenter(function)
		if errors.Is(err, model.ErrUnsupported) {
			continue
		}

		imports = imports.Append(function.NeededImports.Get()...)       // импорты, которые нужны для функции
		imports = imports.Append(presenter.PresentImports(function)...) // импорты, которые нужны для теста

		tests = append(tests, getTest(file, function, presenter))
	}

	if len(tests) == 0 {
		return model.TestFile{}, errEmpty
	}

	return model.TestFile{
		Path: file.Path.ToTest().String(),
		Package: model.Package{
			Name:              fmt.Sprintf("%s_test", file.Package.Name),
			ProjectModuleName: file.Package.ProjectModuleName,
		},
		Imports: imports,
		Tests:   tests,
	}, nil
}

func getTest(file model.File, function model.Function, presenter TestPresenter) model.Test {
	testName := fmt.Sprintf("Test%s", function.Name)

	if function.Receiver != nil {
		testName = fmt.Sprintf("Test%s_%s", function.Receiver.Name, function.Name)
	}

	presentedStructure := presenter.PresentStructure(function)
	presentedLoop := presenter.PresentLoop(function)

	return model.Test{
		Name: testName,
		Source: fmt.Sprintf(
			template,
			testName,
			presentedStructure,
			presentTestcase(file),
			presentedLoop,
		),
	}
}
