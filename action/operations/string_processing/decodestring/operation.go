package decodestring

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

	a.logger.Info("Starting operation decodestring.")
	a.logger.Debug("Input for operation decodestring.", input.Str)

	outBytes, err := b64.StdEncoding.DecodeString(input.Str)
	if nil != err {
		return nil, err
	}
	out := string(outBytes)

	a.logger.Info("Operation decodestring completed.")
	a.logger.Debug("Output of operation decodestring...", out)
	return out, nil
}
