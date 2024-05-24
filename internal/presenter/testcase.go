package presenter

import (
	"fmt"
	"github.com/LaHainee/go_test_template_gen/internal/model"
	"strings"
)

const defaultTestCase = "{},"

var tmsProjectToTagID = map[string]string{
	"service-str-activator":         "19013",
	"service-str-verification":      "19014",
	"service-str-notifier":          "19045",
	"service-str-quality":           "19636",
	"service-str-admin-composition": "19637",
	"service-str-booking-storage":   "23852",
	"service-str-booking-gateway":   "24249",
	"service-calendar-platform":     "23901",
}

func presentTestcase(file model.File) string {
	tagID, ok := tmsProjectToTagID[file.ProjectName]
	if !ok {
		return defaultTestCase
	}

	rows := []string{
		fmt.Sprintf("// tagID %s <Test name>", tagID),
		"// tmsTestType unit",
		defaultTestCase,
	}

	return strings.Join(rows, "\n\t\t")
}
