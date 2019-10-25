package sort

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSortBySingleCol(t *testing.T) {
	/*
		>>> df = pd.DataFrame({
		...     'col1': ['A', 'A', 'B', np.nan, 'D', 'C'],
		...     'col2': [2, 1, 9, 8, 7, 4],
		...     'col3': [0, 1, 9, 4, 2, 3],
		... })
		>>> df
		   col1 col2 col3
		0   A    2    0
		1   A    1    1
		2   B    9    9
		3   NaN  8    4
		4   D    7    2
		5   C    4    3

		>>> df.sort_values(by=['col1'])
		    col1 col2 col3
		0   A    2    0
		1   A    1    1
		2   B    9    9
		5   C    4    3
		4   D    7    2
		3   NaN  8    4
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending: true,
		KeepRow:   true,
		By:        []interface{}{"col1"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestSortByMultipleCols(t *testing.T) {

	/*
	   >>> df.sort_values(by=['col1', 'col2'])
	       col1 col2 col3
	   1   A    1    1
	   0   A    2    0
	   2   B    9    9
	   5   C    4    3
	   4   D    7    2
	   3   NaN  8    4
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending: true,
		KeepRow:   true,
		By:        []interface{}{"col1", "col2"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestSortByMultipleColsByIndex(t *testing.T) {

	/*
	   >>> df.sort_values(by=[0, 1])
	       col1 col2 col3
	   1   A    1    1
	   0   A    2    0
	   2   B    9    9
	   5   C    4    3
	   4   D    7    2
	   3   NaN  8    4
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending: true,
		KeepRow:   true,
		By:        []interface{}{0, 1},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestSortDescending(t *testing.T) {
	/*
	   >>> df.sort_values(by='col1', ascending=False)
	       col1 col2 col3
	   4   D    7    2
	   5   C    4    3
	   2   B    9    9
	   0   A    2    0
	   1   A    1    1
	   3   NaN  8    4
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending: false,
		KeepRow:   true,
		By:        []interface{}{"col1"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestDescentingNilFirst(t *testing.T) {
	/*
	   >>> df.sort_values(by='col1', ascending=False, na_position='first')
	       col1 col2 col3
	   3   NaN  8    4
	   4   D    7    2
	   5   C    4    3
	   2   B    9    9
	   0   A    2    0
	   1   A    1    1
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending:   false,
		NilPosition: "first",
		KeepRow:     true,
		By:          []interface{}{"col1"},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestAscentingNilFirst(t *testing.T) {
	/*
	   >>> df.sort_values(by='col1', ascending=True, na_position='first')
	       col1 col2 col3
	   3   NaN  8    4
	   4   D    7    2
	   5   C    4    3
	   2   B    9    9
	   0   A    2    0
	   1   A    1    1
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending:   true,
		NilPosition: "first",
		KeepRow:     true,
		By:          []interface{}{"col1"},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func TestSortByRow(t *testing.T) {
	/*
		   >>> sorted=df.sort_values(by=1,axis=1)
			   col2 col3 col1
			0   2    0    A
			1   1    1    A
			2   9    9    B
			3  	8    4    NaN
			4   7    2    D
			5   4    3    C
	*/

	inputs := make(map[string]interface{})
	dataFrame := make(map[string]interface{})
	inputs["data"] = dataFrame

	dataFrame["order"] = []interface{}{
		"col1", "col2", "col3",
	}
	dataFrame["col1"] = []interface{}{
		"A", "A", "B", nil, "D", "C",
	}
	dataFrame["col2"] = []interface{}{
		2, 1, 9, 8, 7, 4,
	}
	dataFrame["col3"] = []interface{}{
		0, 1, 9, 4, 2, 3,
	}

	/*
		inputs := make(map[string]interface{})
		dataFrame := make([][]interface{})
		inputs["data"] = dataFrame

		dataFrame[0] = []interface{"col1", "col2", "col3"}

		dataFrame[1] = []interface{}{
			"A", "A", "B", "", "D", "C",
		}
		dataFrame[2] = []interface{}{
			"2, 1, 9, 8, 7, 4,
		}
		dataFrame[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	params := Params{
		Ascending: true,
		KeepRow:   true,
		By:        []interface{}{1},
		Axis:      1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
