package split

import (
	"testing"

	"github.com/project-flogo/cml/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSingleSplit(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["text"] = "Hello world"
	inputs["separator"] = " "

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []string{"Hello", "world"}, result)
}

func TestMultisplit(t *testing.T) {
	inputs := make(map[string]interface{})

	inputs["text"] = "Hello world"
	inputs["separator"] = "l"

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	result, err := opt.Eval(inputs)
	assert.Nil(t, err)

	assert.Equal(t, []string{"He", "", "o wor", "d"}, result)
}
