package interpolateMissing

import (
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
)

type Params struct {
	How   string `md:"how"`
	Edges string `md:"edges"`
}

type Input struct {
	Data interface{} `md:"data"`
	Col  interface{} `md:"col"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = common.ToDataFrame(values["data"])
	i.Col = values["col"]

	return err
}
