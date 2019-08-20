package resize

import (
	"fmt"
	"image"
	"os"
	"testing"

	"github.com/project-flogo/cml/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestXOnly(t *testing.T) {

	inputs := make(map[string]interface{})

	p := Params{Xsize: 100, Ysize: 0, Algo: "Linear"}

	file := "/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/dataset2/Cup/Image6.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error Decoding file: %v\n", err)
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

	inputs := make(map[string]interface{})

	p := Params{Ysize: 150, Xsize: 0, Algo: "Lanczos"}

	file := "/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/dataset2/Cup/Image6.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error Decoding file: %v\n", err)
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

	inputs := make(map[string]interface{})

	p := Params{Ysize: 150, Xsize: 100, Algo: "CatmullRom"}

	file := "/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/dataset2/Cup/Image6.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Oopening file: %v\n", err)
		return
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error Decoding file: %v\n", err)
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
