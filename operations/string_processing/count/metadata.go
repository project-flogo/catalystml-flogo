package count

import "github.com/project-flogo/core/data/coerce"

type Params struct {
	// Lang string `md:"lang",required=false`
}

type Input struct {
	S0 string `md:"s0"`
	S1 string `md:"s1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.S0, err = coerce.ToString(values["s0"])
	i.S1, err = coerce.ToString(values["s1"])

	return err
}
