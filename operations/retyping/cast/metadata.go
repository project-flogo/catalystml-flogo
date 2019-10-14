package cast

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Data   interface{} `md:"data"`
	ToType string      `md:"string",allowed=["int64","float64","string","int32","float32","boolean"]`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Data = values["data"]

	i.ToType, err = coerce.ToString(values["string"])

	if err != nil {
		return err
	}

	return nil
}
