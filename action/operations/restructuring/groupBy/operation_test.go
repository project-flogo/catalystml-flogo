package groupBy

import (
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	/*
		>>> arrays = [['Falcon', 'Falcon', 'Parrot', 'Parrot'],
		...           ['Captive', 'Wild', 'Captive', 'Wild']]
		>>> index = pd.MultiIndex.from_arrays(arrays, names=('Animal', 'Type'))
		>>> df = pd.DataFrame({'Max Speed': [390., 350., 30., 20.]},
		...                   index=index)
		>>> df
		                Max Speed
		Animal Type
		Falcon Captive      390.0
		       Wild         350.0
		Parrot Captive       30.0
		       Wild          20.0*/
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	table["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	table["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}
	/*
		params := Params{
			Index:    []string{"Animal", "Type"},
			Target:   "Max Speed",
			Function: "mean",
			Level:    0,
		}
	*/
	aggregate := make(map[string][]string)
	aggregate["Max Speed"] = []string{"mean"}
	params := Params{
		Index:     []string{"Animal", "Type"},
		Aggregate: aggregate,
		Level:     -1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func Test1(t *testing.T) {
	/*
		>>> arrays = [['Falcon', 'Falcon', 'Parrot', 'Parrot'],
		...           ['Captive', 'Wild', 'Captive', 'Wild']]
		>>> index = pd.MultiIndex.from_arrays(arrays, names=('Animal', 'Type'))
		>>> df = pd.DataFrame({'Max Speed': [390., 350., 30., 20.]},
		...                   index=index)
		>>> df.groupby(level=0).mean()
		        Max Speed
		Animal
		Falcon      370.0
		Parrot       25.0*/
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	table["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	table["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}
	/*
		params := Params{
			Index:    []string{"Animal", "Type"},
			Target:   "Max Speed",
			Function: "mean",
			Level:    0,
		}
	*/
	aggregate := make(map[string][]string)
	aggregate["Max Speed"] = []string{"mean"}
	params := Params{
		Index:     []string{"Animal", "Type"},
		Aggregate: aggregate,
		Level:     0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func Test2(t *testing.T) {
	/*
		>>> arrays = [['Falcon', 'Falcon', 'Parrot', 'Parrot'],
		...           ['Captive', 'Wild', 'Captive', 'Wild']]
		>>> index = pd.MultiIndex.from_arrays(arrays, names=('Animal', 'Type'))
		>>> df = pd.DataFrame({'Max Speed': [390., 350., 30., 20.]},
		...                   index=index)
		>>> df.groupby(level=1).mean()
		         Max Speed
		Type
		Captive      210.0
		Wild         185.0	*/
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	table["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	table["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}
	/*
		params := Params{
			Index:    []string{"Animal", "Type"},
			Target:   "Max Speed",
			Function: "mean",
			Level:    1,
		}
	*/
	aggregate := make(map[string][]string)
	aggregate["Max Speed"] = []string{"mean"}
	params := Params{
		Index:     []string{"Animal", "Type"},
		Aggregate: aggregate,
		Level:     1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}

func Test3(t *testing.T) {
	inputs := make(map[string]interface{})
	table := make(map[string]interface{})
	inputs["data"] = table

	table["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	table["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	table["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}
	/*
		params := Params{
			Index:    []string{"Animal", "Type"},
			Target:   "Max Speed",
			Function: "mean",
			Level:    1,
		}
	*/
	aggregate := make(map[string][]string)
	aggregate["Max Speed"] = []string{"mean", "min"}
	params := Params{
		Index:     []string{"Animal", "Type"},
		Aggregate: aggregate,
		Level:     0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	out, err := opt.Eval(inputs)
	assert.Nil(t, err)

	t.Log("Input of Operation Sort : ", table)
	t.Log("Output of Operation Sort : ", out)
}
