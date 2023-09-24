package model

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

func (f File) LookupFunctionByOutput(output string) (*Function, error) {
	for _, function := range f.Functions {
		for _, argument := range function.OutputArguments {
			if argument.Dereference() == output {
				return &function, nil
			}
		}
	}

	return nil, ErrNotFound
}
