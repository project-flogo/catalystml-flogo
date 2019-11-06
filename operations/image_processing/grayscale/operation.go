package grayscale

import (
	"bytes"
	"github.com/project-flogo/catalystml-flogo/action/operation"

	"image"
	"image/color"
	"image/jpeg"

	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}

	input.FromMap(inputs)
	img := input.Img.(image.Image)

	a.logger.Info("Starting operation grayscale.")
	a.logger.Info("Inputs for Operation grayscale.", img.At(0, 0))

	bnds := img.Bounds()
	newimg := image.NewRGBA(bnds)
	for x := 0; x < bnds.Max.X; x++ {
		for y := 0; y < bnds.Max.Y; y++ {
			newimg.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, newimg, nil)
	if err != nil {
		return nil, err
	}
	a.logger.Info("Operation grayscale Completed.")

	return buf, nil
}
