package replace

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	S0 string `md:"s0"`
	S1 string `md:"s1"`
	S2 string `md:"s2"`
	I  int    `md:"i"s`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.S0, err = coerce.ToString(values["s0"])
	i.S1, err = coerce.ToString(values["s1"])
	i.S2, err = coerce.ToString(values["s2"])
	i.I, err = coerce.ToInt(values["i"])

	return err
}
