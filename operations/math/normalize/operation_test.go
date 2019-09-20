package normalize
import (
	"testing"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/coerce"
	"github.com/stretchr/testify/assert"
)
func TestSingleDEval(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]int{1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"]= 2

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func TestSingleDEvalDiff(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([]interface{}{"Age",1,3,45,5})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"]= 2

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func TestMultiple2DEval(t *testing.T) {
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][]int{{1},{3},{45,56},{5}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"]= 2
	
	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func TestMultiple3DEval(t *testing.T) {
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][][]uint8{{{1},{3,5}},{{3}},{{45,56}},{{5}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"]= 2
	
	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestMultiple3DEvalU16(t *testing.T) {
	opt := &Operation{logger : log.RootLogger() }
	val, err := coerce.ToArray([][][]uint16{{{1},{3,5}},{{3}},{{45,56}},{{5}}})
	inputs := make(map[string]interface{})
	inputs["data"] = val
	inputs["value"] = 20
	inputs["minvalue"]= 2
	
	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}