package pipeline

import (
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

	for _, stage := range inst.def.stages {

		if stage.inputMapper != nil {

			in := &StageInputScope{execCtx: ctx}
			ctx.currentInput, err = stage.inputMapper.Apply(in)
			if err != nil {
				return nil, err
			}

		}
		_, err := stage.opt.Eval(ctx)
		if err != nil {
			return nil, err
		}

		in := &StageOutputScope{execCtx: ctx}

		results, err := stage.outputMapper.Apply(in)

		for name, value := range results {

			ctx.currentOutput[name] = value
		}

		if err != nil {
			return nil, err
		}

	}
	return nil, nil

}
