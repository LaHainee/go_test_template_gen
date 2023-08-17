package model

import "github.com/LaHainee/go_test_template_gen/internal/util/set"

var (
	metricsDepNamings = set.New("metric", "metrics", "stat", "stats")
	loggerDepNamings  = set.New("log", "logger")
)

type Dependency struct {
	Name      string
	Type      string
	Interface *Interface
}

type Interface struct {
	Functions []Function
}

func (d Dependency) IsGrace() bool {
	return d.Type == "grace"
}

func (d Dependency) IsMetric() bool {
	return metricsDepNamings.Contains(d.Type)
}

func (d Dependency) IsLogger() bool {
	return loggerDepNamings.Contains(d.Type)
}
