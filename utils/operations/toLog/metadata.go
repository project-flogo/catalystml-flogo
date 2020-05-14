package toLog

import (
	"errors"
)

type Params struct {
	ToFilePath      string `md:"toFilePath"`
	ClearWhileStart bool   `md:"clearWhileStart"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Data = values["data"]
	if nil == i.Data {
		err = errors.New("Nil input data.")
	}
	return err
}
