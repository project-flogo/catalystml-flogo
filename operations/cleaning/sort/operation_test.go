package sort

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSortTableBySingleCol(t *testing.T) {
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
		Ascending: true,
		By:        []interface{}{"col1"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortTableByMultipleCols(t *testing.T) {

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
		Ascending: true,
		By:        []interface{}{"col1", "col2"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortTableByMultipleColsByIndex(t *testing.T) {

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
		Ascending: true,
		By:        []interface{}{0, 1},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestTableSortDescending(t *testing.T) {
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
		Ascending: false,
		By:        []interface{}{"col1"},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortTableDescentingNilFirst(t *testing.T) {
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
		Ascending:   false,
		NilPosition: "first",
		By:          []interface{}{"col1"},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortTableAscentingNilFirst(t *testing.T) {
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
		Ascending:   true,
		NilPosition: "first",
		By:          []interface{}{"col1"},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortTableByRow(t *testing.T) {
	/*
		   >>> sorted=df.sort_values(by=1,axis=1)
			    col3 col2 col1
			0   0    2    A
			1   1    1    A
			2   9    9    B
			3  	4    8    NaN
			4   2    7    D
			5   3    4    C
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
		Ascending: true,
		By:        []interface{}{1},
		Axis:      1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixBySingleCol(t *testing.T) {
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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{0},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixByMultipleCols(t *testing.T) {

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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{0, 1},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixByMultipleColsByIndex(t *testing.T) {

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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{0, 1},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestMatrixSortDescending(t *testing.T) {
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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: false,
		By:        []interface{}{0},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixDescentingNilFirst(t *testing.T) {
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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending:   false,
		NilPosition: "first",
		By:          []interface{}{0},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixAscentingNilFirst(t *testing.T) {
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
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending:   true,
		NilPosition: "first",
		By:          []interface{}{0},
		Axis:        0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortMatrixByRow(t *testing.T) {
	/*
		   >>> sorted=df.sort_values(by=1,axis=1)
			    col3 col2 col1
			0   0    2    A
			1   1    1    A
			2   9    9    B
			3  	4    8    NaN
			4   2    7    D
			5   3    4    C
	*/

	inputs := make(map[string]interface{})
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{1},
		Axis:      1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortArrayByRow(t *testing.T) {
	/*
		   >>> sorted=df.sort_values(by=1,axis=1)
			    col3 col2 col1
			0   0    2    A
			1   1    1    A
			2   9    9    B
			3  	4    8    NaN
			4   2    7    D
			5   3    4    C
	*/

	inputs := make(map[string]interface{})
	matrix := make([][]interface{}, 6)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		"A", 2, 0,
	}
	matrix[1] = []interface{}{
		"A", 1, 1,
	}
	matrix[2] = []interface{}{
		"B", 9, 9,
	}
	matrix[3] = []interface{}{
		nil, 8, 4,
	}
	matrix[4] = []interface{}{
		"D", 7, 2,
	}
	matrix[5] = []interface{}{
		"C", 4, 0,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{0},
		Axis:      0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}

func TestSortArrayByColumn(t *testing.T) {
	/*
		   >>> sorted=df.sort_values(by=1,axis=1)
			    col3 col2 col1
			0   0    2    A
			1   1    1    A
			2   9    9    B
			3  	4    8    NaN
			4   2    7    D
			5   3    4    C
	*/

	inputs := make(map[string]interface{})
	matrix := make([][]interface{}, 1)
	inputs["data"] = matrix

	/*
		matrix[0] = []interface{}{"col1", "col2", "col3"}

		matrix[1] = []interface{}{
			"A", "A", "B", nil, "D", "C",
		}
		matrix[2] = []interface{}{
			2, 1, 9, 8, 7, 4,
		}
		matrix[3] = []interface{}{
			0, 1, 9, 4, 2, 3,
		}
	*/

	matrix[0] = []interface{}{
		0, 0, 1, nil, 3, 2,
	}

	params := Params{
		Ascending: true,
		By:        []interface{}{0},
		Axis:      1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", matrix)
	t.Log("Output of Operation Sort : ", out)

}
