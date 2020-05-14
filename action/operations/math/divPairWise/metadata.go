package divPairWise

import (
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
)

type Input struct {
	Matrix0 []interface{} `md:"matrix0"`
	Matrix1 []interface{} `md:"matrix1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Matrix0, err = common.ToInterfaceArray(values["matrix0"])
	i.Matrix1, err = common.ToInterfaceArray(values["matrix1"])

	return err

}
