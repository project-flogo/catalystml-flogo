package subsectiontoimage

import (
	"fmt"
	"io/ioutil"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestBorder(t *testing.T) {
	t.Parallel()
	inputs := make(map[string]interface{})

	p := Params{Size: []int{2000, 2250}, LowerLeftCorner: []int{3500, 2300}}

	file := "../test_image.jpg"
	img, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	inputs["img"] = img

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestInMiddle(t *testing.T) {
	t.Parallel()
	inputs := make(map[string]interface{})

	p := Params{Size: []int{1500, 1750}, LowerLeftCorner: []int{500, 700}}

	file := "../test_image.jpg"
	img, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	inputs["img"] = img

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
