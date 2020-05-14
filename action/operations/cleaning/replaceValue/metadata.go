package replaceValue

import (
	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	Columns interface{} `md:"col"`
}

type Input struct {
	Data       map[string]interface{} `md:"data"`
	ReplaceMap map[string]interface{} `md:"replaceMap"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToObject(values["data"])
	if err != nil {

		return err
	}
	i.ReplaceMap, err = coerce.ToObject(values["replaceMap"])
	if err != nil {

		return err
	}

	return err
}
