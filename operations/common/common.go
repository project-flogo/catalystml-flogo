package common

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/project-flogo/core/data/coerce"
)

type Index struct {
	Id uint64
}

func NewIndex(elements []interface{}) Index {
	keyBytes := []byte{}
	for _, element := range elements {
		elementBytes, _ := json.Marshal(element)
		keyBytes = append(keyBytes, elementBytes...)
	}
	hasher := md5.New()
	hasher.Write(keyBytes)
	return Index{Id: binary.BigEndian.Uint64(hasher.Sum(nil))}
}

type DataState interface {
	Update(newData interface{}) error
	Value() interface{}
}

func GetFunction(functionName string) DataState {
	var function DataState
	if "sum" == functionName {
		function = &Sum{}
	} else if "count" == functionName {
		function = &Count{}
	} else if "mean" == functionName {
		function = &Mean{}
	} else if "min" == functionName {
		function = &Min{}
	} else if "max" == functionName {
		function = &Max{}
	}
	return function
}

type Data struct {
	data interface{}
}

func (this *Data) Value() interface{} {
	return this.data
}

func (this *Data) Update(newData interface{}) error {
	this.data = newData
	return nil
}

type Sum struct {
	data float64
}

func (this *Sum) Value() interface{} {
	return this.data
}

func (this *Sum) Update(newData interface{}) error {
	delta, _ := coerce.ToFloat64(newData)
	this.data += delta
	return nil
}

type Count struct {
	counter int
}

func (this *Count) Value() interface{} {
	return this.counter
}

func (this *Count) Update(newData interface{}) error {
	this.counter += 1
	return nil
}

type Mean struct {
	sum   float64
	count float64
}

func (this *Mean) Value() interface{} {
	return this.sum / this.count
}

func (this *Mean) Update(newData interface{}) error {
	this.count += 1
	delta, err := coerce.ToFloat64(newData)
	if nil != err {
		return err
	}
	this.sum += delta
	return nil
}

type Min struct {
	min interface{}
}

func (this *Min) Value() interface{} {
	return this.min
}

func (this *Min) Update(newData interface{}) error {
	if nil == this.min {
		this.min = newData
		return nil
	}

	result, err := compare(newData, this.min)

	if nil != err {
		return err
	}

	if 0 > result {
		this.min = newData
	}
	return nil
}

type Max struct {
	max interface{}
}

func (this *Max) Value() interface{} {
	return this.max
}

func (this *Max) Update(newData interface{}) error {
	if nil == this.max {
		this.max = newData
		return nil
	}

	result, err := compare(newData, this.max)

	if nil != err {
		return err
	}

	if 0 < result {
		this.max = newData
	}
	return nil
}

func compare(data1 interface{}, data2 interface{}) (int, error) {

	switch data1.(type) {
	case float64:
		delta1float64, _ := coerce.ToFloat64(data1)
		delta2float64, err := coerce.ToFloat64(data2)
		if nil != err {
			return 0, err
		}
		delta := delta1float64 - delta2float64
		switch {
		case delta > 0:
			return 1, nil
		case delta == 0:
			return 0, nil
		case delta < 0:
			return -1, nil
		}
	case int:
		delta1int, _ := coerce.ToInt(data1)
		delta2int, err := coerce.ToInt(data2)
		if nil != err {
			return 0, err
		}
		delta := delta1int - delta2int
		switch {
		case delta > 0:
			return 1, nil
		case delta == 0:
			return 0, nil
		case delta < 0:
			return -1, nil
		}
	}

	return 0, errors.New("Unable to compare, Uknown type!")
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