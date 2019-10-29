package common

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"strconv"

	"fmt"
	"reflect"

	"github.com/project-flogo/core/data/coerce"
)

const (
	DataFrameOrderLabel = "order"
)

type DataFrame struct {
	fromTable bool
	order     []string
	data      map[string][]interface{}
}

func (dataFrame *DataFrame) GetFromTable() bool {
	return dataFrame.fromTable
}

func (dataFrame *DataFrame) GetLabels() []string {
	return dataFrame.order
}

func (dataFrame *DataFrame) GetColumn(label string) []interface{} {
	return dataFrame.data[label]
}

func (dataFrame *DataFrame) AsTable() map[string]interface{} {
	table := make(map[string]interface{})
	if 0 != len(dataFrame.order) {
		table[DataFrameOrderLabel] = dataFrame.order
	}

	for key, value := range dataFrame.data {
		table[key] = value
	}
	return table
}

func (dataFrame *DataFrame) AsMatrix() [][]interface{} {
	length := len(dataFrame.data)
	matrix := make([][]interface{}, 0)

	ProcessDataFrame(dataFrame, func(tuple *SortableTuple, lastTuple bool) error {
		array := make([]interface{}, length)
		for fieldIndex, fieldName := range tuple.order {
			array[fieldIndex] = tuple.Data[fieldName]
		}
		matrix = append(matrix, array)

		return nil
	})
	return matrix
}

func (dataFrame *DataFrame) AsIs() interface{} {
	if dataFrame.fromTable {
		return dataFrame.AsTable()
	} else {
		return dataFrame.AsMatrix()
	}
}

func (dataFrame *DataFrame) AddColumn(colName string, colValues []interface{}) {
	dataFrame.data[colName] = colValues
	dataFrame.order = append(dataFrame.order, colName)
}

func NewDataFrame() *DataFrame {
	return &DataFrame{
		order: make([]string, 0),
		data:  make(map[string][]interface{}),
	}
}

