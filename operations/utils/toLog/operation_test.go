package toLog

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	/*
			>>> {
		    			"col1": ["A", "A", "B", nil, "D", "C"],
					"col2": [2, 1, 9, 8, 7, 4],
					"col3": [0, 1, 9, 4, 2, 3],
				}
			>>>
	*/

	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table
	table["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	table["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	table["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	params := Params{
		ToFilePath:      "/Users/steven/Downloads/log01.log",
		ClearWhileStart: true,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation ToLog : ", table)
	t.Log("Output of Operation ToLog : ", out)

}

func Test02(t *testing.T) {
	/*
			>>> [
		    			[3, 7, 2, nil, 2, 0],
					[2, 1, 9, 8, 7, 4],
					[0, 1, 9, 4, 2, 3],
				]
			>>>
	*/

	inputs := make(map[string]interface{})
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	matrix[0] = []interface{}{
		3, 2, 0,
	}
	matrix[1] = []interface{}{
		7, 1, 1,
	}
	matrix[2] = []interface{}{
		2, 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		2, 7, 2,
	}
	matrix[5] = []interface{}{
		0, 4, 0,
	}

	params := Params{
		ToFilePath:      "/Users/steven/Downloads/log01.log",
		ClearWhileStart: false,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation ToLog : ", matrix)
	t.Log("Output of Operation ToLog : ", out)

}
