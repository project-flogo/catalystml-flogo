package join

//import (
//	"fmt"
//	"reflect"
//)

type Params struct {
	On  []string `md:"on"`
	How string   `md:"how"`
}

type Input struct {
	Left       interface{} `md:"left"`
	Right      interface{} `md:"right"`
	LeftIndex  interface{} `md:"leftindex"`
	RightIndex interface{} `md:"rightindex"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Left, err = ToDataFrame(values["left"])
	i.Right, err = ToDataFrame(values["right"])
	i.LeftIndex, err = ToDataFrame(values["leftindex"])
	i.RightIndex, err = ToDataFrame(values["rightindex"])

	return err
}

func ToDataFrame(val interface{}) (interface{}, error) {
	return val, nil
}
