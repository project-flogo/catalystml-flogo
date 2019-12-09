package add

import (
	"testing"

	"github.com/project-flogo/core/support/log"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["num0"] = 5.9
	inputs["num1"] = 3.3

	out, err := opt.Eval(inputs)

	log.RootLogger().Info("output of test is:", out)
	assert.Nil(t, err)

}

func TestInt(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["num0"] = 10
	inputs["num1"] = 3

	out, err := opt.Eval(inputs)

	log.RootLogger().Info("output of test is:", out)
	assert.Nil(t, err)

}

func TestDiv0l(t *testing.T) {
	opt := &Operation{logger: log.RootLogger()}

	inputs := make(map[string]interface{})
	inputs["num0"] = 5.9
	inputs["num1"] = 1

	out, err := opt.Eval(inputs)
	log.RootLogger().Info("output of test is:", out)
	assert.Nil(t, err)

}
