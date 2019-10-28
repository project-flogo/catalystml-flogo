package subsectiontoimage

import (
	"bytes"
	"image"
	// "github.com/project-flogo/core/data/coerce"
)

type Params struct {
	Size            []int `md:"size",required=false`
	LowerLeftCorner []int `md:"lowerLeftCorner",required=false`
	// OppositeCorners [][]int `md:"oppositeCorners",required=false`
	// Center          []int   `md:"center",required=false`
}

type Config struct {
	Operation string                 `json:"operation"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Input     map[string]interface{} `json:"input,omitempty"`
}

type Input struct {
	Img image.Image `md:"img"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	// i.Img = values["img"].(image.Image)
	i.Img, _, _ = image.Decode(bytes.NewReader(values["img"].([]byte)))

	return nil
}
