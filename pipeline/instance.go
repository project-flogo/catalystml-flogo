package pipeline

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
)

type Instance struct {
	def *Definition
	id  string

	logger log.Logger
}

func NewInstance(definition *Definition, id string, logger log.Logger) *Instance {

	return &Instance{def: definition, id: id, logger: logger}
}

func (inst *Instance) Id() string {
	return inst.id
}

func (inst *Instance) Run(input map[string]interface{}) (output map[string]interface{}, err error) {
	currentInput := make(map[string]interface{})

	scope := NewPipelineScope(input)

	//Check the type of the input of the pipeline.
	for key, _ := range inst.def.input {

		temp, ok := inst.def.input[key].(PipelineInput)
		if !ok {
			continue
		}

		definedType, _ := data.ToTypeEnum(temp.Type)
		givenType, _ := data.GetType(input[key])
		if definedType != givenType {
			return nil, fmt.Errorf("Type mismatch in input. Defined type [%s] passed type [%s]", definedType, givenType)
		}
	}

	//Execute the pipeline.
	for _, stage := range inst.def.stages {

		inst.logger.Debug("Operation Input Mapper: ", stage.inputMapper)
		if stage.inputMapper != nil {

			currentInput, err = stage.inputMapper.Apply(scope)
			if err != nil {
				return nil, err
			}

		}

		inst.logger.Debug("Executing operation with inputs: ", currentInput)
		stageOutput, err := stage.opt.Eval(currentInput)

		if err != nil {
			return nil, err
		}

		scope.SetValue(stage.ID(), stageOutput)
		inst.logger.Debug("Setting operation outputs: ", stageOutput)

		_, err = stage.outputMapper.Apply(scope)

		if err != nil {
			return nil, err
		}

	}
	output = scope.values

	if inst.def.output.Data != nil {
		mf := GetMapperFactory()
		mappings := make(map[string]interface{})

		switch t := inst.def.output.Data.(type) {
		case map[string]interface{}:
			for key, val := range t {
				mappings[key] = val
			}
		default:
			mappings["data"] = inst.def.output.Data
		}

		outMapper, err := mf.NewMapper(mappings)
		output, err = outMapper.Apply(scope)

		if err != nil {
			return nil, err
		}
		var definedType data.Type
		if inst.def.output.Type == "dataframe" || inst.def.output.Type == "map" {
			definedType, _ = data.ToTypeEnum("object")
		} else {
			definedType, _ = data.ToTypeEnum(inst.def.output.Type)
		}

		givenType, _ := data.GetType(output)

		if definedType != givenType {
			return nil, fmt.Errorf("Type mismatch in output. Defined type [%s] passed type [%s]", definedType, givenType)
		}
	}
	inst.logger.Debug("Output of the action is...", output)
	return output, nil

}
