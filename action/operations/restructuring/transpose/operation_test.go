package transpose

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	table["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	table["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	table["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Transpose : ", table)
	t.Log("Output of Operation Transpose : ", out)
}

func TestMatrix(t *testing.T) {

	matrix := [][]interface{}{
		{"A", "A", "B", nil, "D", "C"},
		{2, 1, 9, 8, 7, 4},
		{0, 1, 9, 4, 2, 0},
	}

	inputs := make(map[string]interface{})
	inputs["data"] = matrix

	optInitConext := test.NewOperationInitContext(nil, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Transpose : ", matrix)
	t.Log("Output of Operation Transpose : ", out)
}
