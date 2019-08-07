package dropCol

import (
	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	Columns     []interface{} `md:"col"`
	SeperateOut bool          `md:"seperateOut"`
}

type Input struct {
	Data map[string]interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToObject(values["data"])

	return err
}
