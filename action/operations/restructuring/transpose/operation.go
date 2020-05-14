package transpose

import (
	"strconv"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result *common.DataFrame

	a.logger.Info("Starting Operation Transpose.")
	a.logger.Debug("Input for Operation Transpose. Matrix. ", in.Data)

	data, err := common.ToDataFrame(in.Data)
	if nil != err {
		return nil, err
	}
	result, err = transpose(data)

	a.logger.Info("Operation Transpose Completed.")
	a.logger.Debug("Output for Operation Transpose. Matrix.", result)

	return result.AsIs(), err
}

func transpose(dataFrame *common.DataFrame) (result *common.DataFrame, err error) {

	newDataFrame := common.NewDataFrame()

	counter := 0
	common.ProcessDataFrame(dataFrame, func(sTuple *common.SortableTuple, lastTuple bool) error {
		order := sTuple.GetDataArray()
		newDataFrame.AddColumn(strconv.Itoa(counter), order)
		counter++
		return nil
	})

	newDataFrame.SetFromTable(dataFrame.GetFromTable())
	return newDataFrame, nil
}
