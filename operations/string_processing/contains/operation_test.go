package contains

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["s0"] = "Hi Ho Hi Ho."
	inputs["s1"] = " Ho"

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestFalse(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["s0"] = "Hi Ho Hi Ho."
	inputs["s1"] = " Eep"

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
