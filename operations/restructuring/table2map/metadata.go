package table2map

import (
	"errors"

	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	Axis int `md:"axis"`
}

type Input struct {
	Table   []interface{} `md:"table"`
	ColKeys []interface{} `md:"colKeys"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Table, err = coerce.ToArray(values["table"])

	if err != nil {
		return err
	}

	i.ColKeys, err = coerce.ToArray(values["colKeys"])

	if err != nil {
		return errors.New("Column Keys not found")
	}

	return nil
}
