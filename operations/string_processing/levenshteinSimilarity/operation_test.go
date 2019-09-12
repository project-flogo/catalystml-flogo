package levenshteinSimilarity

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	// params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["s0"] = "Password"
	inputs["s1"] = "p@ssW0rd"

	output, err := opt.Eval(inputs)
	assert.Equal(t, output.(float64), float64(0.5))
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
