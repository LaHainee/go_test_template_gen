package imports

import (
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/util/slice"
)

func (s *Source) update(rows []string, num int, testFile model.TestFile) ([]string, error) {
	// Чтобы не парсить вручную строки
	file, err := s.filesystem.Parse(model.FilePath(testFile.Path))
	if err != nil {
		return rows, err
	}

	// К существующим импортам добавим новые
	imports := file.Imports
	imports = imports.Append(testFile.Imports.Get()...)

	// Сохраняем
	rows = slice.Insert(rows, num, "import (")
	num++
	for _, concreteImport := range imports.PresentReformatted() {
		rows = slice.Insert(rows, num, concreteImport)
		num++
	}
	rows = slice.Insert(rows, num, ")")
	num++

	// Удаляем старый блок импортов
	if !strings.Contains(rows[num], "import (") { // Однострочный блок импорта
		return slice.Remove(rows, num), nil
	}

	for rows[num] != ")" {
		rows = slice.Remove(rows, num)
	}
	rows = slice.Remove(rows, num)

	return rows, nil
}
