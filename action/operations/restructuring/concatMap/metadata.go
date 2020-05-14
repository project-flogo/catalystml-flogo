package concatMap

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Data []interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToArray(values["data"])
	if err != nil {
		return err
	}

	return nil
}
