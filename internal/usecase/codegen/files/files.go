package files

import (
	"errors"
	"os"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Getter struct {
	functions   functionsRepository
	byDirectory source
	byFilepath  source
}

func NewGetter(fr functionsRepository, byDirectory, byFilepath source) *Getter {
	return &Getter{
		functions:   fr,
		byDirectory: byDirectory,
		byFilepath:  byFilepath,
	}
}

func (g *Getter) Get(path string) ([]model.File, error) {
	files, err := g.getFiles(path)
	if err != nil {
		return nil, err
	}

	filesToCover := filterFilesToCover(files)

	files = make([]model.File, 0)

	for _, file := range filesToCover {
		uncoveredFunctions, err := g.functions.GetUncovered(file.Functions, file.Path.ToTest())
		if errors.Is(err, model.ErrNotFound) {
			// файл не найден, значит его надо покрыть полностью
			files = append(files, file)
			continue
		}
		if err != nil {
			return nil, err
		}

		functionToCover := filterFunctionToCover(uncoveredFunctions)

		if len(functionToCover) == 0 {
			continue
		}

		file.Functions = functionToCover

		files = append(files, file)
	}

	return files, nil
}

func (g *Getter) getFiles(path string) ([]model.File, error) {
	file, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if file.IsDir() {
		return g.byDirectory.Get(model.FilePath(path))
	}
	return g.byFilepath.Get(model.FilePath(path))
}

func filterFunctionToCover(functions []model.Function) []model.Function {
	toCover := make([]model.Function, 0, len(functions))

	for _, function := range functions {
		if function.IsConstructor() {
			continue
		}

		toCover = append(toCover, function)
	}

	return toCover
}

func filterFilesToCover(files []model.File) []model.File {
	toCover := make([]model.File, 0, len(files))

	for _, file := range files {
		if file.IsMock() || file.IsTest() {
			continue
		}

		if len(file.Functions) == 0 {
			continue
		}

		toCover = append(toCover, file)
	}

	return toCover
}
