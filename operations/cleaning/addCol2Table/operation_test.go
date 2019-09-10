package addCol2Table

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	// params := Params{}
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["matrix"] = [][]float32{[]float32{1.4}, []float32{1.1}}
	inputs["col"] = []float32{1.3, 1.6}

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
