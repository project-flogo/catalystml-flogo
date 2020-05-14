package divide

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Num   interface{} `md:"num"`
	Denom interface{} `md:"denom"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Num, err = coerce.ToFloat64(values["num"])
	i.Denom, err = coerce.ToFloat64(values["denom"])

	return err

}
