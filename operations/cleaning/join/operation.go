package join

import (
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

	operation.logger.Info("Starting Operation Join.")
	operation.logger.Debug("Input for Operation Join Left dataFrame.", in.Left)
	operation.logger.Debug("Input for Operation Join Right dataFrame. ", in.Right)
	operation.logger.Debug("Input for Operation Join Left Index map.", in.LeftIndex)
	operation.logger.Debug("Input for Operation Join Right Index map.", in.RightIndex)
	operation.logger.Debug("Input for Operation Join Parameter. ", operation.params)

	result, err = operation.join(
		in.Left.(map[string][]interface{}),
		in.Right.(map[string][]interface{}),
		in.LeftIndex.([]string),
		in.RightIndex.([]string),
	)

	operation.logger.Info("Operation Join Completed.")
	operation.logger.Debug("Output of Operation Join Joined dataFrame is : ", result)

	return result, err
}

func (operation *Operation) join(
	leftDataFrame map[string][]interface{},
	rightDataFrame map[string][]interface{},
	leftIndex []string,
	rightIndex []string,
) (result map[string][]interface{}, err error) {

	dataFrame01 := leftDataFrame
	dataFrame02 := rightDataFrame
	index01 := leftIndex
	index02 := rightIndex
	if "right" == operation.params.How {
		dataFrame02 = leftDataFrame
		dataFrame01 = rightDataFrame
		index02 = leftIndex
		index01 = rightIndex
	}

	/* dataFrame01 join dataFrame02 */
	dataFrame := make(map[string][]interface{})

	/* check right tuple size */
	tupleSize := -1
	for fieldname, filedsArray := range dataFrame02 {
		tupleSize = len(filedsArray)
		if nil == dataFrame[fieldname] {
			dataFrame[fieldname] = nil
		}
	}

	//	newDataFrame := make(map[string][]interface{})
	dataSet02 := make(map[common.Index]map[string]interface{})
	var tuple map[string]interface{}
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		tuple = make(map[string]interface{})
		for fieldname, filedsArray := range dataFrame02 {
			tuple[fieldname] = filedsArray[i]
		}
		dataSet02[GetKey(index02, tuple)] = tuple
		operation.logger.Debug("Tuple - ", tuple, ", DataSet02 - ", dataSet02)
	}

	tupleSize = -1
	for fieldname, filedsArray := range dataFrame01 {
		tupleSize = len(filedsArray)
		if nil == dataFrame[fieldname] {
			dataFrame[fieldname] = nil
		}
	}

	dataSet := make(map[common.Index]map[string]interface{})
	for i := 0; i < tupleSize; i++ {
		tuple = make(map[string]interface{})
		for fieldname, fieldsArray := range dataFrame01 {
			tuple[fieldname] = fieldsArray[i]
		}

		key := GetKey(index02, tuple)
		tuple02 := dataSet02[key]
		dataSet02[key] = nil
		if 0 != len(tuple02) || "inner" != operation.params.How {
			dataSet[GetKey(index01, tuple)] = tuple
			for fieldname, fieldvalue := range tuple02 {
				tuple[fieldname] = fieldvalue
			}
		}

		operation.logger.Debug("Tuple - tuple01 = ", tuple, ", tuple02 = ", tuple02, ", len(tuple02) = ", len(tuple02))
	}

	if "outer" == operation.params.How {
		for key, tuple := range dataSet02 {
			if nil != tuple {
				dataSet[key] = tuple
			}
		}
	}
	operation.logger.Info("DataSet - ", dataSet)

	return operation.dataSetToDataFrame(dataSet, dataFrame)
}

func (operation *Operation) dataSetToDataFrame(
	dataSet map[common.Index]map[string]interface{},
	newDataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {

	counter := 0
	for _, tuple := range dataSet {
		for column, columnValues := range newDataFrame {
			if nil == columnValues {
				columnValues = make([]interface{}, len(dataSet))
				newDataFrame[column] = columnValues
			}
			if nil != tuple[column] {
				columnValues[counter] = tuple[column]
			}
		}
		counter++
	}

	return newDataFrame, nil
}

func GetKey(indexColumn []string, tuple map[string]interface{}) common.Index {
	index := make([]interface{}, len(indexColumn))
	for i, element := range indexColumn {
		index[i] = tuple[element]
	}
	return common.NewIndex(index)
}
