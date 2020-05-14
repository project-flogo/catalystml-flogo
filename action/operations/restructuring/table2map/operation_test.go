package table2map


import (

	"testing"

	"github.com/project-flogo/core/support/log"
	
	"github.com/stretchr/testify/assert"
)

func Test1dArrAxis1(t *testing.T) {
	//params := Params{}
	var err error
	opt := &Operation{logger : log.RootLogger(), params: &Params{Axis : 1} }
	
	inputs := make(map[string]interface{})
	inputs["colKeys"] = []interface{}{"a",1, "c","b","k"}
	inputs["table"] = []interface{}{1,3,4,5,6}

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func Test2dArrAxis1(t *testing.T) {
	//params := Params{}
	var err error
	opt := &Operation{logger : log.RootLogger(), params: &Params{Axis : 1} }
	
	inputs := make(map[string]interface{})
	inputs["colKeys"] = []interface{}{"a",1, "c","b","k"}
	inputs["table"] = [][]interface{}{{1,4},{3,6},{4,7,8,9},{5},{6}}

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func Test2dArrAxis0(t *testing.T) {
	//params := Params{}
	var err error
	opt := &Operation{logger : log.RootLogger(), params: &Params{Axis : 0} }
	
	inputs := make(map[string]interface{})
	inputs["colKeys"] = []interface{}{"a"}
	inputs["table"] = [][]interface{}{{1,4},{3,6},{4,7,8,9},{5},{6}}

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
/*
func Test2dArrAxis0(t *testing.T) {
	//params := Params{}
	var err error
	opt := &Operation{logger : log.RootLogger(), params: &Params{Axis : 1} }
	
	inputs := make(map[string]interface{})
	inputs["colKeys"] = []interface{}{"a",1, "c","b","k"}
	inputs["table"] = [][]interface{}{{1,4},{3,6},{4,7,8,9},{5},{6}}

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}*/