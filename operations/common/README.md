# Package common

import "github.com/project-flogo/catalystml-flogo/operations/common"

Package common is used as common library for operations

## Constant

DataFrameOrderLabel = "order"

## DataFrame

### func ToDataFrame
```
func ToDataFrame(data interface{}) (*DataFrame, error)
``` 
### func ProcessDataFrame
```
func ProcessDataFrame(dataFrame *DataFrame, callback Callback) error 
```
### type Callback
```
type Callback func(tuple *SortableTuple, lastTuple bool) error
```
### func Transpose
```
func Transpose(dataFrame *DataFrame, newLabels []string) *DataFrame 
```
### func TupleArrayToDataframe
```
func TupleArrayToDataframe(tuples []map[string]interface{}, dataFrame *DataFrame) error 
```
### func TupleArrayToDataframe
```
/* fast but requires predefined dataframe size */
func TupleAssignToDataframe(index int, tuple map[string]interface{}, dataFrame *DataFrame) error 
```
### func TupleArrayToDataframe
```
/* slow but flexible dataframe size */
func TupleAppendToDataframe(tuple map[string]interface{}, dataFrame *DataFrame) error 
```
### func TupleArrayToDataframe
```
/* slow but flexible dataframe size */
func SortableTupleAppendToDataframe(tuple SortableTuple, dataFrame *DataFrame) error 
```
### func NewDataFrame
```
func NewDataFrame() *DataFrame 
```
### type DataFrame
```
type DataFrame struct {
	fromTable bool
	order     []string
	data      map[string][]interface{}
}
func (dataFrame *DataFrame) GetFromTable() bool 
func (dataFrame *DataFrame) SetFromTable(fromTable bool) 
func (dataFrame *DataFrame) GetLabels() []string 
func (dataFrame *DataFrame) GetColumn(label string) []interface{} 
func (dataFrame *DataFrame) AsTable() map[string]interface{} 
func (dataFrame *DataFrame) AsMatrix() [][]interface{} 
func (dataFrame *DataFrame) AsIs() interface{} 
func (dataFrame *DataFrame) AddColumn(colName string, colValues []interface{}) 
```
## for sorting data frame
### func NewDataFrameSorter
```
func NewDataFrameSorter(Axis int, Ascending bool, NilLast bool, ByKey bool, SortBy []interface{}, dataFrame *DataFrame) *DataFrameSorter 
```
### type DataFrameSorter
```
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
func (s DataFrameSorter) GetDataFrame() *DataFrame 
func (s DataFrameSorter) Len() int 
func (s DataFrameSorter) Less(i, j int) bool 
func (s DataFrameSorter) Swap(i, j int) 
func (s DataFrameSorter) compare(valuei interface{}, valuej interface{}) int 
```
### func NewSortableTuple
```
func NewSortableTuple(data map[string]interface{}, fieldOrder []string) *SortableTuple 
```
### type SortableTuple
```
type SortableTuple struct {
	order []string
	Data  map[string]interface{}
}
func (t SortableTuple) GetData() map[string]interface{} 
func (t SortableTuple) GetDataArray() []interface{} 
func (t SortableTuple) GetByKey(key string) interface{} 
func (t SortableTuple) GetByIndex(index int) interface{}
```
## Aggregates
### func GetFunction
```
func GetFunction(functionName string) DataState 
```
## type DataState
```
type DataState interface
Update(newData interface{}) error
Value() interface{}
```
### type Data struct
```
type Data struct
data interface{}
func (this *Data) Value() interface{} 
func (this *Data) Update(newData interface{}) error 
```
### type First struct
```
type First struct
gotValue bool
data     interface{}
func (this *First) Value() interface{} 
func (this *First) Update(newData interface{}) error 
```
### type Sum struct
```
type Sum struct
data float64
func (this *Sum) Value() interface{} 
func (this *Sum) Update(newData interface{}) error 
```
### type Count struct
```
type Count struct
counter int
func (this *Count) Value() interface{} 
func (this *Count) Update(newData interface{}) error 
```
### type Mean
```
type Mean struct
sum   float64
count float64
func (this *Mean) Value() interface{} 
func (this *Mean) Update(newData interface{}) error 
```
### type Min
```
type Min struct 
min interface{}
func (this *Min) Value() interface{} 
func (this *Min) Update(newData interface{}) error 
```
### type Max
```
type Max struct 
max interface{}
func (this *Max) Value() interface{} 
func (this *Max) Update(newData interface{}) error 
```
## func ToInterfaceArray
```
func ToInterfaceArray(val interface{}) ([]interface{}, error) 
```
## func compare
```
func compare(data1 interface{}, data2 interface{}) (int, error) 
```
## type Index
```
type Index struct
Id uint64
func NewIndex(elements []interface{}) Index 
```