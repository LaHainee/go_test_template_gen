package file

import (
	"github.com/LaHainee/go_test_template_gen/internal/model"
)

func bindConstructorsToReceivers(files []model.File) {
	for _, file := range files {
		for _, function := range file.Functions {
			if function.Receiver == nil {
				continue
			}

			constructor, err := file.LookupFunctionByOutput(function.Receiver.Name)
			if err != nil {
				continue
			}

			function.Receiver.Constructor = constructor
		}
	}
}
