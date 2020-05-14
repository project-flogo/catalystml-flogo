package ifin

import (
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

	a.logger.Info("Starting operation ifIn.")
	a.logger.Debug("Inputs for Operation ifIn.", inputs)

	arr0 := inputs["arr0"]
	arr1 := inputs["arr1"]

	// making hash to check against for if it is in
	checkmap := make(map[interface{}]bool)
	for _, val := range arr1.([]interface{}) {
		checkmap[val] = true
	}

	// appending to out if val is in checkmap/arr1
	var out []interface{}
	for _, val := range arr0.([]interface{}) {
		if checkmap[val] == true {
			out = append(out, val)
		}
	}

	a.logger.Info("Operation ifIn Completed.")
	a.logger.Debug("Output of Operation ifIn.", out)
	return out, nil
}
