# Описание
Библиотека для генерации шаблонов Unit-тестов на Golang
```go
import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	. "go.avito.ru/str/service-str-quality/internal/gateway/item_aggregator"
)

func TestGateway_GetOwnerIDs(t *testing.T) {
	t.Parallel()
	
	tests := []struct {
		name         string
		itemIDs      []int64
		prepare      func(transport *Mocktransport)
		expectations func(t assert.TestingT, got map[int64]int64, err error)
	}{
		{},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockTransport := NewMocktransport(ctrl)
			
			tc.prepare(mockTransport)
			
			instance := NewGateway(mockTransport)

			out, err := instance.GetOwnerIDs(context.Background(), tc.itemIDs)
			
			tc.expectations(t, out, err)
		})
	}
}
```

# Особенности

- Для создания моков используется `gomock`
- Для приватных функций тесты не генерируются

# Установка
1) В корне проекта создать директорию `tools`
2) Перейти в созданную директорию и выполнить команду `go mod init tools`
3) Скачать библиотеку `go get github.com/LaHainee/go_test_template_gen`
4) Создать файл `tools.go`
```
//go:build tools
// +build tools

//go:generate go build -mod=mod -o ../bin/testgen github.com/LaHainee/go_test_template_gen/cmd

package tools

import _ "github.com/LaHainee/go_test_template_gen/cmd"
```
5) Выполнить команду `go mod vendor`
6) Сгенерировать бинарь, который будет сохранен в директорию `bin` в корне проекта. Чтобы игнорировать эту директорию можно внутри нее создать  `.gitignore`
```
*
!.gitignore
```
7) Добавить команду в `Makefile`
```
testgen:
	./bin/testgen -path "$(CURDIR)/$(PATH)"
```

# Использование
Выполнить команду
```text
make testgen PATH=internal/gateway/item_aggregator
```
Необходимо указывать относительный путь к файлу или директории. Если указать путь к директории, то генерация будет
происходить для всех файлов внутри директории. Если указать путь к файлу, то генерация будет происходить только для
указанного файла