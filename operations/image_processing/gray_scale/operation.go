package gray_scale

import (
	"image"
	_ "image/jpeg"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = operation.Register(&Operation{}, New)
}

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	// Decode image to JPEG
	img, _, err := image.Decode()
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	log.Printf("Image type: %T", img)

	// Converting image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return grayImg, nil
}
