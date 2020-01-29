package filter

import (
	"errors"
	"strconv"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/common"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	sortByKey bool
	params    *Params
	logger    log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	dataFrame, ok := in.Data.(*common.DataFrame)

	operation.logger.Info("Starting Operation Filter.")
	operation.logger.Debug("DataFrame is...", in.Data)
	operation.logger.Debug("Value is...", in.Value)
	operation.logger.Debug("FilterType is...", in.FilterType)
	operation.logger.Debug("Column is...", operation.params.Col)
	operation.logger.Debug("Axis is...", operation.params.Axis)

	if !ok {
		errors.New("Input data should be DataFrame type.")
	}

	var result *common.DataFrame

	operation.logger.Debug("Before filtering : ", dataFrame)
	result, err = operation.Filter(in.Value, in.FilterType, dataFrame)
	operation.logger.Debug("After filtered : ", result)

	operation.logger.Info("Operation Filter completed.")
	operation.logger.Debug("The output of Operation Filter, As DataFrame : ", result)

	return result.AsIs(), err
}

func (operation *Operation) Filter(
	targetValue interface{},
	filterType string,
	dataFrame *common.DataFrame,
) (*common.DataFrame, error) {
	newDataFrame := common.NewDataFrame()
	switch operation.params.Axis {
	case 0: /* Check column will keep/remove rows */
		common.ProcessDataFrame(dataFrame, func(tuple *common.SortableTuple, lastTuple bool) error {
			// Each tupple is a row of data
			if "Keep" == filterType {
				if tuple.GetByKey(operation.params.Col) != targetValue {
					return nil
				}
			} else {
				if tuple.GetByKey(operation.params.Col) == targetValue {
					return nil
				}
			}
			operation.logger.Debug(*tuple)
			common.SortableTupleAppendToDataframe(*tuple, newDataFrame)
			operation.logger.Debug(newDataFrame)
			return nil
		})
	case 1:
		for _, label := range dataFrame.GetLabels() {
			column := dataFrame.GetColumn(label)
			tuple := make(map[string]interface{})
			fieldOrder := make([]string, len(column))
			for rIndex, value := range column {
				tuple[strconv.Itoa(rIndex)] = value
				fieldOrder[rIndex] = strconv.Itoa(rIndex)
			}

			if "Keep" == filterType {
				if tuple[operation.params.Col] != targetValue {
					continue
				}
			} else {
				if tuple[operation.params.Col] == targetValue {
					continue
				}
			}
			sTuple := common.NewSortableTuple(tuple, fieldOrder)
			common.SortableTupleAppendToDataframe(*sTuple, newDataFrame)
			operation.logger.Debug(*sTuple)
		}
		operation.logger.Debug("Before Transpose ........... ")
		operation.logger.Debug(newDataFrame)
		operation.logger.Debug("After Transpose ........... ")
		newDataFrame = common.Transpose(newDataFrame, nil)
	}
	newDataFrame.SetFromTable(dataFrame.GetFromTable())

	return newDataFrame, nil
}
