package pipeline

import (
	"strconv"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/data"
)

type DefinitionConfig struct {
	Name   string          `json:"name"`
	Stages []*StageConfig  `json:"structure"`
	Input  []PipelineInput `json:"input"`
	Output PipelineOutput  `json:"output"`
}

func NewDefinition(config *DefinitionConfig, mf mapper.Factory, resolver resolve.CompositeResolver) (*Definition, error) {

	def := &Definition{name: config.Name, output: config.Output}

	for _, sconfig := range config.Stages {
		stage, err := NewStage(sconfig, mf, resolver)

		if err != nil {
			return nil, err
		}

		def.stages = append(def.stages, stage)
	}
	def.input = make(map[string]interface{})
	def.labels = make(map[string]interface{})
	for key, val := range config.Input {
		switch t := val.Label.(type) {
		case string:
			def.input[t] = val
		default:
			def.labels[strconv.Itoa(key)] = val.Label
		}
	}

	return def, nil
}

type Definition struct {
	name   string
	stages []*Stage
	input  map[string]interface{}
	labels map[string]interface{}
	output PipelineOutput
}

func (d *Definition) Name() string {
	return d.name
}

func (d *Definition) MetaData() *metadata.IOMetadata  {
	
	result := make(map[string]data.TypedValue)
	for key, _ := range d.input {
		result[key] = nil
	}
	
	return &metadata.IOMetadata{Input: result, Output: nil}
}
