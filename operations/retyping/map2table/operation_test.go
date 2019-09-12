package map2table
import (
	"testing"
	"github.com/project-flogo/core/support/log"
	
	"github.com/stretchr/testify/assert"
)
func TestSample0(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger(), params :&Params{Axis:0}}
	var err error
	inputs := make(map[string]interface{})
	inputs["colOrder"] = []string{"a","b","c"}
	inputs["map"] = map[string]interface{}{"a":[]interface{}{1,2,4}, "b":[]interface{}{3,5,6}, "c":[]interface{}{7,8,9}}	

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func TestSample1(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger(), params :&Params{Axis:1}}
	var err error
	inputs := make(map[string]interface{})
	inputs["colOrder"] = []string{"a","b","c"}
	inputs["map"] = map[string]interface{}{"a":[]interface{}{1,2,4}, "b":[]interface{}{3,5,6}, "c":[]interface{}{7,8,9}}	

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}