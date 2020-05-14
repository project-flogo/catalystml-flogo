package valToArray

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"

	"github.com/stretchr/testify/assert"
)

func TestOnel(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})

	inputs["value"] = 7.8
	inputs["shape"] = []int{1, 2, 3}

	out, err := opt.Eval(inputs)
	fmt.Println(out)
	assert.Nil(t, err)
}

func TestTwo(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})

	inputs["value"] = 7.8
	inputs["shape"] = []int{3, 2, 1}

	out, err := opt.Eval(inputs)
	fmt.Println(out)
	assert.Nil(t, err)
}

func TestThree(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})

	inputs["value"] = 5
	inputs["shape"] = []int{3, 3, 3}

	out, err := opt.Eval(inputs)
	fmt.Println(out)
	assert.Nil(t, err)
}
