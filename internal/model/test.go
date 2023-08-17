package model

type TestFile struct {
	Path    string  // путь до файла
	Package Package // пакет
	Imports Imports // импорты
	Tests   []Test  // тесты
}

type Test struct {
	Name   string
	Source string
}
