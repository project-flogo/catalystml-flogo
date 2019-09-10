package removeStop

import "github.com/project-flogo/core/data/coerce"

type Params struct {
	Lang string `md:"lang",required=false`
}

type Input struct {
	Str string `md:"str"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Str, err = coerce.ToString(values["str"])

	return err
}
