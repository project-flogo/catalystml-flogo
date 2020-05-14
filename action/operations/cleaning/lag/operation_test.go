package lag

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"
	// "github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestMapIn(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["table"] = map[string]interface{}{"state": []interface{}{"CA", "NC", "TX"}}
	inputs["lagnum"] = 1
	inputs["col"] = "state"

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	assert.Nil(t, err)
	fmt.Println("OUTPUT:", output)
}

func TestTableIn(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["table"] = []interface{}{[]interface{}{1, 2, 3, 4, 5}, []interface{}{"a", "b", "c", "d", "e"}}
	inputs["lagnum"] = 1
	inputs["col"] = 0

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	assert.Nil(t, err)
	fmt.Println("OUTPUT:", output)
}
