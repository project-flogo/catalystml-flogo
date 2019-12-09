# Package common

import "github.com/project-flogo/catalystml-flogo/operations/common"

Package common is used as common library for operations

## DataFrame

### func ToDataFrame
```
func ToDataFrame(data interface{}) (*DataFrame, error)
``` 
Convert array or map to data frame.
### func ProcessDataFrame
```
func ProcessDataFrame(dataFrame *DataFrame, callback Callback) error 
```
This function will iterate though all tuples in a dataframe and pass them to callback function.
### type Callback
```
type Callback func(tuple *SortableTuple, lastTuple bool) error
```
User implement it and pass it to ProcessDataFrame for processing dataframe tuple by tuple.
### func Transpose
```
func Transpose(dataFrame *DataFrame, newLabels []string) *DataFrame 
```
Transpose input dataframe and applying new labels to new data frame.
### func TupleArrayToDataframe
```
func TupleArrayToDataframe(tuples []map[string]interface{}, dataFrame *DataFrame) error 
```
Convert an tuple array to a data frame.
### func TupleArrayToDataframe
```
/* fast but requires predefined dataframe size */
func TupleAssignToDataframe(index int, tuple map[string]interface{}, dataFrame *DataFrame) error 
```
Add a tuple to a data frame.
### func TupleArrayToDataframe
```
/* slow but flexible dataframe size */
func TupleAppendToDataframe(tuple map[string]interface{}, dataFrame *DataFrame) error 
```
Add a tuple to a data frame.
### func TupleArrayToDataframe
```
/* slow but flexible dataframe size */
func SortableTupleAppendToDataframe(tuple SortableTuple, dataFrame *DataFrame) error 
```
Add a sortable tuple to a data frame.
### func NewDataFrame
```
func NewDataFrame() *DataFrame 
```
Construct a new data frame.
### type DataFrame
```
type DataFrame struct {
	// contains filtered or unexported fields
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
Construct a new data frame sorter.
### type DataFrameSorter
```
type DataFrameSorter struct {
	// contains filtered or unexported fields
}
func (s DataFrameSorter) GetDataFrame() *DataFrame 
func (s DataFrameSorter) Len() int 
func (s DataFrameSorter) Less(i, j int) bool 
func (s DataFrameSorter) Swap(i, j int) 
func (s DataFrameSorter) compare(valuei interface{}, valuej interface{}) int 
```
### func NewSortableTuple
```
Construct a new sortable tuple.
func NewSortableTuple(data map[string]interface{}, fieldOrder []string) *SortableTuple 
```
### type SortableTuple
```
type SortableTuple struct {
	// contains filtered or unexported fields
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
type DataState interface {
	Update(newData interface{}) error
	Value() interface{}
}
```
The interface defined for all aggregate types to implement. Call Update to perform aggregation and call Value to get aggregated data.
### type Data struct
```
type Data struct{
	// contains filtered or unexported fields
}
```
A simple data holder. Call Update will overwrite previous value.
```
func (this *Data) Value() interface{} 
func (this *Data) Update(newData interface{}) error 
```
### type First struct
```
type First struct{
	// contains filtered or unexported fields
}
```
A simple data holder. Call Update will accept the first none nil data.
```
func (this *First) Value() interface{} 
func (this *First) Update(newData interface{}) error 
```
### type Sum struct
```
type Sum struct{
	// contains filtered or unexported fields
}
```
Summation calculation. Call Update will convert input data to float64 then add it to current value (float64 type).
```
func (this *Sum) Value() interface{} 
func (this *Sum) Update(newData interface{}) error 
```
### type Count struct
```
type Count struct{
	// contains filtered or unexported fields
}
```
Call Update will increment current value by one (int type).
```
func (this *Count) Value() interface{} 
func (this *Count) Update(newData interface{}) error 
```
### type Mean
```
type Mean struct{
	// contains filtered or unexported fields
}
```
Mean calculation. Call Update will convert input data to float64 then replace current value with the average of all input data.
```
func (this *Mean) Value() interface{} 
func (this *Mean) Update(newData interface{}) error 
```
### type Min
```
type Min struct {
	// contains filtered or unexported fields
}
```
Accept float64, int and string type of input. Call Update to replace the current value with input value if it is smaller than current value.
```
func (this *Min) Value() interface{} 
func (this *Min) Update(newData interface{}) error 
```
### type Max
```
type Max struct {
	// contains filtered or unexported fields
}
```
Accept float64, int and string type of input. Call Update to replace the current value with input value if it is bigger than current value.
```
func (this *Max) Value() interface{} 
func (this *Max) Update(newData interface{}) error 
```
## func ToInterfaceArray
```
func ToInterfaceArray(val interface{}) ([]interface{}, error) 
```
Coerce array, slice or map to interface array. Return nil array and error if the conversion is impossible.
