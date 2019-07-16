package pipeline

import (
	"errors"
)

type ScopeId int

const (
	ScopeDefault ScopeId = iota
	ScopePipeline
	ScopePassthru
)

type MultiScope interface {
	GetValueByScope(scope string, name string) (value interface{}, exists bool)
}

type StageInputScope struct {
	execCtx *ExecutionContext
}

func (s *StageInputScope) GetValue(name string) (value interface{}, exists bool) {

	attrs := s.execCtx.currentOutput

	attr, found := attrs[name]

	if found {
		return attr, true
	}

	return attr, found
}

func (s *StageInputScope) SetValue(name string, value interface{}) error {
	return errors.New("read-only scope")
}

func (s *StageInputScope) GetValueByScope(scopeId ScopeId, name string) (value interface{}, exists bool) {

	attrs := s.execCtx.currentOutput

	switch scopeId {
	case ScopePipeline:
		attrs = s.execCtx.pipelineInput
	}

	attr, found := attrs[name]

	if found {
		return attr, true
	}

	return attr, found
}

// SimpleScope is a basic implementation of a scope
type StageOutputScope struct {
	execCtx *ExecutionContext
	output  map[string]interface{}
}

func (s *StageOutputScope) GetValue(name string) (value interface{}, exists bool) {
	attrs := s.execCtx.currentOutput

	attr, found := attrs[name]

	if found {
		return attr, true
	}
	attr, found = s.output[name]

	if found {
		return attr, true
	}

	return attr, found
}

func (s *StageOutputScope) SetValue(name string, value interface{}) error {
	return errors.New("read-only scope")
}

func (s *StageOutputScope) GetValueByScope(scopeId ScopeId, name string) (value interface{}, exists bool) {
	attrs := s.execCtx.currentOutput

	switch scopeId {
	case ScopePipeline:
		attrs = s.execCtx.pipelineInput
	}

	attr, found := attrs[name]

	if found {
		return attr, true
	}

	return attr, found
}
