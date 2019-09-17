package groupBy

import (
	"testing"

	"github.com/project-flogo/cml/action/support/test"
	"github.com/project-flogo/cml/operations/common"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
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
		       Wild          20.0
		>>> df.groupby(level=0).mean()
		        Max Speed
		Animal
		Falcon      370.0
		Parrot       25.0
		>>> df.groupby(level=1).mean()
		         Max Speed
		Type
		Captive      210.0
		Wild         185.0	*/
	inputs := make(map[string]interface{})
	dataFrame := make(common.DataFrame)
	inputs["data"] = dataFrame

	dataFrame["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	dataFrame["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	dataFrame["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}

	params := Params{
		Index:    []string{"Animal", "Type"},
		Target:   "Max Speed",
		Function: "mean",
		Level:    0,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}

func Test2(t *testing.T) {
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
		       Wild          20.0
		>>> df.groupby(level=0).mean()
		        Max Speed
		Animal
		Falcon      370.0
		Parrot       25.0
		>>> df.groupby(level=1).mean()
		         Max Speed
		Type
		Captive      210.0
		Wild         185.0	*/
	inputs := make(map[string]interface{})
	dataFrame := make(common.DataFrame)
	inputs["data"] = dataFrame

	dataFrame["Animal"] = []interface{}{
		"Falcon", "Falcon", "Parrot", "Parrot",
	}
	dataFrame["Type"] = []interface{}{
		"Captive", "Wild", "Captive", "Wild",
	}
	dataFrame["Max Speed"] = []interface{}{
		390.0, 350.0, 30.0, 20.0,
	}

	params := Params{
		Index:    []string{"Animal", "Type"},
		Target:   "Max Speed",
		Function: "mean",
		Level:    1,
	}

	optInitConext := test.NewOperationInitContext(params, nil)

	opt, err := New(optInitConext)
	assert.Nil(t, err)

	_, err = opt.Eval(inputs)
	assert.Nil(t, err)
}
