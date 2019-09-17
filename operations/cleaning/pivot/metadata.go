package pivot

import (
	"github.com/project-flogo/cml/operations/common"
)

type Params struct {
	Index     []string            `md:"index"`
	Columns   []string            `md:"columns"`
	Aggregate map[string][]string `md:"aggregate"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = common.ToDataFrame(values["data"])

	return err
}

func ToDataFrame(val interface{}) (interface{}, error) {
	return val, nil
}
