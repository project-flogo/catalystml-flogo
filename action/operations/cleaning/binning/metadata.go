package binning

import (
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
)

type Params struct {
	Quantile   int       `md:"quantile"`
	Bins       []float64 `md:"bins"`
	Labels     []string  `md:"labels"`
	Column     string    `md:"column"`
	Retbins    bool      `md:"retbins"`
	Precision  int       `md:"precision"`
	Duplicates string    `md:"duplicates"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = common.ToDataFrame(values["data"])

	return err
}
