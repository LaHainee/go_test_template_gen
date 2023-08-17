package model

import (
	"strings"
	"unicode"
)

type Function struct {
	Name            string
	Receiver        *Structure
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

func (functions Functions) LookupByOutputArgument(argumentName string) (Function, error) {
	for _, function := range functions {
		for _, argument := range function.OutputArguments {
			if argument.Dereference() == argumentName {
				return function, nil
			}
		}
	}

	return Function{}, ErrNotFound
}

type Argument struct {
	Name *string
	Type string
}

func (argument Argument) Dereference() string {
	return strings.TrimLeft(argument.Type, "*")
}

type Structure struct {
	Name            string
	ConstructorName *string
	Dependencies    []Dependency
}
