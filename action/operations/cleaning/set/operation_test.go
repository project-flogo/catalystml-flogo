package set

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["arr"] = []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4}
	// inputs["arr"] = []string{"a", "abr", "b", "abr", "a"}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	fmt.Println(output)
	assert.NotNil(t, output)
	assert.ElementsMatch(t, output, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}, "two should be the same")
	assert.Nil(t, err)

}
