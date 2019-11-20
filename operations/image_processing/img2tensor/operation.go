package img2tensor

import (
	// "bytes"

	"fmt"
	"image"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/mapper"
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
	// removeAlpha and includebatch both default to flase which is the
	//     zero value of a bool

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

type initContextImpl struct {
	params   map[string]interface{}
	mFactory mapper.Factory
	name     string
}

func (ctx *initContextImpl) Params() map[string]interface{} {
	return ctx.params
}

func (ctx *initContextImpl) MapperFactory() mapper.Factory {
	return ctx.mFactory
}

func (ctx *initContextImpl) Logger() log.Logger {
	return log.ChildLogger(log.RootLogger(), ctx.name)
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	rmAlpha := a.params.RemoveAlpha

	// a.logger.Debug("inputs", inputs)
	a.logger.Info("Starting Operation img2tensor.")
	a.logger.Debug("RemoveAlpha is set to ...", rmAlpha)
	a.logger.Debug("IncludeBatch is set to ...", a.params.IncludeBatch)

	src := input.Img.(image.Image)

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

	if a.params.IncludeBatch {
		batchsize := 1
		var img [][][][]uint8
		for j := 0; j < batchsize; j++ {
			img = append(img, singleimg)
		}
		a.logger.Info("Operation img2tensor Completed.")
		a.logger.Debug("Values of first pixel ", img[0][0][0])
		fmt.Println(img[0][0])
		return img, nil

	}
	a.logger.Info("Operation img2tensor Completed.")
	a.logger.Debug("Values of first pixel ", singleimg[0][0])
	fmt.Println(singleimg[0])
	return singleimg, nil

}
