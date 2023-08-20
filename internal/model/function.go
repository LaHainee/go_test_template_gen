package model

import (
	"strings"
	"unicode"
)

type Function struct {
	Name            string
	Receiver        *Structure
	NeededImports   Imports
	InputArguments  []Argument
	OutputArguments []Argument
}
type Functions []Function

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

// LookupByOutputArgument - найти функцию, которая возвращает аргумент с полученным типом
func (functions Functions) LookupByOutputArgument(argumentType string) (Function, error) {
	for _, function := range functions {
		for _, argument := range function.OutputArguments {

			// Аргумент может быть указателем, поэтому необходимо его разименовать
			if argument.Dereference() == argumentType {
				return function, nil
			}
		}
	}

	return Function{}, ErrNotFound
}

type Structure struct {
	Name            string
	ConstructorName *string
	Dependencies    []Dependency
}
