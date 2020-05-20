package resize

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestXOnly(t *testing.T) {
	t.Parallel()

	inputs := make(map[string]interface{})

	p := Params{Xsize: 100, Ysize: 0, Algo: "Linear"}

	file := "../test_image.jpg"
	img, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
		return
	}

	inputs["img"] = img

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	// fmt.Println(output)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestYOnly(t *testing.T) {
	t.Parallel()

	inputs := make(map[string]interface{})

	p := Params{Ysize: 150, Xsize: 0, Algo: "Lanczos"}

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
	// fmt.Println(output)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestBoth(t *testing.T) {
	t.Parallel()

	inputs := make(map[string]interface{})

	p := Params{Ysize: 150, Xsize: 100, Algo: "CatmullRom"}

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
	// fmt.Println(output)
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
