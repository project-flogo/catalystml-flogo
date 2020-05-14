package pivot

import (
	"bytes"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
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

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result *common.DataFrame

	operation.logger.Info("Starting Operation Pivot.")
	operation.logger.Debug("Input dataFrame is : ", in.Data)
	operation.logger.Debug("Parameter is : ", operation.params)

	data, err := common.ToDataFrame(in.Data)
	if nil != err {
		return nil, err
	}
	result, err = operation.pivot(data)

	operation.logger.Debug("Pivoted dataFrame is : ", result)
	operation.logger.Info("Operation Pivot Completed.")

	return result.AsIs(), err
}

func (operation *Operation) pivot(dataFrame *common.DataFrame) (result *common.DataFrame, err error) {

	newDataFrame := common.NewDataFrame()
	aggregatedTupleMap := make(map[common.Index]map[string]common.DataState)
	var key []interface{}
	common.ProcessDataFrame(dataFrame, func(sTuple *common.SortableTuple, lastTuple bool) error {
		tuple := sTuple.GetData()

		/* build key for output data*/
		key = make([]interface{}, len(operation.params.Index))
		for j, keyElement := range operation.params.Index {
			key[j] = tuple[keyElement]
		}

		index := common.NewIndex(key)
		aggregatedTuple := aggregatedTupleMap[index]
		if nil == aggregatedTuple {
			aggregatedTuple = make(map[string]common.DataState)
		}

		for _, keyElement := range operation.params.Index {
			data := aggregatedTuple[keyElement]
			if nil == data {
				data = &common.Data{}
				aggregatedTuple[keyElement] = data
			}
			data.Update(tuple[keyElement])
		}

		operation.aggregate(tuple, aggregatedTuple)
		aggregatedTupleMap[index] = aggregatedTuple

		operation.logger.Debug("Tuple - ", tuple, ", aggregatedTuple - ", aggregatedTuple)

		if lastTuple {
			for _, aggregatedTuple := range aggregatedTupleMap {
				newTuple := make(map[string]interface{})

				for key, value := range aggregatedTuple {
					newTuple[key] = value.Value()
				}
				common.TupleAppendToDataframe(newTuple, newDataFrame)
			}
		}

		return nil
	})

	newDataFrame.SetFromTable(dataFrame.GetFromTable())
	return newDataFrame, nil
}

func (operation *Operation) aggregate(
	tuple map[string]interface{},
	aggregatedTuple map[string]common.DataState,
) {
	for valueColumn, functionNames := range operation.params.Aggregate {
		for _, functionName := range functionNames {
			dataKey := operation.dataKey(tuple, functionName, valueColumn)
			function := aggregatedTuple[dataKey]
			if nil == function {
				function = common.GetFunction(functionName)
				aggregatedTuple[dataKey] = function
			}
			err := function.Update(tuple[valueColumn])
			if nil != err {
				operation.logger.Info("Error : ", err)
			}
		}
	}
}

func (operation *Operation) dataKey(
	tuple map[string]interface{},
	functionName string,
	valueColumn string,
) string {
	var groupKey bytes.Buffer
	for _, group := range operation.params.Columns {
		groupKey.WriteString(tuple[group].(string))
		groupKey.WriteString("_")
	}
	groupKey.WriteString(functionName)
	groupKey.WriteString("_")
	groupKey.WriteString(valueColumn)
	return groupKey.String()
}
