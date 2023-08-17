package model

import (
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

type Structure struct {
	TestName          Statement
	FunctionArguments []Statement
	Prepare           *Statement
	Expectation       Statement
}

func (structure *Structure) Present() string {
	structure.reformat()

	rows := make([]string, 0)

	prepare := structure.PresentPrepare()

	rows = append(rows, structure.PresentTestName())
	rows = append(rows, structure.PresentFunctionArguments()...)
	if prepare != nil {
		rows = append(rows, pointer.Val(prepare))
	}
	rows = append(rows, structure.PresentExpectation())

	return strings.Join(rows, "\n\t\t")
}

func (structure *Structure) PresentTestName() string {
	return structure.TestName.Lhs + " " + structure.TestName.Rhs
}

func (structure *Structure) PresentFunctionArguments() []string {
	rows := make([]string, 0, len(structure.FunctionArguments))

	for _, statement := range structure.FunctionArguments {
		rows = append(rows, statement.Lhs+" "+statement.Rhs)
	}

	return rows
}

func (structure *Structure) PresentPrepare() *string {
	if structure.Prepare == nil {
		return nil
	}

	return pointer.To(structure.Prepare.Lhs + " " + structure.Prepare.Rhs)
}

func (structure *Structure) PresentExpectation() string {
	return structure.Expectation.Lhs + " " + structure.Expectation.Rhs
}

// reformat - выравнивание по одному уровню
func (structure *Structure) reformat() {
	maxFieldLength := structure.getMaxFieldLength()

	structure.TestName.Rhs = strings.Repeat(" ", maxFieldLength-len(structure.TestName.Lhs)) + structure.TestName.Rhs

	for i := range structure.FunctionArguments {
		structure.FunctionArguments[i].Rhs = strings.Repeat(" ", maxFieldLength-len(structure.FunctionArguments[i].Lhs)) + structure.FunctionArguments[i].Rhs
	}

	if structure.Prepare != nil {
		structure.Prepare.Rhs = strings.Repeat(" ", maxFieldLength-len(structure.Prepare.Lhs)) + structure.Prepare.Rhs
	}

	structure.Expectation.Rhs = strings.Repeat(" ", maxFieldLength-len(structure.Expectation.Lhs)) + structure.Expectation.Rhs
}

func (structure *Structure) getMaxFieldLength() int {
	var maxFieldLength int

	if len(structure.TestName.Lhs) > maxFieldLength {
		maxFieldLength = len(structure.TestName.Lhs)
	}

	if len(structure.Expectation.Lhs) > maxFieldLength {
		maxFieldLength = len(structure.Expectation.Lhs)
	}

	if structure.Prepare != nil && len(structure.Prepare.Lhs) > maxFieldLength {
		maxFieldLength = len(structure.Prepare.Lhs)
	}

	for _, arg := range structure.FunctionArguments {
		if len(arg.Lhs) > maxFieldLength {
			maxFieldLength = len(arg.Lhs)
		}
	}

	return maxFieldLength
}
