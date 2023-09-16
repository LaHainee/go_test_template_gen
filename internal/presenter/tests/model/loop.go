package model

import (
	"fmt"
	"strings"

	"github.com/LaHainee/go_test_template_gen/internal/util/pointer"
)

type Loop struct {
	MockCreate      []Statement
	MockPrepare     []string
	InstanceCreate  *Statement
	FunctionCall    Statement
	ExpectationCall string
}

func (loop Loop) Present() string {
	rows := make([]string, 0)

	mockCreate := loop.PresentMockCreate()
	if len(mockCreate) != 0 {
		rows = append(rows, mockCreate...)
		rows = append(rows, "")
	}

	mockPrepare := loop.PresentMockPrepare()
	if len(mockPrepare) != 0 {
		rows = append(rows, mockPrepare...)
		rows = append(rows, "")
	}

	instanceInit := loop.PresentInstanceInit()
	if instanceInit != nil {
		rows = append(rows, pointer.Val(instanceInit))
	}

	rows = append(rows, loop.PresentFunctionCall())
	rows = append(rows, "")
	rows = append(rows, loop.PresentExpectationCall())

	return strings.Join(rows, "\n\t\t\t")
}

func (loop Loop) PresentMockCreate() []string {
	rows := make([]string, 0, len(loop.MockCreate))

	for _, statement := range loop.MockCreate {
		rows = append(rows, fmt.Sprintf("%s := %s", statement.Lhs, statement.Rhs))
	}

	return rows
}

func (loop Loop) PresentMockPrepare() []string {
	return loop.MockPrepare
}

func (loop Loop) PresentInstanceInit() *string {
	if loop.InstanceCreate == nil {
		return nil
	}

	return pointer.To(fmt.Sprintf("%s := %s\n", loop.InstanceCreate.Lhs, loop.InstanceCreate.Rhs))
}

func (loop Loop) PresentFunctionCall() string {
	return fmt.Sprintf("%s := %s", loop.FunctionCall.Lhs, loop.FunctionCall.Rhs)
}

func (loop Loop) PresentExpectationCall() string {
	return loop.ExpectationCall
}
