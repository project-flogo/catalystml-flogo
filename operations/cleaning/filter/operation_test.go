package filter

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	/*
		{
		    	"col1": ["A", "A", "B", nil, "D", "C"],
			"col2": [2, 1, 9, 8, 7, 4],
			"col3": [0, 1, 9, 4, 2, 3],
		}

			*
			col1 col2 col3
		0   	A    2    0
		1   	A    1    1
		2  	B    9    9
		3   	nil  8    4
		4   	D    7    2
		5   	C    4    3

		filter with : axis=0, col="col1", value=nil, filterType="Remove"

			*
			col1 col2 col3
		0   	A    2    0
		1   	A    1    1
		2  	B    9    9
		3   	D    7    2
		4   	C    4    3

		{
		    	"col1": ["A", "A", "B", "D", "C"],
			"col2": [2, 1, 9, 7, 4],
			"col3": [0, 1, 9, 2, 3],
		}

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
	inputs["value"] = nil
	inputs["filterType"] = "Remove"

	params := Params{
		Col:  "col1",
		Axis: 0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", table)
	t.Log("Output of Operation Filter : ", out)

}

func Test02(t *testing.T) {
	/*
		{
		    	"col1": ["A", "A", "B", nil, "D", "C"],
			"col2": [2, 1, 9, 8, 7, 4],
			"col3": [0, 1, 9, 4, 2, 3],
		}

			*
			col1 col2 col3
		0   	A    2    0
		1   	A    1    1
		2  	B    9    9
		3   	nil  8    4
		4   	D    7    2
		5   	C    4    3

		filter with : axis=0, col="col1", value="A", filterType="Keep"

			*
			col1 col2 col3
		0   	A    2    0
		1   	A    1    1

		{
		    	"col1": ["A", "A"],
			"col2": [2, 1],
			"col3": [0, 1],
		}

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
	inputs["value"] = "A"
	inputs["filterType"] = "Keep"

	params := Params{
		Col:  "col1",
		Axis: 0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", table)
	t.Log("Output of Operation Filter : ", out)

}

func Test03(t *testing.T) {
	/*
		[[3 2 0] [7 1 1] [2 9 9] [<nil> 8 4] [2 7 2] [0 4 0]]

			*
			0	 1	  2
		0   	3    2    0
		1   	7    1    1
		2  	2    9    9
		3   	nil  8    4
		4   	2    7    2
		5   	0    4    3

		filter with : axis=1, col="0", value=nil, filterType="Remove"

			*
			0 	 1 	  2
		0   	3    2    0
		1   	7    1    1
		2  	2    9    9
		3   	2    7    2
		4   	0    4    3

		[[3 2 0] [7 1 1] [2 9 9] [2 7 2] [0 4 0]]

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

	inputs["value"] = nil
	inputs["filterType"] = "Remove"

	params := Params{
		Col:  "0",
		Axis: 0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", matrix)
	t.Log("Output of Operation Filter : ", out)

}

func Test04(t *testing.T) {
	/*
		[[3 2 0] [7 1 1] [2 9 9] [<nil> 8 4] [2 7 2] [0 4 0]]

			*
			0	 1	  2
		0   	3    2    0
		1   	7    1    1
		2  	2    9    9
		3   	nil  8    4
		4   	2    7    2
		5   	0    4    3

		filter with : axis=0, col="0", value="2", filterType="Keep"

			*
			0 	 1 	  2
		0  	2    9    9
		1   	2    7    2

		[[2 9 9] [2 7 2]]

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

	inputs["value"] = 2
	inputs["filterType"] = "Keep"

	params := Params{
		Col:  "0",
		Axis: 0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", matrix)
	t.Log("Output of Operation Filter : ", out)

}

func Test05(t *testing.T) {
	/*
		[[3 2 0] [7 1 1] [2 9 9] [<nil> 8 4] [2 7 2] [0 4 0]]

				0	 1	  2
			0   	3    2    0
			1   	7    1    1
			2  	2    9    9
		*	3   	nil  8    4
			4   	2    7    2
			5   	0    4    3

			filter with : axis=1, col="3", value=nil, filterType="Remove"

				0 	 1
			0   	2    0
			1   	1    1
			2  	9    9
		*	3   	8    4
			4   	7    2
			5   	4    3

		[[0 2] [1 1] [9 9] [4 8] [2 7] [0 4]]

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

	inputs["value"] = nil
	inputs["filterType"] = "Remove"

	params := Params{
		Col:  "3",
		Axis: 1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", matrix)
	t.Log("Output of Operation Filter : ", out)

}

func Test06(t *testing.T) {
	/*
		[[3 2 0] [7 1 1] [2 9 9] [<nil> 8 4] [2 7 2] [0 4 0]]

				0	 1	  2
			0   	3    2    0
			1   	7    1    1
		*	2  	2    9    9
			3   	nil  8    4
			4   	2    7    2
			5   	0    4    3

			filter with : axis=1, col="2", value=9, filterType="Keep"

				1	 2
			0   	2    0
			1   	1    1
		*	2  	9    9
			3   	8    4
			4   	7    2
			5   	4    3

		[[2 0] [1 1] [9 9] [8 4] [7 2] [4 0]]

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

	inputs["value"] = 9
	inputs["filterType"] = "Keep"

	params := Params{
		Col:  "2",
		Axis: 1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Filter : ", matrix)
	t.Log("Output of Operation Filter : ", out)

}
