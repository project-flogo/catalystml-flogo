package normalize

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Data  interface{} `md:"data"`
	Value float32     `md:"value"`
	Min   float32     `md:"minvalue"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToArray(values["data"])
	if err != nil {
		return err
	}
	i.Value, err = coerce.ToFloat32(values["value"])
	if err != nil {
		return err
	}
	i.Min, err = coerce.ToFloat32(values["minvalue"])
	if err != nil {
		return err
	}

	return err
}
