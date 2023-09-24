package model

import (
	"strings"
	"unicode"
)

type Function struct {
	Name            string
	Receiver        *Structure
	NeededImports   Imports
	Return          *Return
	InputArguments  []Argument
	OutputArguments []Argument
}

func (function Function) HasInterfaceDependencies() bool {
	if function.Receiver == nil {
		return false
	}

	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface != nil {
			return true
		}
	}

	return false
}

func (function Function) IsPrivate() bool {
	return unicode.IsLower(rune(function.Name[0]))
}

func (function Function) IsConstructor() bool {
	return strings.HasPrefix(function.Name, "New")
}

type Structure struct {
	Name         string
	Constructor  *Function
	Dependencies []Dependency
}
