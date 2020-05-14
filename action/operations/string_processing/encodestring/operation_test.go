package encodestring

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	input := make(map[string]interface{})
	input["str"] = "Hi Ho off to work we go!!"

	output, err := opt.Eval(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)

	t.Log("Input of Operation Encodestring : ", input["str"])
	t.Log("Output of Operation Encodestring : ", output)

}
