package cmlmapper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
)

// CML Mapper is a mapper private to CML.
// It gets the required value from the underlying data
// in the CML pipeline.

//Eg. If  `math1` is defined as :
// [ 1  2    3
//   4 "Abc" 6 ]
// Using this mapper we can directly do $math1[0][2] within the CML Specification
// To get the desired value.
type CmlMapper struct {
	defStructs []DerefernceStruct
}

// A struct which helps to navigate the data
// to find the value.
type DerefernceStruct struct {
	Id    string //The id of data i.e. label of the data.
	Index string // The identifier of the value in the data. Eg math[0] or math['sample'] . 0 and 'sample' will be the Index.
}

// NewExpression returns the deference struct (which will help to navigate the data to find the value)
// from the given string.
func NewExpression(str string) []DerefernceStruct {
	var derefStructs []DerefernceStruct
	var id string

	//Get label of the data.
	if str[0] == '$' {
		id = str[1:strings.Index(str, "[")]
	} else {
		id = str[0:strings.Index(str, "[")]
	}

	// Split the string on "[". The first value of array will be label so
	// ignore that.
	// For subsequent value construct the deference struct
	for key, val := range strings.Split(str, "[") {
		var derefStruct DerefernceStruct
		derefStruct.Id = id
		if key == 0 {

			continue
		}
		//Remove "[" or "]" or "\" from the string.
		val = strings.TrimFunc(val, removeChars)

		derefStruct.Index = val

		derefStructs = append(derefStructs, derefStruct)
	}

	return derefStructs
}

// Resolve resolves the value from the data using dereference struct and scope.
// scope is the collection of the data.
func Resolve(deStructs []DerefernceStruct, scope data.Scope) (temp interface{}, err error) {

	if scope == nil {
		return nil, fmt.Errorf("Scope cannot be nil")
	}
	//Iterate over dereference struct.
	for _, val := range deStructs {
		//Declare a temporary map and array

		var tempArray []interface{}
		var tempMap map[string]interface{}

		//Get the data from the scope.
		if temp == nil {
			var ok bool
			temp, ok = scope.GetValue(val.Id)
			if !ok {
				return nil, fmt.Errorf("Unable to find value related to %s", val.Id)
			}
		}
		// Try to convert the temp calue to Array.
		tempArray, err = coerce.ToArray(temp)
		if err != nil {
			// If error convert it to map.
			tempMap, err = coerce.ToObject(temp)
			if err != nil {
				//If error return
				return nil, err
			}
			var ok bool
			// Get the value from the map using the
			// dereference struct's index
			temp, ok = tempMap[val.Index]
			if !ok {
				//If not found; return error
				return nil, err
			}

		} else {
			// Convert the index of dereference struct to int
			index, err := strconv.Atoi(val.Index)
			if err != nil {
				//If error return
				return nil, err
			}
			// Get the value from the array.
			temp = tempArray[index]
		}

	}

	return temp, nil
}

func removeChars(r rune) bool {
	if r == ']' || r == '\'' {
		return true
	}
	return false
}

func Apply(deStructs []DerefernceStruct, scope data.Scope, value interface{}) {

	var temp interface{}
	var err error
	size := len(deStructs)

	for key, val := range deStructs {

		var temp2 []interface{}
		var temp3 map[string]interface{}

		if temp == nil {
			temp, _ = scope.GetValue(val.Id)
		}
		temp2, err = coerce.ToArray(temp)
		if err != nil {
			temp3, err = coerce.ToObject(temp)
			if err != nil {
				//return nil, err
			}
			var ok bool

			temp, ok = temp3[val.Index]
			if key == size-1 {

				temp3[val.Index] = value
			}
			if !ok {
				//return nil, err
			}

		} else {
			index, _ := strconv.Atoi(val.Index)
			temp = temp2[index]
			if key == size-1 {
				temp2[index] = val

			}
		}

	}
	//fmt.Println("Scope:", temp, value)
	//return temp, nil

}
