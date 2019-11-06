package common

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMap(t *testing.T) {
	mapin := make(map[string]interface{})
	mapin["test"] = []string{"a", "b", "c"}
	mapin["blah"] = []int{1, 2, 3}

	out, err := ToDataFrame(mapin)
	// fmt.Println(err)
	fmt.Println(out, err)

	assert.Nil(t, err)

}

func TestSimpleMap2(t *testing.T) {
	mapin := make(map[string]interface{})
	mapin["test"] = "a"
	mapin["blah"] = []interface{}{"a"}

	out, err := ToDataFrame(mapin)
	// fmt.Println(err)
	fmt.Println(out, err)

	assert.Nil(t, err)

}

func TestSimpleMatrix(t *testing.T) {
	matrixin := [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}}

	out, err := ToDataFrame(matrixin)
	// fmt.Println(err)
	fmt.Println(out, err)

	assert.Nil(t, err)

}

func TestSimpleMatrix2(t *testing.T) {
	matrixin := []interface{}{1, 2}

	out, err := ToDataFrame(matrixin)
	// fmt.Println(err)
	fmt.Println(out, err)

	assert.Nil(t, err)

}

func TestSimpleMatrix3(t *testing.T) {
	matrixin := []interface{}{[]interface{}{[]interface{}{1}, []interface{}{2}}, []interface{}{[]interface{}{3}, []interface{}{4}}, []interface{}{[]interface{}{5}, []interface{}{6}}}

	out, err := ToDataFrame(matrixin)
	// fmt.Println(err)
	fmt.Println(out, err)

	assert.Nil(t, err)

}

func TestProcessDataFrame(t *testing.T) {
	mapin := make(map[string]interface{})
	mapin["test"] = []string{"a", "b", "c"}
	mapin["blah"] = []int{1, 2, 3}

	dataframe, _ := ToDataFrame(mapin)

	newTuple := make(map[string]interface{})
	newTuple["sum"] = 0
	newTuple["count"] = 0
	newDataFrame := NewDataFrame()
	ProcessDataFrame(dataframe, func(tuple *SortableTuple, lastTuple bool) error {
		newTuple["sum"] = newTuple["sum"].(int) + tuple.GetByKey("blah").(int)
		newTuple["count"] = newTuple["count"].(int) + 1
		if lastTuple {
			TupleAppendToDataframe(newTuple, newDataFrame)
		}
		return nil
	})

	fmt.Println(newDataFrame)
}

func TestSortableTuple(t *testing.T) {

	table := make(map[string]interface{})

	table[DataFrameOrderLabel] = []interface{}{
		"col1", "col2", "col3", "col4",
	}
	table["col1"] = []interface{}{
		3718, 3711, 3718, 3703,
	}
	table["col2"] = []interface{}{
		2138.0, 2138.0, 2138.0, 2125.0,
	}
	table["col3"] = []interface{}{
		1908, 1908, 1940, 1933,
	}
	table["col4"] = []interface{}{
		912, 912, 970, 943,
	}

	dataFrame, _ := ToDataFrame(table)

	tuples := NewDataFrameSorter(
		0,
		true,
		true,
		true,
		[]interface{}{"col1", "col2", "col3", "col4"},
		dataFrame,
	)

	fmt.Println("Before dataFrame : ", tuples.GetDataFrame())
	sort.Sort(tuples)
	fmt.Println("After dataFrame : ", tuples.GetDataFrame())
}

func TestSortableTupleByIndex(t *testing.T) {

	table := make(map[string]interface{})

	table[DataFrameOrderLabel] = []interface{}{
		"col1", "col2", "col3", "col4",
	}
	table["col1"] = []interface{}{
		3718, 3711, 3718, 3703,
	}
	table["col2"] = []interface{}{
		2138.0, 2138.0, 2138.0, 2125.0,
	}
	table["col3"] = []interface{}{
		1908, 1908, 1940, 1933,
	}
	table["col4"] = []interface{}{
		912, 912, 970, 943,
	}

	dataFrame, _ := ToDataFrame(table)

	tuples := NewDataFrameSorter(
		0,
		true,
		true,
		false,
		[]interface{}{0, 1, 2, 3},
		dataFrame,
	)

	fmt.Println("Before dataFrame : ", tuples.GetDataFrame())
	sort.Sort(tuples)
	fmt.Println("After dataFrame : ", tuples.GetDataFrame())

}
