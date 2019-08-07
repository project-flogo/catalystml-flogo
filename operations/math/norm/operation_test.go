package norm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
)

func TestSingleDEval(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]int{1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestMultipleDEval(t *testing.T) {
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][]int{{1},{3},{45,56},{5}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	//input := &Input{Array: [][]float32{{1.0}, {2.0}, {3.6}, {4}}}
	

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
