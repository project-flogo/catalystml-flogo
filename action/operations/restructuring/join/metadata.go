package join

import (
	"errors"
	"reflect"
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
	i.Left, err = CheckDataFrame(values["left"])
	i.Right, err = CheckDataFrame(values["right"])
	i.LeftIndex = values["leftindex"]
	i.RightIndex = values["rightindex"]

	return err
}

func CheckDataFrame(val interface{}) (interface{}, error) {
	if nil == val {
		return nil, errors.New("Data frame should not be nil.")
	}

	if reflect.ValueOf(val).Kind() != reflect.Map {
		return nil, errors.New("Data frame should be map type.")
	}

	return val, nil
}
