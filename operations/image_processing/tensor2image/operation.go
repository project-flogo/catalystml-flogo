package tensor2image

import (
	"errors"
	"fmt"
	"image"
	"image/color"

	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	toFile bool
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()
	toFile := false
	filePath := p.ToFileFolder
	if "" != p.ToFileFolder {
		err = prepareForOutputFile(filePath, logger)
		if err != nil {
			return nil, err
		}
		toFile = true
	}

	return &Operation{
		params: p,
		toFile: toFile,
		logger: logger}, nil
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

	input := &Input{}
	input.FromMap(inputs)

	extension := a.params.Extension

	a.logger.Info("Starting Operation tensor2img.")
	a.logger.Debug("Extension is set to ...", extension)

	var out []image.Image
	tensor4, ok := input.Tensor.([][][][]uint8)
	if ok {
		out = make([]image.Image, len(tensor4))
		for index, value := range tensor4 {
			out[index] = tensor2Image(value)
		}
	} else {
		tensor3, ok := input.Tensor.([][][]uint8)
		if ok {
			out = make([]image.Image, 1)
			out[0] = tensor2Image(tensor3)
		} else {
			return nil, errors.New("Illegal tensor for image.")
		}
	}

	if a.toFile {
		filename := a.params.Filename
		if "" == filename {
			filename = "outimage"
		}

		for index, _ := range out {
			image2File(
				a.params.ToFileFolder,
				fmt.Sprintf("%s_%d", filename, index),
				a.params.Extension,
				out[index],
			)
		}
	}

	a.logger.Info("Operation tensor2image Completed.")
	return out, nil

}

func tensor2Image(tensor [][][]uint8) image.Image {
	width := len(tensor)
	height := len(tensor[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	var alpha uint8
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colorCode := tensor[x][y]
			red := colorCode[0]
			green := colorCode[1]
			blue := colorCode[2]
			alpha = 0xff
			if 4 == len(colorCode) {
				alpha = colorCode[3]
			}
			color := color.RGBA{red, green, blue, alpha}
			img.Set(x, y, color)
		}
	}
	return img
}

func image2File(filePath string, filename string, extension string, img image.Image) error {
	f, err := os.Create(fmt.Sprintf("%s/%s.%s", filePath, filename, extension))
	if err != nil {
		return err
	}
	defer f.Close()

	switch strings.ToLower(extension) {
	case "jpeg", "jpg":
		opt := jpeg.Options{}
		err = jpeg.Encode(f, img, &opt)
	case "png":
		err = png.Encode(f, img)
	case "gif":
		opt := gif.Options{}
		err = gif.Encode(f, img, &opt)
	}

	return err
}

func prepareForOutputFile(
	outputFolderPath string,
	logger log.Logger) error {

	err := os.MkdirAll(outputFolderPath, os.ModePerm)
	if nil != err {
		logger.Error("Unable to create folder ...")
	}

	return err
}
