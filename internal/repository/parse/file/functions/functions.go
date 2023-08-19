package functions

import "github.com/LaHainee/go_test_template_gen/internal/model"

func SetConstructors(files []model.File) {
	allFunctions := getFunctions(files)

	for _, file := range files {
		for _, function := range file.Functions {
			if function.Receiver == nil {
				continue
			}

			constructor, err := model.Functions(allFunctions).LookupByOutputArgument(function.Receiver.Name)
			if err != nil {
				continue
			}

			function.Receiver.ConstructorName = &constructor.Name
		}
	}
}

func getFunctions(files []model.File) []model.Function {
	functions := make([]model.Function, 0)

	for _, file := range files {
		functions = append(functions, file.Functions...)
	}

	return functions
}
