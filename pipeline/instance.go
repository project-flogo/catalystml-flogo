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
	ctx := &ExecutionContext{pipeline: inst}
	ctx.pipelineInput = input

	ctx.currentOutput = input

	output = make(map[string]interface{})
	for _, stage := range inst.def.stages {

		if stage.inputMapper != nil {

			in := &StageInputScope{execCtx: ctx}
			ctx.currentInput, err = stage.inputMapper.Apply(in)
			if err != nil {
				return nil, err
			}

		}

		for key, val := range ctx.pipelineInput {
			temp, ok := inst.def.input[key].(PipelineInput)
			//fmt.Println("Val...", val)
			if !ok {
				continue
			}

			definedType, _ := data.ToTypeEnum(temp.Type)
			givenType, _ := data.GetType(val.(map[string]interface{})[temp.Label])
			if definedType != givenType {
				return nil, fmt.Errorf("Type mismatch in input. Defined type [%s] passed type [%s]", definedType, givenType)
			}
		}

		err := stage.opt.Eval(ctx)
		if err != nil {
			return nil, err
		}

		in := &StageOutputScope{execCtx: ctx}

		results, err := stage.outputMapper.Apply(in)

		for name, value := range results {

			_, ok := stage.outputAttrs[name]
			if !ok {
				output[name] = value
				ctx.currentOutput[name] = value
			} else {
				output[stage.outputAttrs[name].(string)] = value
				ctx.currentOutput[stage.outputAttrs[name].(string)] = value
			}

		}

		if err != nil {
			return nil, err
		}

	}

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
		output, err = outMapper.Apply(&StageOutputScope{output: output, execCtx: ctx})

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

	return output, nil

}
