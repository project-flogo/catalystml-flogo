package pivot

//import (
//	"fmt"
//	"reflect"
//)

type Params struct {
	Index     []string            `md:"index"`
	Columns   []string            `md:"columns"`
	Aggregate map[string][]string `md:"aggregate"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	i.Data = values["data"]

	return nil
}
