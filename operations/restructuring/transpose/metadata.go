package transpose

import (
	"github.com/project-flogo/catalystml-flogo/operations/common"
)

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = common.ToInterfaceArray(values["data"])

	return err
}
