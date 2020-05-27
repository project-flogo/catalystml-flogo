package cmlmapper

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data"
	"github.com/stretchr/testify/assert"
)

func TestMapping(t *testing.T) {
	deref := NewExpression("$math1[2][1]")

	inputValues := make(map[string]interface{})

	inputValues["math1"] = [][]interface{}{{1}, {2, 6}, {3, 4, 5}}

	scope := data.NewSimpleScope(inputValues, nil)
	result, _ := Resolve(deref, scope)
	fmt.Println("Result..", result)
}

func TestMapping2(t *testing.T) {
	deref := NewExpression("$math1[2]")

	inputValues := make(map[string]interface{})

	inputValues["math1"] = []interface{}{1, 2, 3, 4, 5}

	scope := data.NewSimpleScope(inputValues, nil)
	result, _ := Resolve(deref, scope)
	fmt.Println("Result..", result)
}

func TestMapping3(t *testing.T) {
	deref := NewExpression("$math1['col']")

	inputValues := make(map[string]interface{})

	inputValues["math1"] = map[string]interface{}{"col": "sample"}

	scope := data.NewSimpleScope(inputValues, nil)
	result, err := Resolve(deref, scope)
	assert.Nil(t, err)
	assert.Equal(t, result, "sample")
	fmt.Println("Result..", result)
}

func TestMapping4(t *testing.T) {
	deref := NewExpression("$math1['col'][1]")

	inputValues := make(map[string]interface{})
	temp := []int{2, 3}
	inputValues["math1"] = map[string]interface{}{"col": temp}

	scope := data.NewSimpleScope(inputValues, nil)
	result, err := Resolve(deref, scope)
	assert.Nil(t, err)
	assert.Equal(t, result, 3)
	fmt.Println("Result..", result)

}
