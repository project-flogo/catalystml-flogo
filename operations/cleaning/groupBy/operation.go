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

func (this *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	var result interface{}

	this.logger.Info("Input dataFrame is : ", in.Data)
	this.logger.Info("Parameter is : ", this.params)

	result, err = this.groupBy(in.Data.(map[string][]interface{}))

	this.logger.Info("Grouped dataFrame is : ", result)

	return result, err
}

func (this *Operation) groupBy(dataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {

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

		outputKeyValue := tuple[this.params.Index[this.params.Level]].(string)
		data := groupedData[outputKeyValue]
		if nil == data {
			data = common.GetFunction(this.params.Function)
			groupedData[outputKeyValue] = data
		}
		data.Update(tuple[this.params.Target])

		this.logger.Debug("Tuple - ", tuple, ", groupedData - ", groupedData)
	}

	return this.tupleToDataFrame(groupedData)
}

func (this *Operation) groupedKey() string {
	var groupKey bytes.Buffer
	groupKey.WriteString(this.params.Function)
	groupKey.WriteString("_")
	groupKey.WriteString(this.params.Target)
	return groupKey.String()
}

func (this *Operation) tupleToDataFrame(
	groupedData map[string]common.DataState) (result map[string][]interface{}, err error) {
	newDataFrame := make(map[string][]interface{})
	tupleSize := len(groupedData)
	dataKey := this.params.Index[this.params.Level]
	groupedKey := this.groupedKey()
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
