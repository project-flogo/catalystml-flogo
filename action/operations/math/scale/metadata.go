package scale

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Data   interface{} `md:"data"`
	Scaler float32     `md:"scaler"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToArray(values["data"])
	if err != nil {
		return err
	}
	i.Scaler, err = coerce.ToFloat32(values["scaler"])
	if err != nil {
		return err
	}

	return err
}
