package cmlmapper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
)

type CamMapper struct {
	defStructs []DerefernceStruct
}
type DerefernceStruct struct {
	Id    string
	Index string
}

func NewExpression(str string) []DerefernceStruct {
	var derefStructs []DerefernceStruct

	id := str[1:strings.Index(str, "[")]
	for key, val := range strings.Split(str, "[") {
		var derefStruct DerefernceStruct
		derefStruct.Id = id
		if key == 0 {

			continue
		}
		val = strings.TrimFunc(val, removeChars)

		derefStruct.Index = val

		derefStructs = append(derefStructs, derefStruct)
	}

	return derefStructs
}

func Resolve(deStructs []DerefernceStruct, scope data.Scope) (temp interface{}, err error) {

	for _, val := range deStructs {
		var temp2 []interface{}
		var temp3 map[string]interface{}

		if temp == nil {
			temp, _ = scope.GetValue(val.Id)
		}
		temp2, err = coerce.ToArray(temp)
		if err != nil {
			temp3, err = coerce.ToObject(temp)
			if err != nil {
				return nil, err
			}
			var ok bool

			temp, ok = temp3[val.Index]
			if !ok {
				return nil, err
			}

		} else {
			index, _ := strconv.Atoi(val.Index)
			temp = temp2[index]
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

				fmt.Println("Val inde...", val.Index, key)
				fmt.Println("Setting...", temp3[val.Index], value)
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
				fmt.Println("Setting...", temp2[index], value)
			}
		}

	}
	//fmt.Println("Scope:", temp, value)
	//return temp, nil

}
