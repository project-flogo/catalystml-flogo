package join

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	/*
	   In [79]: left = pd.DataFrame({'A': ['A0', 'A1', 'A2'],
	      ....:                      'B': ['B0', 'B1', 'B2']},
	      ....:                     index=['K0', 'K1', 'K2'])
	      ....:

	   In [80]: right = pd.DataFrame({'C': ['C0', 'C2', 'C3'],
	      ....:                       'D': ['D0', 'D2', 'D3']},
	      ....:                      index=['K0', 'K2', 'K3'])
	      ....:

	   In [81]: result = left.join(right)
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["left"] = dataFrame

	dataFrame["A"] = []interface{}{
		"A0", "A1", "A2",
	}
	dataFrame["B"] = []interface{}{
		"B0", "B1", "B2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K1", "K2",
	}

	dataFrame = make(map[string][]interface{})
	inputs["right"] = dataFrame

	dataFrame["C"] = []interface{}{
		"C0", "C1", "C2",
	}
	dataFrame["D"] = []interface{}{
		"D0", "D1", "D2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K2", "K3",
	}

	inputs["leftindex"] = []string{"I"}
	inputs["rightindex"] = []string{"I"}

	params := Params{
		On:  []string{"I"},
		How: "left",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test2(t *testing.T) {
	/*
	   In [79]: left = pd.DataFrame({'A': ['A0', 'A1', 'A2'],
	      ....:                      'B': ['B0', 'B1', 'B2']},
	      ....:                     index=['K0', 'K1', 'K2'])
	      ....:

	   In [80]: right = pd.DataFrame({'C': ['C0', 'C2', 'C3'],
	      ....:                       'D': ['D0', 'D2', 'D3']},
	      ....:                      index=['K0', 'K2', 'K3'])
	      ....:

	   In [81]: result = left.join(right)
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["left"] = dataFrame

	dataFrame["A"] = []interface{}{
		"A0", "A1", "A2",
	}
	dataFrame["B"] = []interface{}{
		"B0", "B1", "B2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K1", "K2",
	}

	dataFrame = make(map[string][]interface{})
	inputs["right"] = dataFrame

	dataFrame["C"] = []interface{}{
		"C0", "C1", "C2",
	}
	dataFrame["D"] = []interface{}{
		"D0", "D1", "D2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K2", "K3",
	}

	inputs["leftindex"] = []string{"I"}
	inputs["rightindex"] = []string{"I"}

	params := Params{
		On:  []string{"I"},
		How: "right",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test3(t *testing.T) {
	/*
	   In [79]: left = pd.DataFrame({'A': ['A0', 'A1', 'A2'],
	      ....:                      'B': ['B0', 'B1', 'B2']},
	      ....:                     index=['K0', 'K1', 'K2'])
	      ....:

	   In [80]: right = pd.DataFrame({'C': ['C0', 'C2', 'C3'],
	      ....:                       'D': ['D0', 'D2', 'D3']},
	      ....:                      index=['K0', 'K2', 'K3'])
	      ....:

	   In [81]: result = left.join(right)
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["left"] = dataFrame

	dataFrame["A"] = []interface{}{
		"A0", "A1", "A2",
	}
	dataFrame["B"] = []interface{}{
		"B0", "B1", "B2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K1", "K2",
	}

	dataFrame = make(map[string][]interface{})
	inputs["right"] = dataFrame

	dataFrame["C"] = []interface{}{
		"C0", "C1", "C2",
	}
	dataFrame["D"] = []interface{}{
		"D0", "D1", "D2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K2", "K3",
	}

	inputs["leftindex"] = []string{"I"}
	inputs["rightindex"] = []string{"I"}

	params := Params{
		On:  []string{"I"},
		How: "outer",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test4(t *testing.T) {
	/*
	   In [79]: left = pd.DataFrame({'A': ['A0', 'A1', 'A2'],
	      ....:                      'B': ['B0', 'B1', 'B2']},
	      ....:                     index=['K0', 'K1', 'K2'])
	      ....:

	   In [80]: right = pd.DataFrame({'C': ['C0', 'C2', 'C3'],
	      ....:                       'D': ['D0', 'D2', 'D3']},
	      ....:                      index=['K0', 'K2', 'K3'])
	      ....:

	   In [81]: result = left.join(right)
	*/
	inputs := make(map[string]interface{})
	dataFrame := make(map[string][]interface{})
	inputs["left"] = dataFrame

	dataFrame["A"] = []interface{}{
		"A0", "A1", "A2",
	}
	dataFrame["B"] = []interface{}{
		"B0", "B1", "B2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K1", "K2",
	}

	dataFrame = make(map[string][]interface{})
	inputs["right"] = dataFrame

	dataFrame["C"] = []interface{}{
		"C0", "C1", "C2",
	}
	dataFrame["D"] = []interface{}{
		"D0", "D1", "D2",
	}
	dataFrame["I"] = []interface{}{
		"K0", "K2", "K3",
	}

	inputs["leftindex"] = []string{"I"}
	inputs["rightindex"] = []string{"I"}

	params := Params{
		On:  []string{"I"},
		How: "inner",
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
