package grayscale

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestXOnly(t *testing.T) {
	t.Parallel()
	inputs := make(map[string]interface{})
	opt := &Operation{logger: log.RootLogger()}

	file := "../test_image.jpg"
	img, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	inputs["img"] = img
	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
