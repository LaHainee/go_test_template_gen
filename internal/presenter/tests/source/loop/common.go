package loop

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/model"
)

func MockName(dependency model.Dependency) string {
	return fmt.Sprintf("mock%s", capitalizeFirst(dependency.Type))
}

func capitalizeFirst(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
