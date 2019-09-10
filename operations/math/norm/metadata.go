package norm

import (
	"errors"
	"reflect"

	"github.com/project-flogo/cml/operations/common"
)

type Params struct {
	Axis int `md:"axis"`
}

type Input struct {
	Data   interface{} `md:"data"`
	isFlat bool
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	dataArray, err := common.ToInterfaceArray(values["data"])
	if nil != err {
		return err
	}

	if nil == dataArray || 0 == len(dataArray) {
		return errors.New("Empty array.")
	}

	elementType := reflect.ValueOf(dataArray[0]).Kind()
	if reflect.Slice == elementType || reflect.Array == elementType {
		i.Data = dataArray
		i.isFlat = false
	} else {
		i.Data = make([]interface{}, 1)
		i.Data.([]interface{})[0] = dataArray
		i.isFlat = true
	}

	return err
}
