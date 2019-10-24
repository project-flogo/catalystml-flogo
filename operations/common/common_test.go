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
	newDataFrame, _ := ProcessDataFrame(*dataframe, func(tuple map[string]interface{}, newDataFrame *DataFrame, lastTuple bool) error {
		newTuple["sum"] = newTuple["sum"].(int) + tuple["blah"].(int)
		newTuple["count"] = newTuple["count"].(int) + 1
		if lastTuple {
			TupleAppendToDataframe(newTuple, newDataFrame)
		}
		return nil
	})

	fmt.Println(newDataFrame)
}

func TestSortableTuple(t *testing.T) {
	tuples := TupleSorter{
		ByKey:  true,
		SortBy: []interface{}{"col1", "col2", "col3", "col4"},
		Tuples: make([]SortableTuple, 4),
	}

	keyToIndex := map[string]int{
		"col1": 0,
		"col2": 1,
		"col3": 2,
		"col4": 3,
	}

	tuples.Tuples[0] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3718,
			2138.0,
			1908,
			912},
	}
	tuples.Tuples[1] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3711,
			2138.0,
			1908,
			912},
	}
	tuples.Tuples[2] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3718,
			2138.0,
			1940,
			970},
	}
	tuples.Tuples[3] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3703,
			2125.0,
			1933,
			943},
	}

	fmt.Println("Before : ", tuples)
	sort.Sort(tuples)
	fmt.Println("After  : ", tuples)
}

func TestSortableTupleByIndex(t *testing.T) {
	tuples := TupleSorter{
		ByKey:  false,
		SortBy: []interface{}{0, 1, 2, 3},
		Tuples: make([]SortableTuple, 4),
	}

	keyToIndex := map[string]int{
		"col1": 0,
		"col2": 1,
		"col3": 2,
		"col4": 3,
	}

	tuples.Tuples[0] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3718,
			2138.0,
			1908,
			912},
	}
	tuples.Tuples[1] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3711,
			2138.0,
			1908,
			912},
	}
	tuples.Tuples[2] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3718,
			2138.0,
			1940,
			970},
	}
	tuples.Tuples[3] = SortableTuple{
		KeyToIndex: keyToIndex,
		Data: []interface{}{
			3703,
			2125.0,
			1933,
			943},
	}

	fmt.Println("Before : ", tuples)
	sort.Sort(tuples)
	fmt.Println("After  : ", tuples)
}
