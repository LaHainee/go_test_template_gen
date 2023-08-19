package file

import (
	"os"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/file/functions"
)

type Parser struct {
	filesystem filesystem
}

func NewParser(fs filesystem) *Parser {
	return &Parser{
		filesystem: fs,
	}
}

func (p *Parser) ParseDirectory(directoryPath model.FilePath) ([]model.File, error) {
	filePaths, err := directoryPath.DirectoryFilePaths()
	if err != nil {
		return nil, err
	}

	files := make([]model.File, 0, len(filePaths))

	for _, filePath := range filePaths {
		file, err := p.Parse(filePath)
		if err != nil {
			// При парсинге директории намеренно пропускаем обработку ошибок, чтобы не валить генерацию
			continue
		}

		files = append(files, file)
	}

	// Приходится вызывать после того как спарсили все файлы, поскольку конструктор может быть объявлен в другом файле,
	// который еще не спаршен
	functions.SetConstructors(files)

	return files, nil
}

func (p *Parser) Parse(filePath model.FilePath) (model.File, error) {
	if !filePath.IsGolangSource() {
		return model.File{}, model.ErrInvalidFile
	}

	if _, err := os.Stat(filePath.String()); os.IsNotExist(err) {
		return model.File{}, model.ErrNotFound
	}

	file, err := p.filesystem.Parse(filePath)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}
