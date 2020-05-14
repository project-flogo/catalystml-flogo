package normalize

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSingleDEval(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]int{1, 3, 45, 5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"] = 2

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	out, err := opt.Eval(inputs)
	fmt.Println(out)
	var test []float32
	for _, val := range out.([]interface{}) {
		test = append(test, val.(float32))
	}
	assert.Nil(t, err)
	assert.Equal(t, []float32{-0.055555556, 0.055555556, 2.3888888, 0.16666667}, test)

}
func TestSingleDEvalDiff(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]float32{1, 3, 45, 5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"] = 2

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	out, err := opt.Eval(inputs)
	fmt.Println(out)
	var test []float32
	for _, val := range out.([]interface{}) {
		test = append(test, val.(float32))
	}
	assert.Nil(t, err)
	assert.Equal(t, []float32{-0.055555556, 0.055555556, 2.3888888, 0.16666667}, test)

}
func TestMultiple2DEval(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][]int{{1}, {3}, {45, 56}, {5}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"] = 2

	out, err := opt.Eval(inputs)
	fmt.Println(out)

	assert.Nil(t, err)

}
func TestMultiple3DEval(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][][]uint8{{{1}, {3, 5}}, {{3}}, {{45, 56}}, {{5}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"] = 2

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestMultiple3DEvalU16(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][][]uint16{{{1}, {3, 5}}, {{3}}, {{45, 56}}, {{5}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"] = 2

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
