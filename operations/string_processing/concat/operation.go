package concat

import (
	// "strings"
	"bytes"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	// params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Executing operation concat...", input.S0, " to ", input.S1)

	var b bytes.Buffer
	if input.S0 != "" {
		b.WriteString(input.S0)
	}
	if input.S1 != "" {
		b.WriteString(input.S1)
	}
	if len(input.Slist) > 0 {
		for _, s := range input.Slist {
			b.WriteString(s.(string))
		}
	}
	out := b.String()
	a.logger.Debug("result of concat...", out)

	return out, nil
}
