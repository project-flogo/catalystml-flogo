package stem

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["str"] = "Running"

	// p := Params{Algo: "Porter"}
	p := Params{Algo: "Snowball"}
	// p := Params{}

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Equal(t, output, "run", "two should be the same")
	assert.Nil(t, err)

}
