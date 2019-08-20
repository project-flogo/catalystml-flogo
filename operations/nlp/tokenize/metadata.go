package tokenize

import "github.com/project-flogo/core/data/coerce"

type Params struct {
}

type Input struct {
	Str string `md:"str"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Str, err = coerce.ToString(values["str"])

	return err
}
