package arguments

import (
	"fmt"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/source/loop"
	"strings"
)

type Constructor struct{}

func NewConstructor() *Constructor {
	return &Constructor{}
}

func (c *Constructor) Get(function model.Function) string {
	arguments := make([]string, 0)

	dependencies := make(map[string]model.Dependency)
	for _, dependency := range function.Receiver.Dependencies {
		dependencies[dependency.Name] = dependency
	}

	for _, argument := range function.Receiver.Constructor.InputArguments {
		if argument.Name == nil {
			continue
		}

		dependencyName := function.Receiver.Constructor.Return.Structure.ArgumentBindings[*argument.Name]

		dependency := dependencies[dependencyName]

		if dependency.Interface == nil {
			arguments = append(arguments, fmt.Sprintf("tc.%s", dependency.Name))
			continue
		}

		if dependency.IsLogger() {
			arguments = append(arguments, "log")
			continue
		}

		arguments = append(arguments, loop.MockName(dependency))
	}

	return strings.Join(arguments, ", ")
}
