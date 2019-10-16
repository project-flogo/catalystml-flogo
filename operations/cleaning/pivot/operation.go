package pivot

import (
	"bytes"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/common"
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

	var result interface{}

	operation.logger.Debug("Input dataFrame is : ", in.Data)
	operation.logger.Debug("Parameter is : ", operation.params)

	result, err = operation.pivot(in.Data.(map[string][]interface{}))

	operation.logger.Debug("Pivoted dataFrame is : ", result)

	return result, err
}

func (operation *Operation) pivot(dataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {

	/* check tuple size */
	tupleSize := -1
	for _, filedsArray := range dataFrame {
		tupleSize = len(filedsArray)
		if 0 < tupleSize {
			break
		}
	}

	newDataFrame := make(map[string][]interface{})
	aggregatedTupleMap := make(map[common.Index]map[string]common.DataState)
	tuple := make(map[string]interface{})
	var key []interface{}
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		for fieldname, filedsArray := range dataFrame {
			tuple[fieldname] = filedsArray[i]
		}

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
			newDataFrame[keyElement] = nil
		}

		operation.aggregate(tuple, aggregatedTuple, newDataFrame)
		aggregatedTupleMap[index] = aggregatedTuple

		operation.logger.Debug("Tuple - ", tuple, ", aggregatedTuple - ", aggregatedTuple)
	}

	return operation.transform(aggregatedTupleMap, newDataFrame)
}

func (operation *Operation) aggregate(
	tuple map[string]interface{},
	aggregatedTuple map[string]common.DataState,
	newDataFrame map[string][]interface{},
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
			newDataFrame[dataKey] = nil
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

func (operation *Operation) transform(
	tupleMap map[common.Index]map[string]common.DataState,
	newDataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {
	counter := 0
	for _, tuple := range tupleMap {
		for column, columnValus := range newDataFrame {
			if nil == columnValus {
				columnValus = make([]interface{}, len(tupleMap))
				newDataFrame[column] = columnValus
			}
			if nil != tuple[column] {
				columnValus[counter] = tuple[column].Value()
			}
		}
		counter++
	}

	return newDataFrame, nil
}
