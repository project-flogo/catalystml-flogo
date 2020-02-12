package tensor2image

import (
	"fmt"
	"image"

	// _ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func Test2PNG(t *testing.T) {

	var err error
	inputs := make(map[string]interface{})
	inputs["tensor"], err = getImageAsSlice(t, "../test_image.jpg", true, false)

	p := Params{ToFileFolder: "../", Filename: "test_image_out", Extension: "PNG"}

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)
}

func Test2JPG(t *testing.T) {

	var err error
	inputs := make(map[string]interface{})
	inputs["tensor"], err = getImageAsSlice(t, "../test_image.jpg", true, false)

	p := Params{ToFileFolder: "../", Filename: "test_image_out", Extension: "jpg"}

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)
}

func Test2GIF(t *testing.T) {

	var err error
	inputs := make(map[string]interface{})
	inputs["tensor"], err = getImageAsSlice(t, "../test_image.jpg", true, false)

	p := Params{ToFileFolder: "../", Filename: "test_image_out", Extension: "gif"}

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	assert.NotNil(t, output)
	assert.Nil(t, err)
}

func getImageAsSlice(t *testing.T, file string, rmAlpha bool, includeBatch bool) (interface{}, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
		return nil, err
	}

	src, str, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error Decoding file: %v\n", err)
		return nil, err
	}
	t.Log("image file format : ", str)

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	pixsize := 4 // size of the pixel
	if rmAlpha {
		pixsize = 3
	}

	// //Converting Image to array

	var singleimg [][][]uint8
	for x := 0; x < w; x++ {
		var row [][]uint8
		for y := 0; y < h; y++ {
			var col []uint8
			for i := 0; i < pixsize; i++ {
				col = append(col, 0)
			}

			row = append(row, col)
		}
		singleimg = append(singleimg, row)
	}

	scale := uint32(257) // RGBA scale factor
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := src.At(x, y)
			rr, bb, gg, aa := imageColor.RGBA()
			color := []uint8{uint8(rr / scale), uint8(bb / scale), uint8(gg / scale), uint8(aa)}
			for i := 0; i < pixsize; i++ {
				singleimg[x][y][i] = color[i]
			}

		}
	}

	if includeBatch {
		batchsize := 1
		var img [][][][]uint8
		for j := 0; j < batchsize; j++ {
			img = append(img, singleimg)
		}
		t.Log("Operation img2tensor Completed.")
		t.Log("Values of first pixel ", img[0][0][0])
		return img, nil

	}
	t.Log("Operation img2tensor Completed.")
	t.Log("Values of first pixel ", singleimg[0][0])
	return singleimg, nil

}
