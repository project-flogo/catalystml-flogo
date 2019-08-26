package mean

import (
	"errors"
	"fmt"
	"reflect"
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
	dataArray, err := ToInterfaceArray(values["data"])
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

func ToInterfaceArray(val interface{}) ([]interface{}, error) {

	switch t := val.(type) {
	case []interface{}:
		return t, nil

	case []map[string]interface{}:
		var a []interface{}
		for _, v := range t {
			a = append(a, v)
		}
		return a, nil
	case nil:
		return nil, nil
	default:
		s := reflect.ValueOf(val)
		if s.Kind() == reflect.Slice {
			a := make([]interface{}, s.Len())
			for i := 0; i < s.Len(); i++ {
				element := s.Index(i).Interface()
				elementType := reflect.TypeOf(element)
				switch elementType.Kind() {
				case reflect.Slice:
					a[i], _ = ToInterfaceArray(element)
				case reflect.Array:
					a[i], _ = ToInterfaceArray(element)
				default:
					a[i] = element
				}
			}
			return a, nil
		}
		return nil, fmt.Errorf("unable to coerce %#v to []interface{}", val)
	}
}
