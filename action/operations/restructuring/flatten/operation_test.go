package flatten

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1dArr(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]int{1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func Test2dArr(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][]interface{}{{1},{3},{4},{5},{6}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	assert.NotNil(t, result)

}

func Test3dArr(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][][]interface{}{{{1,4}},{{3}},{{4}},{{5}},{{6}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)
	assert.NotNil(t, result)

}