func ToDataFrame(data interface{}) (*DataFrame, error) {
	// This function takes an slices, matrix, tensors, and maps to a dataframe, does not handle other type caases

	out := NewDataFrame()

	var err error

	switch v := data.(type) {
	case []interface{}, [][]interface{}, [][][]interface{}, [][][][]interface{}, [][][][][]interface{},
		[]int, [][]int, [][][]int, [][][][]int, [][][][][]int,
		[]float64, [][]float64, [][][]float64, [][][][]float64, [][][][][]float64,
		[]string, [][]string, [][][]string, [][][][]string, [][][][][]string:

		out.fromTable = false

		//Test dimensionality of matrix
		//fmt.Println("data in is a slice", v)
		vcon, err := ToInterfaceArray(v)
		if err != nil {
			return nil, err
		}

		e := 0
		for i, val := range vcon {
			s := fmt.Sprintf("%d", i)
			switch w := val.(type) {
			case string, int, int32, int64, float32, float64, bool:
				out.AddColumn(s, []interface{}{w})
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
			out.AddColumn(s, arr)
		}

	case map[string]interface{}:

		out.fromTable = true

		colOrder, ok := v["order"].([]interface{})
		if !ok || 0 == len(colOrder) {
			colOrder = make([]interface{}, 0)
			for key, _ := range v {
				if "order" == key {
					continue
				}
				colOrder = append(colOrder, key)
			}
		}

		l := -1
		for _, key := range colOrder {
			s := key.(string)
			val := v[key.(string)]

			// out[s], l, err = blah(val, l)
			switch va := val.(type) {
			case string, int, int32, int64, float32, float64, bool:

				if l == -1 || l == 1 {
					l = 1
				} else {
					return nil, fmt.Errorf("length of columns not consistent")
				}
				out.AddColumn(s, []interface{}{va})
			default:
				cur, _ := ToInterfaceArray(va)
				lcur := len(cur)
				if l == -1 || l == lcur {
					l = lcur
				} else {
					return nil, fmt.Errorf("length of columns not consistent")
				}
				out.AddColumn(s, cur)

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
	case string:
		delta1str, _ := coerce.ToString(data1)
		delta2str, err := coerce.ToString(data2)
		if nil != err {
			return 0, err
		}
		switch {
		case delta1str > delta2str:
			return 1, nil
		case delta1str == delta2str:
			return 0, nil
		case delta1str < delta2str:
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

type Callback func(tuple *SortableTuple, lastTuple bool) error

func ProcessDataFrame(dataFrame *DataFrame, callback Callback) error {

	/* check tuple size */
	tupleSize := -1
	var count int
	for _, columnValues := range dataFrame.data {
		if 0 == count {
			tupleSize = len(columnValues)
		} else {
			if tupleSize != len(columnValues) {
				return errors.New("Illegel dataframe : column value array with different size!")
			}
		}
	}

	tuple := make(map[string]interface{})
	for i := 0; i < tupleSize; i++ {
		/* build tuple */
		for fieldname, filedsArray := range dataFrame.data {
			tuple[fieldname] = filedsArray[i]
		}
		err := callback(NewSortableTuple(tuple, dataFrame.GetLabels()), i == (tupleSize-1))
		if nil != err {
			return err
		}
	}

	return nil
}

func Transpose(dataFrame *DataFrame, newLabels []string) *DataFrame {
	table := make(map[string]interface{})
	var newDataFrame *DataFrame
	index := 0
	ProcessDataFrame(dataFrame, func(tuple *SortableTuple, lastTuple bool) error {
		var label string
		if nil == newLabels || 0 == len(newLabels) {
			label = strconv.Itoa(index)
		} else {
			label = newLabels[index]
		}

		table[label] = tuple.GetDataArray()
		if lastTuple {
			newDataFrame, _ = ToDataFrame(table)
			fmt.Println(newDataFrame)
		}
		index++
		return nil
	})
	return newDataFrame
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
			columnValueArray := (*dataFrame).data[columnName]
			if nil == columnValueArray {
				columnValueArray = make([]interface{}, len(tuples))
				(*dataFrame).data[columnName] = columnValueArray
			}

			(*dataFrame).data[columnName][index] = columnValue
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
		columnValueArray := (*dataFrame).data[columnName]
		if nil == columnValueArray || index >= len(columnValueArray) {
			return errors.New("Index out of bound !")
		}
		(*dataFrame).data[columnName][index] = columnValue
	}

	return nil
}

/* slow but flexible dataframe size */
func TupleAppendToDataframe(
	tuple map[string]interface{},
	dataFrame *DataFrame) error {
	dataframeSize := -1
	for columnName, columnValue := range tuple {
		columnValueArray := (*dataFrame).data[columnName]
		if nil == columnValueArray {
			columnValueArray = make([]interface{}, 0)
			(*dataFrame).data[columnName] = columnValueArray
		}

		if dataframeSize < 0 {
			dataframeSize = len(columnValueArray)
		} else if dataframeSize != len(columnValueArray) {
			return errors.New("Unequal column value array size !")
		}

		(*dataFrame).data[columnName] = append(columnValueArray, columnValue)
	}

	return nil
}

/* slow but flexible dataframe size */
func SortableTupleAppendToDataframe(
	tuple SortableTuple,
	dataFrame *DataFrame) error {
	dataframeSize := -1
	fmt.Println(tuple)
	for _, columnName := range tuple.order {
		columnValueArray := (*dataFrame).data[columnName]
		if nil == columnValueArray {
			columnValueArray = make([]interface{}, 0)
			(*dataFrame).data[columnName] = columnValueArray
		}

		if dataframeSize < 0 {
			dataframeSize = len(columnValueArray)
		} else if dataframeSize != len(columnValueArray) {
			return errors.New("Unequal column value array size !")
		}

		(*dataFrame).data[columnName] = append(columnValueArray, tuple.Data[columnName])
	}

	return nil
}

func NewSortableTuple(data map[string]interface{}, fieldOrder []string) *SortableTuple {
	sTuple := &SortableTuple{
		order: make([]string, len(data)),
		Data:  make(map[string]interface{}),
	}

	if nil == fieldOrder || 0 == len(fieldOrder) {
		//fmt.Println("no predefined order : ", fieldOrder)
		index := 0
		for key, value := range data {
			sTuple.order[index] = key
			sTuple.Data[key] = value
			index++
		}
	} else {
		//fmt.Println("has predefined order : ", fieldOrder)
		for index, key := range fieldOrder {
			sTuple.order[index] = key
			sTuple.Data[key] = data[key]
			index++
		}
	}

	return sTuple
}

type SortableTuple struct {
	order []string
	Data  map[string]interface{}
}

func (t SortableTuple) GetData() map[string]interface{} {
	return t.Data
}

func (t SortableTuple) GetDataArray() []interface{} {
	dataArray := make([]interface{}, len(t.Data))
	for index, key := range t.order {
		dataArray[index] = t.Data[key]
	}
	return dataArray
}

func (t SortableTuple) GetByKey(key string) interface{} {
	//	fmt.Println("Key = ", key, ", Data = ", t.Data)
	return t.Data[key]
}

func (t SortableTuple) GetByIndex(index int) interface{} {
	return t.Data[t.order[index]]
}

func NewDataFrameSorter(
	Axis int,
	Ascending bool,
	NilLast bool,
	ByKey bool,
	SortBy []interface{},
	dataFrame *DataFrame,
) *DataFrameSorter {
	sorter := &DataFrameSorter{
		Axis:              Axis,
		Ascending:         Ascending,
		NilLast:           NilLast,
		ByKey:             ByKey,
		SortBy:            SortBy,
		RowLabels:         make([]string, 0),
		ColumnLabels:      dataFrame.order,
		Tuples:            make([]SortableTuple, 0),
		OriginalFromTable: dataFrame.fromTable,
	}

	switch Axis {
	case 0:
		index := 0
		ProcessDataFrame(dataFrame, func(tuple *SortableTuple, lastTuple bool) error {
			sorter.Tuples = append(sorter.Tuples, *tuple)
			sorter.RowLabels = append(sorter.RowLabels, strconv.Itoa(index))
			index++
			return nil
		})
	case 1:
		for index, label := range dataFrame.GetLabels() {
			column := dataFrame.GetColumn(label)
			tuples := make(map[string]interface{})
			fieldOrder := make([]string, len(column))
			for rIndex, value := range column {
				tuples[strconv.Itoa(rIndex)] = value
				fieldOrder[rIndex] = strconv.Itoa(rIndex)
			}
			sorter.Tuples = append(sorter.Tuples, *NewSortableTuple(tuples, fieldOrder))
			sorter.RowLabels = append(sorter.RowLabels, strconv.Itoa(index))
		}
	}

	return sorter
}

type DataFrameSorter struct {
	Axis              int
	Ascending         bool
	NilLast           bool
	ByKey             bool
	SortBy            []interface{}
	Tuples            []SortableTuple
	ColumnLabels      []string
	RowLabels         []string
	OriginalFromTable bool
}

func (s DataFrameSorter) GetDataFrame() *DataFrame {
	var dataFrame *DataFrame
	table := make(map[string]interface{})
	switch s.Axis {
	case 0:
		for index, sTuple := range s.Tuples {
			for _, label := range s.ColumnLabels {
				columns := table[label]
				if nil == columns {
					columns = make([]interface{}, len(s.RowLabels))
					table[label] = columns
				}
				columns.([]interface{})[index] = sTuple.GetByKey(label)
			}
		}

		table[DataFrameOrderLabel] = make([]interface{}, len(s.ColumnLabels))
		for index, label := range s.ColumnLabels {
			table[DataFrameOrderLabel].([]interface{})[index] = label
		}

	case 1:
		table[DataFrameOrderLabel] = make([]interface{}, len(s.ColumnLabels))

		for index, RowLabel := range s.RowLabels {
			indexRowLabel, _ := strconv.Atoi(RowLabel)
			table[s.ColumnLabels[index]] = s.Tuples[indexRowLabel].GetDataArray()
			table[DataFrameOrderLabel].([]interface{})[index] = s.ColumnLabels[indexRowLabel]
		}
	}
	var err error
	dataFrame, err = ToDataFrame(table)
	if nil != err {
		fmt.Println(err.Error())
	}
	dataFrame.fromTable = s.OriginalFromTable

	return dataFrame
}

func (s DataFrameSorter) Len() int {
	return len(s.Tuples)
}

func (s DataFrameSorter) Less(i, j int) bool {
	for _, sortKey := range s.SortBy {
		var result int
		var valuei interface{}
		var valuej interface{}
		if s.ByKey {
			valuei = s.Tuples[i].GetByKey(sortKey.(string))
			valuej = s.Tuples[j].GetByKey(sortKey.(string))
		} else {
			valuei = s.Tuples[i].GetByIndex(sortKey.(int))
			valuej = s.Tuples[j].GetByIndex(sortKey.(int))
		}
		result = s.compare(valuei, valuej)

		if 0 == result {
			continue
		} else {
			if s.Ascending {
				return 0 > result
			} else {
				return 0 < result
			}
		}
	}

	return true
}

func (s DataFrameSorter) Swap(i, j int) {
	s.Tuples[i], s.Tuples[j] = s.Tuples[j], s.Tuples[i]
	s.RowLabels[i], s.RowLabels[j] = s.RowLabels[j], s.RowLabels[i]
}

func (s DataFrameSorter) compare(valuei interface{}, valuej interface{}) int {
	var result int
	if nil == valuei && nil == valuej {
		return 0
	} else if nil == valuei {
		if s.Ascending == s.NilLast {
			result = 1
		} else if s.Ascending != s.NilLast {
			result = -1
		}
	} else if nil == valuej {
		if s.Ascending == s.NilLast {
			result = -1
		} else if s.Ascending != s.NilLast {
			result = 1
		}
	} else {
		result, _ = compare(valuei, valuej)
	}
	return result
}
