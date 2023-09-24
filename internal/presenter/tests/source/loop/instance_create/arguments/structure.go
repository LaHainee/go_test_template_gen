package arguments

import (
	"fmt"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	"github.com/LaHainee/go_test_template_gen/internal/presenter/tests/source/loop"
	"strings"
)

type Structure struct{}

func NewStructure() *Structure {
	return &Structure{}
}

func (s *Structure) Get(function model.Function) string {
	arguments := make([]string, 0)

	for _, dependency := range function.Receiver.Dependencies {
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
