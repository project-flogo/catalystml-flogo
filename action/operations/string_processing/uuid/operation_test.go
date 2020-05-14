package uuid

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
