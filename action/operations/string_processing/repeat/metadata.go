package repeat

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	S string `md:"s"`
	I int `md:"i"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.S, err = coerce.ToString(values["s"])
	i.I, err = coerce.ToInt(values["i"])

	return err
}
