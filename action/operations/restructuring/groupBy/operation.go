package groupBy

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

	operation.logger.Info("Starting Operation GroupBy.")
	operation.logger.Debug("Input dataFrame is : ", in.Data)
	operation.logger.Debug("Parameter is : ", operation.params)

	data, err := common.ToDataFrame(in.Data)
	if nil != err {
		return nil, err
	}
	result, err = operation.groupBy(data)

	operation.logger.Debug("Grouped dataFrame is : ", result)
	operation.logger.Info("Operation GroupBy Completed.")

	return result.AsIs(), err
}

func (operation *Operation) groupBy(dataFrame *common.DataFrame) (result *common.DataFrame, err error) {
	var keyColumns []string
	// Use all index columns, if level < 0
	if 0 > operation.params.Level {
		keyColumns = operation.params.Index
	} else {
		keyColumns = make([]string, 1)
		keyColumns[0] = operation.params.Index[operation.params.Level]
	}

	newDataFrame := common.NewDataFrame()
	aggregatedTupleByGroup := make(map[common.Index]map[string]common.DataState)
	var key []interface{}
	common.ProcessDataFrame(dataFrame, func(sTuple *common.SortableTuple, lastTuple bool) error {
		tuple := sTuple.GetData()
		key = make([]interface{}, len(keyColumns))
		for j, keyElement := range keyColumns {
			key[j] = tuple[keyElement]
		}

		index := common.NewIndex(key)
		aggregatedTuple := aggregatedTupleByGroup[index]
		if nil == aggregatedTuple {
			aggregatedTuple = make(map[string]common.DataState)
			for _, keyColumn := range keyColumns {
				keyData := &common.First{}
				keyData.Update(tuple[keyColumn])
				aggregatedTuple[keyColumn] = keyData
			}
			aggregatedTupleByGroup[index] = aggregatedTuple
		}

		operation.aggregate(tuple, aggregatedTuple)

		if lastTuple {
			for _, aggregatedTuple := range aggregatedTupleByGroup {
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
	groupKey.WriteString(functionName)
	groupKey.WriteString("_")
	groupKey.WriteString(valueColumn)
	return groupKey.String()
}
