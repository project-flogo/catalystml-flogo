package pivot

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = operation.Register(&Operation{}, New)
}

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

	result, err = this.pivot(in.Data.(map[string][]interface{}))

	this.logger.Info("Pivoted dataFrame is : ", result)

	return result, err
}

func (this *Operation) pivot(dataFrame map[string][]interface{}) (result map[string][]interface{}, err error) {

	/* check tuple size */
	tupleSize := -1
	for _, filedsArray := range dataFrame {
		tupleSize = len(filedsArray)
		if 0 < tupleSize {
			break
		}
	}

	newDataFrame := make(map[string][]interface{})
	aggregatedTupleMap := make(map[Index]map[string]DataState)
	tuple := make(map[string]interface{})
	var key []interface{}
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		for fieldname, filedsArray := range dataFrame {
			tuple[fieldname] = filedsArray[i]
		}

		/* build key for output data*/
		key = make([]interface{}, len(this.params.Index))
		for j, keyElement := range this.params.Index {
			key[j] = tuple[keyElement]
		}

		index := NewIndex(key)
		aggregatedTuple := aggregatedTupleMap[index]
		if nil == aggregatedTuple {
			aggregatedTuple = make(map[string]DataState)
		}

		for _, keyElement := range this.params.Index {
			data := aggregatedTuple[keyElement]
			if nil == data {
				data = &Data{}
				aggregatedTuple[keyElement] = data
			}
			data.Update(tuple[keyElement])
			newDataFrame[keyElement] = nil
		}

		this.aggregate(tuple, aggregatedTuple, newDataFrame)
		aggregatedTupleMap[index] = aggregatedTuple

		this.logger.Debug("Tuple - ", tuple, ", aggregatedTuple - ", aggregatedTuple)
	}

	return this.transform(aggregatedTupleMap, newDataFrame)
}

func (this *Operation) aggregate(
	tuple map[string]interface{},
	aggregatedTuple map[string]DataState,
	newDataFrame map[string][]interface{},
) {
	for valueColumn, functionNames := range this.params.Aggregate {
		for _, functionName := range functionNames {
			dataKey := this.dataKey(tuple, functionName, valueColumn)
			function := aggregatedTuple[dataKey]
			if nil == function {
				function = getFunction(functionName)
				aggregatedTuple[dataKey] = function
			}
			err := function.Update(tuple[valueColumn])
			if nil != err {
				this.logger.Info("Error : ", err)
			}
			newDataFrame[dataKey] = nil
		}
	}
}

func (this *Operation) dataKey(
	tuple map[string]interface{},
	functionName string,
	valueColumn string,
) string {
	var groupKey bytes.Buffer
	for _, group := range this.params.Columns {
		groupKey.WriteString(tuple[group].(string))
		groupKey.WriteString("_")
	}
	groupKey.WriteString(functionName)
	groupKey.WriteString("_")
	groupKey.WriteString(valueColumn)
	return groupKey.String()
}

func (this *Operation) transform(
	tupleMap map[Index]map[string]DataState,
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

type DataState interface {
	Update(newData interface{}) error
	Value() interface{}
}

func getFunction(functionName string) DataState {
	var function DataState
	if "sum" == functionName {
		function = &Sum{}
	} else if "count" == functionName {
		function = &Count{}
	} else if "mean" == functionName {
		function = &Mean{}
	} else if "min" == functionName {
		function = &Min{}
	} else if "max" == functionName {
		function = &Max{}
	}
	return function
}

type Data struct {
	data interface{}
}

func (this *Data) Value() interface{} {
	return this.data
}

func (this *Data) Update(newData interface{}) error {
	this.data = newData
	return nil
}

type Sum struct {
	data float64
}

func (this *Sum) Value() interface{} {
	return this.data
}

func (this *Sum) Update(newData interface{}) error {
	delta, _ := coerce.ToFloat64(newData)
	this.data += delta
	return nil
}

type Count struct {
	counter int
}

func (this *Count) Value() interface{} {
	return this.counter
}

func (this *Count) Update(newData interface{}) error {
	this.counter += 1
	return nil
}

type Mean struct {
	sum   float64
	count float64
}

func (this *Mean) Value() interface{} {
	return this.sum / this.count
}

func (this *Mean) Update(newData interface{}) error {
	this.count += 1
	delta, err := coerce.ToFloat64(newData)
	if nil != err {
		return err
	}
	this.sum += delta
	return nil
}

type Min struct {
	min interface{}
}

func (this *Min) Value() interface{} {
	return this.min
}

func (this *Min) Update(newData interface{}) error {
	if nil == this.min {
		this.min = newData
		return nil
	}

	result, err := compare(newData, this.min)

	if nil != err {
		return err
	}

	if 0 > result {
		this.min = newData
	}
	return nil
}

type Max struct {
	max interface{}
}

func (this *Max) Value() interface{} {
	return this.max
}

func (this *Max) Update(newData interface{}) error {
	if nil == this.max {
		this.max = newData
		return nil
	}

	result, err := compare(newData, this.max)

	if nil != err {
		return err
	}

	if 0 < result {
		this.max = newData
	}
	return nil
}

func compare(data1 interface{}, data2 interface{}) (int, error) {

	switch data1.(type) {
	case float64:
		delta1float64, _ := coerce.ToFloat64(data1)
		delta2float64, err := coerce.ToFloat64(data2)
		if nil != err {
			return 0, err
		}
		delta := delta1float64 - delta2float64
		switch {
		case delta > 0:
			return 1, nil
		case delta == 0:
			return 0, nil
		case delta < 0:
			return -1, nil
		}
	case int:
		delta1int, _ := coerce.ToInt(data1)
		delta2int, err := coerce.ToInt(data2)
		if nil != err {
			return 0, err
		}
		delta := delta1int - delta2int
		switch {
		case delta > 0:
			return 1, nil
		case delta == 0:
			return 0, nil
		case delta < 0:
			return -1, nil
		}
	}

	return 0, errors.New("Unable to compare, Uknown type!")
}

type Index struct {
	Id uint64
}

func NewIndex(elements []interface{}) Index {
	keyBytes := []byte{}
	for _, element := range elements {
		elementBytes, _ := json.Marshal(element)
		keyBytes = append(keyBytes, elementBytes...)
	}
	hasher := md5.New()
	hasher.Write(keyBytes)
	return Index{Id: binary.BigEndian.Uint64(hasher.Sum(nil))}
}
