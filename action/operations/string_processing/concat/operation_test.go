package concat

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["s0"] = "Hi Ho Hi Ho."
	inputs["s1"] = " Off to work we go!!"
	inputs["slist"] = []interface{}{"one ", "two ", "three ", "four"}

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
