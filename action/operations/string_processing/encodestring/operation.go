package encodestring

import (
	b64 "encoding/base64"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Starting operation encodestring.")
	a.logger.Debug("Input for operation encodestring.", input.Str)

	out := b64.StdEncoding.EncodeToString([]byte(input.Str))

	a.logger.Info("Operation encodestring completed.")
	a.logger.Debug("Output of operation encodestring...", out)
	return out, nil
}
