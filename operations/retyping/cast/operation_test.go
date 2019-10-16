package cast

import (
	"testing"

	"github.com/project-flogo/core/support/log"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	var err error
	inputs := make(map[string]interface{})
	inputs["data"] = []string{"1", "2.0", "2.5", "7.9"}
	inputs["toType"] = "float64"

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestSample2(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	var err error
	inputs := make(map[string]interface{})
	m := make(map[string]interface{})
	m["blah"] = []interface{}{3.4, 5.6, 7.8}
	inputs["data"] = m
	inputs["toType"] = "int32"

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestSampleBool(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	var err error
	inputs := make(map[string]interface{})
	m := make(map[string]interface{})
	m["blah"] = []interface{}{true, false, true}
	inputs["data"] = m
	inputs["toType"] = "int32"

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
func TestSampleBool2String(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	var err error
	inputs := make(map[string]interface{})
	m := make(map[string]interface{})
	m["blah"] = []interface{}{true, false, true}
	inputs["data"] = m
	inputs["toType"] = "string"

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}

func TestSamplebaseDataType(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}
	var err error
	inputs := make(map[string]interface{})
	m := "3.14159265359"
	inputs["data"] = m
	inputs["toType"] = "float32"

	_, err = opt.Eval(inputs)

	assert.Nil(t, err)

}
