package split

import (
	"strings"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (this *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	this.logger.Info("before : in.Text = ", in.Text, ", in.Separator = ", in.Separator)

	stringArray := strings.Split(in.Text, in.Separator)

	this.logger.Info("after : stringArray = ", stringArray)

	return stringArray, nil
}
