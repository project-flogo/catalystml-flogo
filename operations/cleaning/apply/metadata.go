package apply

import (
	"encoding/json"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	_ "github.com/project-flogo/catalystml-flogo/operations/string_processing/count"
	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	MapOrArray string `md:"mapOrArray",allowed=["map","array"],required=false`
}

type Input struct {
	Data     interface{}      `md:"data"`
	Function operation.Config `md:"function"`
}

// type Config struct {
// 	Operation string                 `md:"operation"`
// 	Params    map[string]interface{} `md:"params"`
// 	Input     map[string]interface{} `md:"input"`
// }

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToAny(values["data"])
	b, err := coerce.ToBytes(values["function"])
	err = json.Unmarshal(b, &i.Function)

	return err
}
