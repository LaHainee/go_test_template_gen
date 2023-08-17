package model

const (
	ArgumentError   = "error"
	ArgumentContext = "context.Context"
)

type File struct {
	Path      FilePath
	Imports   Imports
	Package   Package
	Functions []Function
}

type Package struct {
	Name              string
	Path              string
	ProjectModuleName string // module из go.mod
}

func (f File) IsMock() bool {
	return f.Path.IsMock()
}

func (f File) IsTest() bool {
	return f.Path.IsTest()
}
