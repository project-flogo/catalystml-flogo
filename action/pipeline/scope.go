package pipeline

import (
	"strconv"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/path"
)

type scopeImpl struct {
	values map[string]interface{}
}

func NewPipelineScope(input map[string]interface{}) data.Scope {

	if input != nil {
		return &scopeImpl{values: input}
	}
	values := make(map[string]interface{})
	return &scopeImpl{values: values}
}

func (s *scopeImpl) GetValue(name string) (value interface{}, exists bool) {
	val, ok := s.values[name]

	if !ok {
		return nil, false
	}
	return val, true
}

//Check if the name resolves to existing values in scope
//If not then set a new value
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
