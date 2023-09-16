package function

import "github.com/LaHainee/go_test_template_gen/internal/model"

const (
	TypePublicMethodWithInterfaceDeps     = "public_method_with_interface_deps"
	TypePublicMethodWithoutInterfaceDeps  = "public_method_without_interface_deps"
	TypePrivateMethodWithInterfaceDeps    = "private_method_with_interface_deps"
	TypePrivateMethodWithoutInterfaceDeps = "private_method_without_interface_deps"
	TypePublicFunction                    = "public_function"
	TypePrivateFunction                   = "private_function"
)

func GetFunctionType(function model.Function) string {
	if function.Receiver == nil {
		// Функция, приватная или публичная
		return getFunctionType(function)
	}

	// Метод:
	// 1) С интерфейсными зависимостями
	// 		1.1) Приватный
	// 		1.2) Публичный
	// 2) Без интерфейсных зависимостей
	// 		2.1) Приватный
	// 		2.2) Публичный
	return getMethodType(function)
}

func getMethodType(function model.Function) string {
	for _, dependency := range function.Receiver.Dependencies {
		if dependency.Interface == nil {
			continue
		}

		// Нашли интерфейсную зависимость
		if function.IsPrivate() {
			return TypePrivateMethodWithInterfaceDeps
		}
		return TypePublicMethodWithInterfaceDeps
	}

	if function.IsPrivate() {
		return TypePrivateMethodWithoutInterfaceDeps
	}
	return TypePublicMethodWithoutInterfaceDeps
}

func getFunctionType(function model.Function) string {
	if function.IsPrivate() {
		return TypePrivateFunction
	}

	return TypePublicFunction
}
