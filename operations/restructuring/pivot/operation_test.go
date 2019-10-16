package pivot

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	/*
	   >>> table = pd.pivot_table(df, values='D', index=['A', 'B'],
	   ...                     columns=['C'], aggfunc=np.sum)
	   >>> table
	   C        large  small
	   A   B
	   bar one    4.0    5.0
	       two    7.0    6.0
	   foo one    4.0    1.0
	       two    NaN    6.0
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["data"] = dataFrame

	dataFrame["A"] = []interface{}{
		"foo", "foo", "foo", "foo", "foo", "bar", "bar", "bar", "bar",
	}
	dataFrame["B"] = []interface{}{
		"one", "one", "one", "two", "two", "one", "one", "two", "two",
	}
	dataFrame["C"] = []interface{}{
		"small", "large", "large", "small", "small", "large", "small", "small", "large",
	}
	dataFrame["D"] = []interface{}{
		1, 2, 2, 3, 3, 4, 5, 6, 7,
	}
	dataFrame["E"] = []interface{}{
		2, 4, 5, 5, 6, 6, 8, 9, 9,
	}

	aggregate := make(map[string][]string)
	aggregate["D"] = []string{"sum"}
	params := Params{
		Index:     []string{"A", "B"},
		Columns:   []string{"C"},
		Aggregate: aggregate,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test2(t *testing.T) {
	/*
	   >>> table = pd.pivot_table(df, values=['D', 'E'], index=['A', 'C'],
	   ...                     aggfunc={'D': np.mean,
	   ...                              'E': np.mean})
	   >>> table
	                   D         E
	   A   C
	   bar large  5.500000  7.500000
	       small  5.500000  8.500000
	   foo large  2.000000  4.500000
	       small  2.333333  4.333333
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["data"] = dataFrame

	dataFrame["A"] = []interface{}{
		"foo", "foo", "foo", "foo", "foo", "bar", "bar", "bar", "bar",
	}
	dataFrame["B"] = []interface{}{
		"one", "one", "one", "two", "two", "one", "one", "two", "two",
	}
	dataFrame["C"] = []interface{}{
		"small", "large", "large", "small", "small", "large", "small", "small", "large",
	}
	dataFrame["D"] = []interface{}{
		1, 2, 2, 3, 3, 4, 5, 6, 7,
	}
	dataFrame["E"] = []interface{}{
		2, 4, 5, 5, 6, 6, 8, 9, 9,
	}

	aggregate := make(map[string][]string)
	aggregate["D"] = []string{"mean"}
	aggregate["E"] = []string{"mean"}

	params := Params{
		Index:     []string{"A", "C"},
		Columns:   []string{},
		Aggregate: aggregate,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test3(t *testing.T) {
	/*
	   >>> table = pd.pivot_table(df, values=['D', 'E'], index=['A', 'C'],
	   ...                     aggfunc={'D': np.mean,
	   ...                              'E': [min, max, np.mean]})
	   >>> table
	                   D    E
	               mean  max      mean  min
	   A   C
	   bar large  5.500000  9.0  7.500000  6.0
	       small  5.500000  9.0  8.500000  8.0
	   foo large  2.000000  5.0  4.500000  4.0
	       small  2.333333  6.0  4.333333  2.0
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["data"] = dataFrame

	dataFrame["A"] = []interface{}{
		"foo", "foo", "foo", "foo", "foo", "bar", "bar", "bar", "bar",
	}
	dataFrame["B"] = []interface{}{
		"one", "one", "one", "two", "two", "one", "one", "two", "two",
	}
	dataFrame["C"] = []interface{}{
		"small", "large", "large", "small", "small", "large", "small", "small", "large",
	}
	dataFrame["D"] = []interface{}{
		1, 2, 2, 3, 3, 4, 5, 6, 7,
	}
	dataFrame["E"] = []interface{}{
		2.0, 4.0, 5.0, 5.0, 6.0, 6.0, 8.0, 9.0, 9.0,
	}

	aggregate := make(map[string][]string)
	aggregate["D"] = []string{"mean"}
	aggregate["E"] = []string{"min", "max", "mean"}

	params := Params{
		Index:     []string{"A", "C"},
		Columns:   []string{},
		Aggregate: aggregate,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test4(t *testing.T) {
	/*
	 */
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["data"] = dataFrame

	dataFrame["A"] = []interface{}{
		"foo", "foo", "foo", "foo", "foo", "bar", "bar", "bar", "bar",
	}
	dataFrame["B"] = []interface{}{
		"one", "one", "one", "two", "two", "one", "one", "two", "two",
	}
	dataFrame["C"] = []interface{}{
		"small", "large", "large", "small", "small", "large", "small", "small", "large",
	}
	dataFrame["D"] = []interface{}{
		1, 2, 2, 3, 3, 4, 5, 6, 7,
	}
	dataFrame["E"] = []interface{}{
		2, 4, 5, 5, 6, 6, 8, 9, 9,
	}

	aggregate := make(map[string][]string)
	aggregate["D"] = []string{"sum", "count"}
	aggregate["E"] = []string{"mean"}

	params := Params{
		Index:     []string{"A", "C"},
		Columns:   []string{},
		Aggregate: aggregate,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
