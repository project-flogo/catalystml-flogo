package concatMap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/project-flogo/core/support/log"
	"fmt"
)

func TestMultipleDEval(t *testing.T) {
	temp1 := make(map[string]interface{})
	temp2 := make(map[string]interface{})
	temp1["a"] = "Abc"
	temp1["x"] = "XYZ"

	temp2["a"] = "Abc"
	temp2["x"] = "XYZ"

	inputs := make(map[string]interface{})

	inputs["data"] = []interface{}{temp1, temp2}
	opt := &Operation{logger: log.RootLogger()}

	result, err := opt.Eval(inputs)
	fmt.Println(result)
	assert.Nil(t, err)
}

func TestSingleDEval(t *testing.T) {
	temp1 := make(map[string]interface{})
	
	temp1["a"] = "Abc"
	temp1["x"] = "XYZ"


	inputs := make(map[string]interface{})

	inputs["data"] = []interface{}{temp1}
	opt := &Operation{logger: log.RootLogger()}

	result, err := opt.Eval(inputs)
	fmt.Println(result)
	assert.Nil(t, err)
}