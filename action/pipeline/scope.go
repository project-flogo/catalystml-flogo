package pipeline

import (
	"errors"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/path"
)

type scopeImpl struct {
	values map[string]interface{}
}

// NewPipelineScope gets the scope from the input and label map.
func NewPipelineScope(input map[string]interface{}, labels map[string]interface{}) (data.Scope, error) {

	if input != nil {
		val, err := preProcessInputs(input, labels)
		if err != nil {
			return nil, err
		}
		return &scopeImpl{values: val}, nil
	}
	values := make(map[string]interface{})
	return &scopeImpl{values: values}, nil
}

func (s *scopeImpl) GetValue(name string) (value interface{}, exists bool) {
	val, ok := s.values[name]

	if !ok {
		return nil, false
	}
	return val, true
}

// SetValue sets a key, value pair in the scope/
func (s *scopeImpl) SetValue(name string, value interface{}) error {

	if strings.Contains(name, "[") {

		path.SetValue(s.values, getPath(name), value)

	} else {
		s.values[name] = value
	}

	return nil
}

func getPath(name string) string {
	var result string
	for _, val := range strings.Split(name, "[") {
		temp := strings.TrimFunc(val, func(r rune) bool {
			if r == '\'' || r == ']' {
				return true
			}
			return false
		})
		if _, err := strconv.Atoi(temp); err == nil {
			result = result + "[" + temp + "]"
		} else {
			result = result + "." + temp
		}

	}
	return result
}

// PreProcessInput maps the labels to the data from input.
// Eg. For input  {"a": "2", "b": 3} and label being ["firstInput", "secondInput"]. firstInput maps to "2"
// Also is input is [3, 4] and label being ["fInput", "sOutput"] ; fInput maps to 3
// For more information : //TODO: Add the link to the label issue.
func preProcessInputs(inputs map[string]interface{}, labels map[string]interface{}) (map[string]interface{}, error) {
	inputMap := make(map[string]interface{})

	if val, ok := inputs["input"]; ok && len(labels) != 0 {
		vArr, _ := coerce.ToArray(val)

		for key, in := range vArr {

			switch t := labels[strconv.Itoa(key)].(type) {

			case string:
				inputMap[t] = in
			case []interface{}:
				if len(t) > len(vArr) {
					return nil, errors.New("Mismatch in Data and Labels ")
				}
				for i := 0; i < len(t); i++ {
					inputMap[t[i].(string)] = vArr[i]
				}
			}
		}
		for key, val := range inputs {
			inputMap[key] = val
		}

		return inputMap, nil
	}
	if _, ok := inputs["0"]; ok || len(labels) != 0 {

		for key, val := range labels {

			switch t := val.(type) {
			case string:
				inputMap[t] = inputs[key]
			case []interface{}:
				vArr, _ := coerce.ToArray(inputs[key])
				if len(t) > len(vArr) {
					return nil, errors.New("Mismatch in Data and Labels ")
				}
				for i := 0; i < len(t); i++ {
					inputMap[t[i].(string)] = vArr[i]
				}
			}

		}
		for key, val := range inputs {
			inputMap[key] = val
		}

		return inputMap, nil
	}

	return inputs, nil

}
