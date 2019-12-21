package runcml

import (
	"context"
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/action"
)

type Operation struct {
	logger log.Logger
	act action.Action
}



func New(ctx operation.InitContext) (operation.Operation, error) {

	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	factory := action.GetFactory("github.com/project-flogo/catalystml-flogo/action")

	act, err := factory.New(&action.Config{Settings: map[string]interface{}{"catalystMlURI": p.CatalystMlURI}})

	if err != nil {
		return nil, err
	}

	return &Operation{logger: ctx.Logger(), act: act}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	in := &Input{}

	err := in.FromMap(inputs)

	if err != nil {
		return nil, err
	}

	a.logger.Info("Starting operation Run CMl.")
	a.logger.Debug("The inputs of operation Run CML.", inputs)

	out, err := a.act.(action.SyncAction).Run(context.Background(), in.Data)

	a.logger.Info("Operation Run CML completed.")
	a.logger.Debug("The output of operation Run CML.", out)

	return out, err

}