package pipeline

import (
	"fmt"
	"time"

	"github.com/project-flogo/catalystml-flogo/action/types"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
)

type Instance struct {
	def *Definition
	id  string

	logger log.Logger
}

// NewInstance gets new instance from defination
func NewInstance(definition *Definition, id string, logger log.Logger) *Instance {

	return &Instance{def: definition, id: id, logger: logger}
}

func (inst *Instance) Id() string {
	return inst.id
}

// Run runs the instance of the CML.
func (inst *Instance) Run(input map[string]interface{}) (output map[string]interface{}, err error) {

	// Get the Scope of the CML pipeline.
	// Scope is the collection of the data in the CML
	scope, err := NewPipelineScope(input, inst.def.labels)

	if err != nil {
		return nil, err
	}

	// Log the time
	start := time.Now()

	//Check the type of the input of the pipeline.
	for key, _ := range inst.def.input {

		temp, ok := inst.def.input[key].(PipelineInput)
		if !ok {
			continue
		}

		err = types.ValidateType(temp.Type, input[key])

		if err != nil {
			return nil, err
		}

	}

	//Run the tasks.
	for key, task := range inst.def.tasks {
		task.Position()
		scope, err = task.Eval(scope, inst.logger)

		if err != nil {
			return nil, fmt.Errorf("Error %s in task \"%s-%v\" ", err.Error(), task.Name(), key)
		}

	}

	// Set the output.

	if inst.def.output.Data != nil {
		mf := GetMapperFactory()
		mappings := make(map[string]interface{})

		// Type Switch
		switch t := inst.def.output.Data.(type) {
		case map[string]interface{}:
			for key, val := range t {
				mappings[key] = val
			}
		default:
			mappings["data"] = inst.def.output.Data
		}

		// Get the data from output expression
		outMapper, err := mf.NewMapper(mappings)
		if err != nil {
			return nil, err
		}
		output, err = outMapper.Apply(scope)

		if err != nil {
			return nil, err
		}
		var definedType data.Type

		// Check if the output is defined as dataframe or map.
		if inst.def.output.Type == "dataframe" || inst.def.output.Type == "map" {
			definedType, err = data.ToTypeEnum("object")
			if err != nil {
				return nil, err
			}

			givenType, err := data.GetType(output)
			if err != nil {
				return nil, err
			}

			if definedType != givenType {
				return nil, fmt.Errorf("Type mismatch in output. Defined type [%s] passed type [%s]", definedType, givenType)
			}

			inst.logger.Infof("The output took %v to calculate", time.Since(start))

			return output, nil
		}

		definedType, _ = data.ToTypeEnum(inst.def.output.Type)

		for key, _ := range output {

			givenType, err := data.GetType(output[key])
			if err != nil {
				return nil, err
			}

			if definedType != givenType {
				return nil, fmt.Errorf("Type mismatch in output. Defined type [%s] passed type [%s]", definedType, givenType)
			}
		}

	}
	inst.logger.Infof("The output took %v to calculate", time.Since(start))

	return output, nil

}
