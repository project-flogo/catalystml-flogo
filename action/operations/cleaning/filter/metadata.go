package filter

import (
	//	"fmt"

	"github.com/project-flogo/catalystml-flogo/action/operations/common"
)

type Params struct {
	Axis int    `md:"axis",allowed=[0,1]`
	Col  string `md:"col"`
}

type Input struct {
	Data       interface{} `md:"data"`
	Value      interface{} `md:"value"`
	FilterType string      `md:"filterType"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Data, err = common.ToDataFrame(values["data"])
	i.Value = values["value"]
	var ok bool
	i.FilterType, ok = values["filterType"].(string)
	if !ok {
		i.FilterType = "Remove"
	}

	return err
}
