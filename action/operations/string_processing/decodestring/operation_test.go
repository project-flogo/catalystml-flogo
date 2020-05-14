package decodestring

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	input := make(map[string]interface{})
	input["str"] = "SGkgSG8gb2ZmIHRvIHdvcmsgd2UgZ28hIQ=="

	output, err := opt.Eval(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)

	t.Log("Input of Operation Decodestring : ", input["str"])
	t.Log("Output of Operation Decodestring : ", output)

}
