package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FilePath string

func (path FilePath) ToTest() FilePath {
	return FilePath(fmt.Sprintf("%s/%s_test.go", path.Dir(), path.FileNameWithoutExtension()))
}

func (path FilePath) String() string {
	return string(path)
}

func (path FilePath) Dir() string {
	return filepath.Dir(path.String())
}

func (path FilePath) Base() string {
	return filepath.Base(path.String())
}

func (path FilePath) FileNameWithoutExtension() string {
	filename := path.Base()
	extension := filepath.Ext(path.String())

	return filename[:len(filename)-len(extension)]
}

func (path FilePath) IsGolangSource() bool {
	extension := filepath.Ext(path.String())

	return extension == ".go"
}

func (path FilePath) DirectoryFilePaths() ([]FilePath, error) {
	// Необходимо для обработки кейса, когда на path это /User/vaershov, т.е. в конце нет слеша
	// Чтобы не вызывать os.Stat() реализовано решение с проверкой длины расширения
	directory := path.String()
	if len(filepath.Ext(path.String())) > 0 {
		directory = filepath.Dir(path.String())
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	filePaths := make([]FilePath, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(directory, file.Name())

		filePaths = append(filePaths, FilePath(filePath))
	}

	return filePaths, nil
}

func (path FilePath) IsMock() bool {
	fileName := filepath.Base(path.String())

	return strings.Contains(fileName, "mock")
}

func (path FilePath) IsTest() bool {
	fileName := filepath.Base(path.String())

	return strings.HasSuffix(fileName, "_test.go") && !path.IsMock()
}
