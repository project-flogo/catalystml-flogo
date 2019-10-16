package common

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/project-flogo/core/data/coerce"
)

type DataFrame map[string][]interface{}

func ToDataFrame(data interface{}) (DataFrame, error) {
	// This function takes an slices, matrix, tensors, and maps to a dataframe, does not handle other type caases

	out := make(DataFrame)
	var err error

	switch v := data.(type) {
	case []interface{}, [][]interface{}, [][][]interface{}, [][][][]interface{}, [][][][][]interface{}, 
	[]int, [][]int,[][][]int,[][][][]int,[][][][][]int, 
	[]float64, [][]float64,[][][]float64,[][][][]float64,[][][][][]float64,
	[]string, [][]string,[][][]string,[][][][]string,[][][][][]string :
		//Test dimensionality of matrix
		fmt.Println("data in is a slice", v)
		vcon, err := ToInterfaceArray(v)
		if err != nil {
			return nil, err
		}

		e := 0
		for i, val := range vcon {
			s := fmt.Sprintf("%d", i)
			switch w := val.(type) {
			case string, int, int32, int64, float32, float64, bool:
				out[s] = []interface{}{w}
				e = 1
			default:
				break
			}
		}
		if e == 1 {
			return out, nil
		}

		var tmp [][]interface{}
		for _, row := range vcon {
			var tmpr []interface{}
			for _, colval := range row.([]interface{}) {
				tmpr = append(tmpr, colval) //vcon[i].([]interface{})[j]
			}
			tmp = append(tmp, tmpr)
		}

		for i := 0; i < len(tmp[0]); i++ {
			var arr []interface{}
			for j := range tmp {
				arr = append(arr, tmp[j][i])
			}
			s := fmt.Sprintf("%d", i)
			out[s] = arr
		}

	case map[string]interface{}:
		fmt.Println("data in is a map", v)

		l := -1
		for key, val := range v {
			s := key

			// out[s], l, err = blah(val, l)
			switch va := val.(type) {
			case string, int, int32, int64, float32, float64, bool:

				if l == -1 || l == 1 {
					l = 1
				} else {
					return nil, fmt.Errorf("length of columns not consistent")
				}
				out[s] = []interface{}{va}
			default:
				cur, _ := ToInterfaceArray(va)
				lcur := len(cur)
				if l == -1 || l == lcur {
					l = lcur
				} else {
					return nil, fmt.Errorf("length of columns not consistent")
				}
				out[s] = cur

			}

			if err != nil {
				return nil, err
			}

		}
	default:
		return nil, fmt.Errorf("only slices/matrices/tensors and maps are accepted types")
	}

	return out, nil
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

type DataState interface {
	Update(newData interface{}) error
	Value() interface{}
}

func GetFunction(functionName string) DataState {
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

type First struct {
	gotValue bool
	data     interface{}
}

func (this *First) Value() interface{} {
	return this.data
}

func (this *First) Update(newData interface{}) error {
	if !this.gotValue {
		this.data = newData
		this.gotValue = true
	}
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

func ToInterfaceArray(val interface{}) ([]interface{}, error) {

	switch t := val.(type) {
	case []interface{}:
		return t, nil

	case []map[string]interface{}:
		var a []interface{}
		for _, v := range t {
			a = append(a, v)
		}
		return a, nil
	case nil:
		return nil, nil
	default:
		s := reflect.ValueOf(val)
		if s.Kind() == reflect.Slice {
			a := make([]interface{}, s.Len())
			for i := 0; i < s.Len(); i++ {
				element := s.Index(i).Interface()
				elementType := reflect.TypeOf(element)
				switch elementType.Kind() {
				case reflect.Slice:
					a[i], _ = ToInterfaceArray(element)
				case reflect.Array:
					a[i], _ = ToInterfaceArray(element)
				default:
					a[i] = element
				}
			}
			return a, nil
		}
		return nil, fmt.Errorf("unable to coerce %#v to []interface{}", val)
	}
}

type Callback func(tuple map[string]interface{}, newDataFrame *DataFrame, lastTuple bool) error

func ProcessDataFrame(dataFrame DataFrame, callback Callback) (result DataFrame, err error) {

	/* check tuple size */
	tupleSize := -1
	var count int
	for _, columnValues := range dataFrame {
		if 0 == count {
			tupleSize = len(columnValues)
		} else {
			if tupleSize != len(columnValues) {
				return nil, errors.New("Illegel dataframe : column value array with different size!")
			}
		}
	}

	newDataFrame := make(DataFrame)
	tuple := make(map[string]interface{})
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		for fieldname, filedsArray := range dataFrame {
			tuple[fieldname] = filedsArray[i]
		}
		err := callback(tuple, &newDataFrame, i == (tupleSize-1))
		if nil != err {
			return nil, err
		}
	}

	return newDataFrame, nil
}

func TupleArrayToDataframe(
	tuples []map[string]interface{},
	dataFrame *DataFrame) error {

	dataframeSize := len(tuples)
	if 0 == dataframeSize {
		return errors.New("Empty tuple array!")
	}

	for index, tuple := range tuples {
		for columnName, columnValue := range tuple {
			columnValueArray := (*dataFrame)[columnName]
			if nil == columnValueArray {
				columnValueArray = make([]interface{}, len(tuples))
				(*dataFrame)[columnName] = columnValueArray
			}

			(*dataFrame)[columnName][index] = columnValue
		}
	}

	return nil
}

/* fast but requires predefined dataframe size */
func TupleAssignToDataframe(
	index int,
	tuple map[string]interface{},
	dataFrame *DataFrame) error {
	for columnName, columnValue := range tuple {
		columnValueArray := (*dataFrame)[columnName]
		if nil == columnValueArray || index >= len(columnValueArray) {
			return errors.New("Index out of bound !")
		}
		(*dataFrame)[columnName][index] = columnValue
	}

	return nil
}

/* slow but flexible dataframe size */
func TupleAppendToDataframe(
	tuple map[string]interface{},
	dataFrame *DataFrame) error {
	dataframeSize := -1
	for columnName, columnValue := range tuple {
		columnValueArray := (*dataFrame)[columnName]
		if nil == columnValueArray {
			columnValueArray = make([]interface{}, 0)
			(*dataFrame)[columnName] = columnValueArray
		}

		if dataframeSize < 0 {
			dataframeSize = len(columnValueArray)
		} else if dataframeSize != len(columnValueArray) {
			return errors.New("Unequal column value array size !")
		}

		(*dataFrame)[columnName] = append(columnValueArray, columnValue)
	}

	return nil
}
