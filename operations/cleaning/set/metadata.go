package set

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Arr []interface{} `md:"arr"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Arr, err = coerce.ToArray(values["arr"])

	return err
}
