package ifin

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Arr0 []interface{} `md:"arr0"`
	Arr1 []interface{} `md:"arr1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Arr0, err = coerce.ToArray(values["arr0"])
	i.Arr1, err = coerce.ToArray(values["arr1"])

	return err
}
