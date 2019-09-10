package groupBy

import (
	"bytes"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/common"
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

	operation.logger.Info("Input dataFrame is : ", in.Data)
	operation.logger.Info("Parameter is : ", operation.params)

	result, err = operation.groupBy(in.Data.(map[string][]interface{}))

	operation.logger.Info("Grouped dataFrame is : ", result)

	return result, err
}

func (operation *Operation) groupBy(dataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {

	/* check tuple size */
	tupleSize := -1
	for _, filedsArray := range dataFrame {
		tupleSize = len(filedsArray)
		if 0 < tupleSize {
			break
		}
	}

	groupedData := make(map[string]common.DataState)
	tuple := make(map[string]interface{})
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		for fieldname, filedsArray := range dataFrame {
			tuple[fieldname] = filedsArray[i]
		}

		outputKeyValue := tuple[operation.params.Index[operation.params.Level]].(string)
		data := groupedData[outputKeyValue]
		if nil == data {
			data = common.GetFunction(operation.params.Function)
			groupedData[outputKeyValue] = data
		}
		data.Update(tuple[operation.params.Target])

		operation.logger.Debug("Tuple - ", tuple, ", groupedData - ", groupedData)
	}

	return operation.tupleToDataFrame(groupedData)
}

func (operation *Operation) groupedKey() string {
	var groupKey bytes.Buffer
	groupKey.WriteString(operation.params.Function)
	groupKey.WriteString("_")
	groupKey.WriteString(operation.params.Target)
	return groupKey.String()
}

func (operation *Operation) tupleToDataFrame(
	groupedData map[string]common.DataState) (result map[string][]interface{}, err error) {
	newDataFrame := make(map[string][]interface{})
	tupleSize := len(groupedData)
	dataKey := operation.params.Index[operation.params.Level]
	groupedKey := operation.groupedKey()
	newDataFrame[dataKey] = make([]interface{}, tupleSize)
	newDataFrame[groupedKey] = make([]interface{}, tupleSize)

	counter := 0
	for key, data := range groupedData {
		newDataFrame[dataKey][counter] = key
		newDataFrame[groupedKey][counter] = data.Value()
		counter++
	}

	return newDataFrame, nil
}
