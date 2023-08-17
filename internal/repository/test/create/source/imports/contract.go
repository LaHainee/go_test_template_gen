package imports

import "github.com/LaHainee/go_test_template_gen/internal/model"

type filesystem interface {
	Parse(filePath model.FilePath) (model.File, error)
}
