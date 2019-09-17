package join

import (
	"errors"
	"reflect"

	"github.com/project-flogo/cml/operations/common"
)

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
	i.Left, err = common.ToDataFrame(values["left"])
	i.Right, err = common.ToDataFrame(values["right"])
	i.LeftIndex, err = CheckIndex(values["leftindex"])
	i.RightIndex, err = CheckIndex(values["rightindex"])

	return err
}

func CheckIndex(val interface{}) (interface{}, error) {
	if nil == val {
		return nil, errors.New("Index should not be nil.")
	}

	if reflect.ValueOf(val).Kind() != reflect.Slice {
		return nil, errors.New("Index should be slice type.")
	}

	return val, nil
}
