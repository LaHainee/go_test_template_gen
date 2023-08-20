package create

import (
	"io"
	"os"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

type Repository struct {
	source Source
}

func NewRepository(source Source) *Repository {
	return &Repository{
		source: source,
	}
}

func (r *Repository) Create(files []model.TestFile) error {
	for _, file := range files {
		err := r.create(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) create(testFile model.TestFile) error {
	file, err := os.OpenFile(testFile.Path, os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	rows := strings.Split(string(content), "\n")

	rows, err = r.source.Add(rows, testFile)
	if err != nil {
		return err
	}

	return os.WriteFile(testFile.Path, []byte(strings.Join(rows, "\n")), 0o644)
}
