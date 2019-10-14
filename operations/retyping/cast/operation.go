package cast

import (
	"reflect"

	c "github.com/spf13/cast"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

var ValLen int
var givenType data.Type

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	in := &Input{}
	err := in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	typ := inputs["toType"].(string)

	v := reflect.ValueOf(inputs["data"])
	switch v.Kind() {
	case reflect.Slice, reflect.Array:

		result := castslice(inputs["data"], typ)
		a.logger.Info("Cast Opertaion commencing on a Slice...", result)
		a.logger.Info("Results...", result)
		// out, _ := coerce.ToArray(result)
		return result, nil
	case reflect.Map:

		result := castmap(inputs["data"], typ)
		a.logger.Info("Cast Opertaion commencing on a Map...", result)
		a.logger.Info("Results...", result)
		// out, _ := coerce.ToObject(result)
		return result, nil

	default:

		result := casttype(inputs["data"], typ)
		a.logger.Info("Cast Opertaion commencing on a base data type ...", result)
		a.logger.Info("Results...", result)
		// out := result
		return result, nil
	}
}

func castslice(data interface{}, toType string) (out []interface{}) {

	d, _ := coerce.ToArray(data)
	for _, s := range d {
		v := reflect.ValueOf(s)
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			val := castslice(s, toType)
			out = append(out, val)
		case reflect.Map:
			val := castmap(s, toType)
			out = append(out, val)
		default:
			val := casttype(s, toType)
			out = append(out, val)
		}
	}

	return out
}

func castmap(data interface{}, toType string) (out map[string]interface{}) {
	out = make(map[string]interface{})

	d, _ := coerce.ToObject(data)
	for k, s := range d {
		v := reflect.ValueOf(s)
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			val := castslice(s, toType)
			out[k] = val
		case reflect.Map:
			val := castmap(s, toType)
			out[k] = val
		default:
			val := casttype(s, toType)
			out[k] = val
		}
	}

	return out
}

func casttype(val interface{}, toType string) (out interface{}) {

	switch toType {
	case "int64":
		out = c.ToInt64(val)
	case "int32":
		out = c.ToInt32(val)
	case "float64":
		out = c.ToFloat64(val)
	case "float32":
		out = c.ToFloat32(val)
	case "string":
		out = c.ToString(val)
	case "bool":
		out = c.ToBool(val)
	default:
		out = c.ToFloat64(val)
	}
	// out = 5
	return out
}

// func toInt32(val interface{}) (out int32) {
// 	return 5
// }
