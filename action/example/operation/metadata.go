package operation

import "github.com/project-flogo/core/data/coerce"

type Params struct {
	Sample     float64       `md:"sample"`
	ListOfKeys []interface{} `md:"listOfKeys"`
}
type Input struct {
	InputSample int `md:"inputSample"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.InputSample, err = coerce.ToInt(values["inputSample"])

	return err
}
