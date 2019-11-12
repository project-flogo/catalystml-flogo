package valToArray

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Value interface{}   `md:"value"`
	Shape []interface{} `md:"shape"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Value = values["value"]
	i.Shape, err = coerce.ToArray(values["shape"])

	return err
}
