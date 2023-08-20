package file

import (
	"errors"
	"fmt"
	"go/ast"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	facade "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast"
	"github.com/LaHainee/go_test_template_gen/internal/util/slice"
)

type Source struct {
	next facade.Source
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) SetNext(next facade.Source) {
	s.next = next
}

func (s *Source) Extend(filePath model.FilePath, astFile *ast.File, file *model.File) error {
	pkg, err := getPackage(string(filePath), astFile)
	if err != nil {
		return err
	}

	file.Path = filePath
	file.Package = pkg

	if s.next != nil {
		return s.next.Extend(filePath, astFile, file)
	}
	return nil
}

func getPackage(filePath string, file *ast.File) (model.Package, error) {
	if file.Name == nil {
		return model.Package{}, errors.New("corrupted file")
	}

	projectRootPath, err := getProjectRootPath(filePath)
	if err != nil {
		return model.Package{}, err
	}

	packagePath, err := getPackagePath(projectRootPath, filePath)
	if err != nil {
		return model.Package{}, err
	}

	projectModuleName, err := getProjectModuleName(projectRootPath)
	if err != nil {
		return model.Package{}, err
	}

	return model.Package{
		Name:              file.Name.Name,
		Path:              packagePath,
		ProjectModuleName: projectModuleName,
	}, nil
}

func getPackagePath(rootPath, filePath string) (string, error) {
	filePath = filepath.Dir(filePath)

	pathParts := make([]string, 0)

	for {
		if filePath == "/" {
			return "", fmt.Errorf("%s not found", rootPath)
		}

		if filePath == rootPath {
			projectModuleName, err := getProjectModuleName(rootPath)
			if err != nil {
				return "", err
			}

			pathParts = append(pathParts, projectModuleName)

			return strings.Join(slice.Reverse(pathParts), "/"), nil
		}

		currentDirectory := filepath.Base(filePath)
		directoryPath := filepath.Dir(filePath)
		filePath = directoryPath

		pathParts = append(pathParts, currentDirectory)
	}
}

func getProjectModuleName(rootPath string) (string, error) {
	file, err := os.Open(filepath.Join(rootPath, "go.mod"))
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if !strings.HasPrefix(line, "module") {
			continue
		}

		return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
	}

	return "", errors.New("module name not found in go.mod")
}

/*
getProjectRootPath – получить путь к корню проекта. Требуется для последующего парсинга module из go.mod, а также
получения импорта до текущего файла

Например /User/vaershov/Desktop/work/service-str-activator/
*/
func getProjectRootPath(filePath string) (string, error) {
	for {
		if filePath == "/" {
			return "", errors.New("file go.mod not found")
		}

		directoryPath := filepath.Dir(filePath)
		filePath = directoryPath

		var rootPath string

		err := filepath.WalkDir(directoryPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.Name() == "go.mod" {
				rootPath = directoryPath
				return nil
			}

			return nil
		})
		if err != nil {
			return "", err
		}

		if rootPath != "" {
			return rootPath, nil
		}
	}
}
