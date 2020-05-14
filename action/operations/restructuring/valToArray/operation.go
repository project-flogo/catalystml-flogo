package valToArray

import (
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	in := &Input{}

	in.FromMap(inputs)
	a.logger.Info("Starting Operation valToArray.")
	a.logger.Debug("Input of Operation valToArray.", in.Value, in.Shape)

	out, err := strucBuilder(in.Value, in.Shape)
	if err != nil {
		return nil, err
	}
	a.logger.Info("Finishing Operation valToArray.")
	a.logger.Debug("Output of Operation valToArray.", out)
	return out, nil
}

func strucBuilder(val interface{}, shape []interface{}) (out interface{}, err error) {

	if len(shape) > 0 {
		var struc []interface{}
		thisdim := shape[0].(int)
		for i := 0; i < thisdim; i++ {
			thispos, _ := strucBuilder(val, shape[1:])
			struc = append(struc, thispos)
		}
		out = struc
	} else {
		out = val
	}

	return out, nil
}
