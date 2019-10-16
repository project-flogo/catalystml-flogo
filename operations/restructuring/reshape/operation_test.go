package reshape

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]int{1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{2,2}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}
func TestSingleRow(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]int{1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{4,-1}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}
func TestMultiRe(t *testing.T) {
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][]int{{1,1},{2,2},{3,3},{4,4}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	
	inputs["shape"]= []int{2,4}
	
	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}