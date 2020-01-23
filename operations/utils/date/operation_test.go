package date

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	input := make(map[string]interface{})
	input["data"] = "2020-01-24T06:28:00.001Z"

	params := Params{
		Format: "2006-01-02T15:04:05.000Z",
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)

	output, err := opt.Eval(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)

	t.Log("Input of Operation Date : ", input["data"])
	t.Log("Output of Operation Date : ", output)

}

func Test_ANSIC(t *testing.T) {
	input := make(map[string]interface{})
	input["data"] = "Fri Jan 24 06:28:00 2020"

	params := Params{
		Format: "Mon Jan _2 15:04:05 2006",
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)

	output, err := opt.Eval(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)

	t.Log("Input of Operation Date : ", input["data"])
	t.Log("Output of Operation Date : ", output)

}

func Test_RFC822(t *testing.T) {
	input := make(map[string]interface{})
	input["data"] = "24 Jan 20 06:28 -0000"

	params := Params{
		Format: "02 Jan 06 15:04 -0500",
	}

	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)

	output, err := opt.Eval(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)

	t.Log("Input of Operation Date : ", input["data"])
	t.Log("Output of Operation Date : ", output)

}
