package apply

import (
	"fmt"
	"reflect"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)
	if p.MapOrArray == "" {
		p.MapOrArray = "array"
	}

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

type initContextImpl struct {
	params   map[string]interface{}
	mFactory mapper.Factory
	name     string
}

func (ctx *initContextImpl) Params() map[string]interface{} {
	return ctx.params
}

func (ctx *initContextImpl) MapperFactory() mapper.Factory {
	return ctx.mFactory
}

func (ctx *initContextImpl) Logger() log.Logger {
	return log.ChildLogger(log.RootLogger(), ctx.name)
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	catchkey := "&item"

	//To get the inputs in the desired types.
	input := &Input{}

	moa := a.params.MapOrArray //moa = Map or Array
	input.FromMap(inputs)

	a.logger.Debug("inputs", inputs)
	a.logger.Info("Executing operation apply to...", input.Function.Operation)

	opt := operation.Get(input.Function.Operation)
	if opt == nil {
		return nil, fmt.Errorf("unsupported Operation:" + input.Function.Operation)

	}

	f := operation.GetFactory(input.Function.Operation)
	if f != nil {
		initCtx := &initContextImpl{params: input.Function.Params, name: input.Function.Operation}
		pa, err := f(initCtx)
		if err != nil {
			return nil, fmt.Errorf("unable to create apply operation '%s' : %s", input.Function.Operation, err.Error())
		}
		opt = pa
	}

	outM := make(map[interface{}]interface{})
	var outSl []interface{}

	v := reflect.ValueOf(input.Data)

	switch v.Kind() {
	case reflect.Slice:

		loopOver, err := coerce.ToArray(input.Data)
		if err != nil {
			return nil, fmt.Errorf("error coercing to array")
		}

		for _, v := range loopOver {
			optint := make(map[string]interface{})

			for key, val := range input.Function.Input {
				if val == catchkey {

					optint[key] = v
				} else {
					optint[key] = val
				}
			}

			outitem, err := opt.Eval(optint)
			if err != nil {
				return nil, err
			}

			if moa == "map" {
				outM[v] = outitem
			} else if moa == "array" {
				outSl = append(outSl, outitem)
			}
		}
	case reflect.Map:
		loopOver := input.Data.(map[interface{}]interface{})

		for _, v := range loopOver {
			optint := make(map[string]interface{})
			for key, val := range input.Function.Input {
				if val == catchkey {
					optint[key] = v
				} else {
					optint[key] = val
				}
			}

			outitem, err := opt.Eval(optint)
			if err != nil {
				return nil, err
			}

			if moa == "map" {
				outM[v] = outitem
			} else if moa == "array" {
				outSl = append(outSl, outitem)
			}
		}
	}

	var out interface{}
	if moa == "map" {
		out = outM
	} else if moa == "array" {
		out = outSl
	}
	a.logger.Info("Output of apply", out)
	return out, nil
}
