package split

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Str string `md:"str"`
	Sep string    `md:"sep"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Str, err = coerce.ToString(values["str"])
	i.Sep, err = coerce.ToString(values["sep"])

	return err
}
