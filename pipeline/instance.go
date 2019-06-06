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
	ctx := &ExecutionContext{discriminator: "discriminator", pipeline: inst}
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

		//For the input of operation... check if the type of the value coming
		//in matches to the type defined in input of the catalystML.
		//Also do we need this ?
		for key, val := range ctx.pipelineInput {
			temp, ok := inst.def.input[key].(PipelineInput)
			fmt.Println("Val...", val)
			if !ok {
				continue
			}

			definedType, _ := data.ToTypeEnum(temp.Type)
			givenType, _ := data.GetType(val)
			if definedType == givenType {
				fmt.Println("matched....")
			}
		}

		_, err := stage.opt.Eval(ctx)
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
		outMapper, err := mf.NewMapper(inst.def.output.Data)

		output, err = outMapper.Apply(&StageInputScope{execCtx: ctx})

		fmt.Println("Output")
		if err != nil {
			return nil, err
		}
	}

	return output, nil

}
