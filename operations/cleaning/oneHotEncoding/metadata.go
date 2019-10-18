package oneHotEncoding

import (
	"github.com/project-flogo/catalystml-flogo/operations/common"
)

type Params struct {
	InputColumns  []interface{} `md:"inputColumns"`
	OutputColumns []interface{} `md:"outputColumns",required=false`
	KeepOrig      bool          `md:"keepOrig"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var m map[string]interface{}

	df, _ := common.ToDataFrame(values["data"])
	m = make(map[string]interface{})
	for k, arr := range df {
		m[k] = arr
	}

	i.Data = m

	return nil
}
