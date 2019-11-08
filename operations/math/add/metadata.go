package add

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Num0 interface{} `md:"num0"`
	Num1 interface{} `md:"num1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Num0, err = coerce.ToFloat64(values["num0"])
	i.Num1, err = coerce.ToFloat64(values["num1"])

	return err

}
