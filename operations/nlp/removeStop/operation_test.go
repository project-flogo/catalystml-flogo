package removeStop

import (
	"testing"

	"github.com/project-flogo/fps/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["str"] = "<p>The hotspot is here</p>"

	p := Params{Lang: "en"}

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}
