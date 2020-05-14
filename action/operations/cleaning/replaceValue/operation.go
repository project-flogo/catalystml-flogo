package replaceValue

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params []interface{}
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	params, err := coerce.ToArray(p.Columns)

	if err != nil {
		return nil, err
	}

	return &Operation{params: params, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	in := &Input{}
	err := in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	a.logger.Info("Starting Operation Replace Value.")
	a.logger.Debug("Input of Operation Replace Value.", in.Data)

	for _, val := range a.params {

		if arr, ok := in.Data[val.(string)]; ok {

			temp, err := coerce.ToArrayIfNecessary(arr)

			if err != nil {
				return nil, err
			}
			if arr, ok := temp.([]interface{}); ok {
				for key, tmp := range arr {

					if _, ok = in.ReplaceMap[tmp.(string)]; ok {
						arr[key] = in.ReplaceMap[tmp.(string)]
					}
				}
				in.Data[val.(string)] = arr
			} else {
				in.Data[val.(string)] = in.ReplaceMap[temp.(string)]
			}

		}

	}
	a.logger.Info("Operation Replace Value Completed.")
	a.logger.Debug("Output of Operation Replace Value.", in.Data)

	return in.Data, nil
}
