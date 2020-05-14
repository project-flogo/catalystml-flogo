package resize

import (
	// "bytes"

	"image"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"

	"github.com/disintegration/imaging"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)
	if p.Algo == "" {
		p.Algo = "Lanczos"
	}

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}


func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}

	algo := a.params.Algo
	xsize := a.params.Xsize
	ysize := a.params.Ysize
	input.FromMap(inputs)

	a.logger.Info("Starting operation resize.")
	a.logger.Debug("Inputs for Operation resize.", inputs)
	a.logger.Info("Resampling Filter is.", algo)

	img := input.Img // image.Image type

	var rFilter imaging.ResampleFilter
	if algo == "Lanczos" {
		rFilter = imaging.Lanczos
	} else if algo == "NearestNeighbor" {
		rFilter = imaging.NearestNeighbor
	} else if algo == "Linear" {
		rFilter = imaging.Linear
	} else if algo == "CatmullRom" {
		rFilter = imaging.CatmullRom
	} else {
		rFilter = imaging.Lanczos
	}

	src := img.(image.Image)
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	if xsize <= 0 && ysize <= 0 {
		a.logger.Infof("WARNING: no resizing done due to lack of dimensions to resize to")

	} else if xsize > 0 && ysize > 0 {
		w = xsize
		h = ysize
		a.logger.Infof("Resizing to the x and y values given: %dx%d", w, h)
	} else if xsize > 0 && ysize <= 0 {
		w = xsize
		h = int(w * bounds.Max.Y / bounds.Max.X)

		a.logger.Infof("Proportionally resizing to the xsize value given: %dx%d", w, h)
	} else if ysize > 0 && xsize <= 0 {
		h = ysize
		w = int(h * bounds.Max.X / bounds.Max.Y)

		a.logger.Infof("Proportionally resizing to the ysize value given: %dx%d", w, h)
	}

	src = imaging.Resize(src, w, h, rFilter)
	a.logger.Info("Operation resize Completed.")

	return src, nil
}
