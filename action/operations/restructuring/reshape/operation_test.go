package reshape

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]int{1, 3, 45, 5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{2, 2}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}

func TestAllNegOne(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]int{1, 3, 45, 5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{-1, -1, -1}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	_, err = opt.Eval(inputs)

	assert.NotNil(t, err)
	fmt.Println("Expected error . . .", err)

}

func TestSingleRow(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]int{1, 3, 45, 5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{4, -1}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}
func TestInvalidGuess(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([]int{1, 3, 45, 5, 9})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["shape"] = []int{2, -1}

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}

	_, err = opt.Eval(inputs)

	assert.NotNil(t, err)
	fmt.Println("Expected error . . .", err)

}

func TestMultiRe(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	inputs["shape"] = []int{2, 4}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}

func Test3DIn(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][][]int{{{1}, {1}}, {{2}, {2}}, {{3}, {3}}, {{4}, {4}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	inputs["shape"] = []int{2, 4}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}

func Test3DOut(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	inputs["shape"] = []int{2, 4, -1}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}

func Test3DOut2(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	inputs["shape"] = []int{2, -1, 4}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}

func Test4DOut(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	val, err := coerce.ToArray([][][][]int{{{{1, 2}}, {{1, 2}}}, {{{2, 3}}, {{2, 3}}}, {{{3, 4}}, {{3, 4}}}, {{{4, 5}}, {{4, 5}}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val

	inputs["shape"] = []int{1, 1, 4, -1}

	result, err := opt.Eval(inputs)

	assert.Nil(t, err)
	fmt.Println("Result...", result)

}
