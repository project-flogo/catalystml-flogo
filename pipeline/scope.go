package pipeline

import (
	"errors"
)

type scopeImpl struct {
	values map[string]interface{}
}

func NewPipelineScope(input map[string]interface{}) *scopeImpl {

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

func (s *scopeImpl) SetValue(name string, value interface{}) error {

	_, ok := s.values[name]
	if ok {
		return errors.New("Value exist")
	}
	s.values[name] = value

	return nil
}
