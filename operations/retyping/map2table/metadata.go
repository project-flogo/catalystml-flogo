package map2table

import (
	"errors"

	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	Axis int `md:"axis"`
}

type Input struct {
	Map      map[string]interface{} `md:"map"`
	ColOrder []interface{}          `md:"colOrder"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Map, err = coerce.ToObject(values["map"])

	if err != nil {
		return err
	}

	i.ColOrder, err = coerce.ToArray(values["colOrder"])

	if err != nil {
		return errors.New("Column Keys not found")
	}

	return nil
}
