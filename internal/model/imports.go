package model

import (
	"fmt"
	"sort"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/util/set"
)

var thirdPartyPrefixes = []string{
	"github.com",
	"go.avito.ru",
	"go.octolab.org",
	"go.uber.org",
	"golang.org",
	"gopkg.in",
}

type Imports struct {
	stdlib            set.Set[string]
	thirdParty        set.Set[string]
	internal          set.Set[string]
	projectModuleName string
}

func NewImports() Imports {
	return Imports{
		stdlib:     set.New[string](),
		thirdParty: set.New[string](),
		internal:   set.New[string](),
	}
}

// Search - найти импорт по полученному пакету
func (imports Imports) Search(pkg string) (string, error) {
	importsList := imports.Get()

	for _, concreteImport := range importsList {
		// Разобьем импорт
		parts := strings.Split(concreteImport, " ")

		if len(parts) == 0 {
			continue
		}

		// Тут будет либо алиас для импорта, например, context, либо полный путь /str-activator/internal/models
		pathOrAlias := strings.Trim(parts[0], "\"")

		if len(pathOrAlias) < len(pkg) {
			continue
		}

		if pathOrAlias[len(pathOrAlias)-len(pkg):] == pkg {
			return concreteImport, nil
		}
	}

	return "", ErrNotFound
}

func (imports Imports) SetProjectModuleName(moduleName string) Imports {
	imports.projectModuleName = moduleName

	// После того как задан moduleName нужно заново заполнить блоки
	values := imports.Get()
	imports.stdlib = set.New[string]()
	imports.thirdParty = set.New[string]()
	imports.internal = set.New[string]()

	return imports.Append(values...)
}

func (imports Imports) Append(values ...string) Imports {
	for _, val := range values {
		if imports.isInternal(val) {
			imports.internal.Add(val)
			continue
		}

		if imports.isThirdParty(val) {
			imports.thirdParty.Add(val)
			continue
		}

		// Импорты, которые не являются ни 3d-party, ни внутренними относим к импортам из stdlib
		imports.stdlib.Add(val)
	}

	return imports
}

func (imports Imports) PresentReformatted() []string {
	rows := make([]string, 0)

	stdlibImports := imports.stdlib.Values()
	thirdPartyImports := imports.thirdParty.Values()
	internalImports := imports.internal.Values()

	if len(stdlibImports) > 0 {
		sort.Sort(importsSlice(stdlibImports))
		rows = append(rows, stdlibImports...) // 1-й блок импорты из стандартной библиотеки
		rows = append(rows, "")
	}

	if len(thirdPartyImports) > 0 {
		sort.Sort(importsSlice(thirdPartyImports))
		rows = append(rows, thirdPartyImports...) // 2-й блок импорты внешних библиотек
		rows = append(rows, "")
	}

	if len(internalImports) > 0 {
		sort.Sort(importsSlice(internalImports))
		rows = append(rows, internalImports...) // 3-й блок внутренние импорты
	}

	// Перед каждой непустой строкой добавим табуляцию
	for i, row := range rows {
		if len(row) == 0 {
			continue
		}

		rows[i] = "\t" + row
	}

	return rows
}

func (imports Imports) Get() []string {
	list := make([]string, 0)

	list = append(list, imports.stdlib.Values()...)
	list = append(list, imports.thirdParty.Values()...)
	list = append(list, imports.internal.Values()...)

	return list
}

func (imports Imports) isThirdParty(value string) bool {
	for _, prefix := range thirdPartyPrefixes {
		// Искать будем строку вида "github.com
		hasPrefix := strings.Contains(value, fmt.Sprintf("\"%s/", prefix))
		if !hasPrefix {
			continue
		}

		// Дополнительно проверим, что импорт не является внутренним
		return !imports.isInternal(value)
	}

	return false
}

func (imports Imports) isInternal(value string) bool {
	// Если не задан module name проекта, то невозможно определить, что импорт относится к внутренним импортам
	if len(imports.projectModuleName) == 0 {
		return false
	}

	return strings.Contains(value, fmt.Sprintf("\"%s/", imports.projectModuleName))
}

// Сортировка импортов
type importsSlice []string

func (s importsSlice) Len() int {
	return len(s)
}

func (s importsSlice) Less(i, j int) bool {
	partsI := strings.Split(s[i], " ")
	partsJ := strings.Split(s[j], " ")

	toCompareI := partsI[0]
	toCompareJ := partsJ[0]

	if len(partsI) > 1 {
		toCompareI = partsI[1]
	}

	if len(partsJ) > 1 {
		toCompareJ = partsJ[1]
	}

	return toCompareI < toCompareJ
}

func (s importsSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
