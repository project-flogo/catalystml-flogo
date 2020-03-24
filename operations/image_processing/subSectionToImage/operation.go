package subsectiontoimage

import (
	// "bytes"

	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	wsub := a.params.Size[0]
	hsub := a.params.Size[1]
	xsub := a.params.LowerLeftCorner[0]
	ysub := a.params.LowerLeftCorner[1]

	// a.logger.Debug("inputs", inputs)
	a.logger.Info("Starting Operation subsectiontoimage.")
	a.logger.Debugf("The subection is %dx%d in size with a lower left corner at (%d,%d)", wsub, hsub, xsub, ysub)

	src := input.Img.(image.Image)

	bnds := src.Bounds()

	x0sub := xsub
	x1sub := xsub + wsub
	y0sub := ysub
	y1sub := ysub + hsub
	if y1sub >= bnds.Max.Y {
		y1sub = bnds.Max.Y
	}
	if x1sub >= bnds.Max.X {
		x1sub = bnds.Max.X
	}

	newbnds := image.Rect(0, 0, x1sub-x0sub, y1sub-y0sub)
	newimg := image.NewRGBA(newbnds)
	for x := x0sub; x < x1sub; x++ {
		for y := x0sub; y < y1sub; y++ {
			newimg.Set(x-x0sub, y-y0sub, color.RGBAModel.Convert(src.At(x, y)))
		}
	}

	out, err := os.Create("./output.png")
	if err != nil {
		return nil, err
	}
	defer out.Close()

	err = png.Encode(out, newimg)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newimg, nil)
	if err != nil {
		return nil, err
	}

	a.logger.Info("Operation img2tensor Completed.")
	a.logger.Debug("Values of first pixel ", newimg.At(0, 0))
	return buf, nil

}
