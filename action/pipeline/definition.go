package pipeline

import (
	"strconv"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/resolve"
)

type DefinitionConfig struct {
	Name   string          `json:"name"`
	Tasks  []TaskConfig    `json:"structure"`
	Input  []PipelineInput `json:"input"`
	Output PipelineOutput  `json:"output"`
}

// NewDefinition gets a defination struct from config.
func NewDefinition(config *DefinitionConfig, mf mapper.Factory, resolver resolve.CompositeResolver) (*Definition, error) {

	def := &Definition{name: config.Name, output: config.Output}

	for _, Tasks := range config.Tasks {

		// Get a New Task from config.
		task, err := NewTask(Tasks, mf, resolver)

		if err != nil {
			return nil, err
		}

		//Append the task to the array.
		def.tasks = append(def.tasks, task)

	}
	def.input = make(map[string]interface{})

	def.labels = make(map[string]interface{})

	// Set the labels from the input
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

// Defination struct
type Definition struct {
	name   string
	tasks  []Task
	input  map[string]interface{}
	labels map[string]interface{}
	output PipelineOutput
}

func (d *Definition) Name() string {
	return d.name
}

func (d *Definition) MetaData() *metadata.IOMetadata {

	result := make(map[string]data.TypedValue)
	for key, _ := range d.input {
		result[key] = nil
	}

	return &metadata.IOMetadata{Input: result, Output: nil}
}
