package date

import (
	"errors"
	"fmt"
)

type Params struct {
	Format string `md:"format"`
}

type Input struct {
	Data string `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	var ok bool
	i.Data, ok = values["data"].(string)
	if !ok {
		err = errors.New(fmt.Sprintf("invalid input data : %v", values["data"]))
	}
	return err
}
