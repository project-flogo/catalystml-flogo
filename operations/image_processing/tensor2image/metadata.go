package tensor2image

import (
	"errors"
)

type Params struct {
	Extension    string `md:"extension", required=true`
	ToFileFolder string `md:"toFileFolder"`
	Filename     string `md:"filename"`
}

type Config struct {
	Operation string                 `json:"operation"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Input     map[string]interface{} `json:"input,omitempty"`
}

type Input struct {
	Tensor interface{} `md:"tensor"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Tensor = values["tensor"]
	if nil == i.Tensor {
		err = errors.New("Input tesnsor is nil.")
	}

	return err
}
