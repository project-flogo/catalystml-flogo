package reshape

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Data  []interface{} `md:"data"`
	Shape []interface{} `md:"shape"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToArray(values["data"])
	if err != nil {
		return err
	}
	i.Shape, err = coerce.ToArray(values["shape"])
	if err != nil {

		return err
	}

	return err
}
