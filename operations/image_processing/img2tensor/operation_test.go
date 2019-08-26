package img2tensor

import (
	"fmt"
	"image"

	// _ "image/gif"
	// _ "image/jpeg"
	_ "image/png"
	"os"
	"testing"

	"github.com/project-flogo/cml/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestFalse(t *testing.T) {

	inputs := make(map[string]interface{})

	p := Params{RemoveAlpha: false, IncludeBatch: false}

	file := "../test_image.jpg"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
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
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestTrue(t *testing.T) {

	inputs := make(map[string]interface{})

	p := Params{RemoveAlpha: true, IncludeBatch: true}

	file := "/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/dataset2/Cup/Image6.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
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
	assert.NotNil(t, output)
	assert.Nil(t, err)

}

func TestDefaults(t *testing.T) {

	inputs := make(map[string]interface{})

	p := Params{}

	file := "/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/dataset2/Cup/Image6.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
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
	assert.NotNil(t, output)
	assert.Nil(t, err)

}
