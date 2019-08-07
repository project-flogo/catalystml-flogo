package pipeline

import (
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
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

	for _, val := range config.Input {
		def.input[val.Label] = val
	}

	return def, nil
}

type Definition struct {
	name   string
	stages []*Stage
	input  map[string]interface{}
	output PipelineOutput
}

func (d *Definition) Name() string {
	return d.name
}
