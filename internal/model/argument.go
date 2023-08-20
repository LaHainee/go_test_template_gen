package model

import (
	"strings"
)

const (
	ArgumentError   = "error"
	ArgumentContext = "context.Context"
)

type Argument struct {
	Name    *string
	Type    string
	Package *string
}

func (argument Argument) IsPointer() bool {
	if len(argument.Type) == 0 {
		return false
	}

	return argument.Type[0] == '*'
}

func (argument Argument) NameContains(substr string) bool {
	if argument.Name == nil {
		return false
	}

	name := *argument.Name

	substr = strings.ToLower(substr)
	for i := 0; i <= len(name)-len(substr); i++ {
		if strings.ToLower(name[i:i+len(substr)]) == substr {
			return true
		}
	}
	return false
}

func (argument Argument) Is(typ string) bool {
	return argument.Type == typ
}

func (argument Argument) Dereference() string {
	return strings.TrimLeft(argument.Type, "*")
}
