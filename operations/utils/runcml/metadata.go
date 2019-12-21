package runcml

import "github.com/project-flogo/core/data/coerce"

type Params struct {
	CatalystMlURI string `md:"catalystMlURI"`
}

type Input struct {
	Data map[string]interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToObject(values["data"])

	return err
}
