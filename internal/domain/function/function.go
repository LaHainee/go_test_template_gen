package function

import "github.com/LaHainee/go_test_template_gen/internal/model"

const (
	TypeMethodWithInterfaceDeps    = "method_with_interface_deps"
	TypeMethodWithoutInterfaceDeps = "method_without_interface_deps"
	TypeFunction                   = "function"
)

func GetFunctionType(function model.Function) string {
	if function.Receiver == nil {
		return TypeFunction
	}

	return getMethodType(function)
}

func getMethodType(function model.Function) string {
	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil {
			continue
		}

		return TypeMethodWithInterfaceDeps
	}

	return TypeMethodWithoutInterfaceDeps
}
