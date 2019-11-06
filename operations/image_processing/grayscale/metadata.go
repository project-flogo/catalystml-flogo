package resize

import (
	"bytes"
	"image"
)

type Config struct {
	Operation string                 `json:"operation"`
	Input     map[string]interface{} `json:"input,omitempty"`
}

type Input struct {
	Img image.Image `md:"img"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	i.Img, _, _ = image.Decode(bytes.NewReader(values["img"].([]byte)))

	return nil
}
