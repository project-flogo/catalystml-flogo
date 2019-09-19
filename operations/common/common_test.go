package common

import (
	"fmt"
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