package sort

import (
	//	"fmt"

	"github.com/project-flogo/catalystml-flogo/action/operations/common"
)

type Params struct {
	Ascending   bool          `md:"ascending"`
	NilPosition string        `md:"nilPosition",allowed=["first","last"],required=false`
	By          []interface{} `md:"by"`
	Axis        int           `md:"axis",allowed=[0,1],required=false`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Data, err = common.ToDataFrame(values["data"])

	return err
}
