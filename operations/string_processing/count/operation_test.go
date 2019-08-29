package count

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["s0"] = "moo moo moo"
	inputs["s1"] = "moo"

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	assert.Equal(t, output, 3, "two should be the same")
	assert.Nil(t, err)

}
