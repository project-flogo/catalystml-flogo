package replaceregex

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Regex string `md:"regex"`
	S0    string `md:"s0"`
	S1    string `md:"s1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.S0, err = coerce.ToString(values["s0"])
	i.S1, err = coerce.ToString(values["s1"])
	i.Regex, err = coerce.ToString(values["regex"])

	return err
}
