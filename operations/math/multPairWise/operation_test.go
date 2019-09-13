package multPairWise

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestSizeError(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["matrix0"] = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputs["matrix1"] = []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	opt.logger.Info("Error expected, got: %s", err)
	assert.Nil(t, output)
	assert.NotNil(t, err)

}

func TestArr(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["matrix0"] = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputs["matrix1"] = []interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	opt.logger.Info("Output expected, got:", output)
	assert.NotNil(t, output)
	// assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}

func TestMtx(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["matrix0"] = []interface{}{[]interface{}{1}, []interface{}{2}}
	inputs["matrix1"] = []interface{}{[]interface{}{10}, []interface{}{9}}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	opt.logger.Info("Output expected, got:", output)
	assert.NotNil(t, output)
	// assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}
