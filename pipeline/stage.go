package pipeline

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/fps/operation"
)

type Stage struct {
	id  string
	opt operation.Operation

	params map[string]interface{}

	outputAttrs map[string]interface{}

	inputMapper  mapper.Mapper
	outputMapper mapper.Mapper
}

type StageConfig struct {
	*operation.Config
}

type initContextImpl struct {
	params   map[string]interface{}
	mFactory mapper.Factory
}

func (ctx *initContextImpl) Params() map[string]interface{} {
	return ctx.params
}

func (ctx *initContextImpl) MapperFactory() mapper.Factory {
	return ctx.mFactory
}

func (ctx *initContextImpl) Logger() log.Logger {
	return log.RootLogger()
}

func NewStage(config *StageConfig, mf mapper.Factory, resolver resolve.CompositeResolver) (*Stage, error) {

	if config.Operation == "" {
		return nil, fmt.Errorf("Operation not specified for Stage")
	}

	opt := operation.Get(config.Operation)

	if opt == nil {
		return nil, fmt.Errorf("unsupported Operation:" + config.Operation)
	}

	f := operation.GetFactory(config.Operation)

	if f != nil {
		initCtx := &initContextImpl{params: config.Config.Params, mFactory: mf}
		pa, err := f(initCtx)
		if err != nil {
			return nil, fmt.Errorf("unable to create stage '%s' : %s", config.Operation, err.Error())
		}
		opt = pa
	}

	stage := &Stage{}
	stage.id = config.Id
	stage.opt = opt

	settingsMd := opt.Metadata().Params

	if len(config.Params) > 0 && settingsMd != nil {
		stage.params = make(map[string]interface{}, len(config.Params))

		for name, value := range config.Params {

			attr := settingsMd[name]

			if attr != nil {
				//todo handle error
				stage.params[name] = resolveParamsValue(resolver, name, value)
			}
		}
	}

	input := make(map[string]interface{})
	mf = GetMapperFactory()

	for k, v := range config.Input {
		if !isExpr(v) {
			fieldMetaddata, ok := opt.Metadata().Input[k]
			if ok {
				v, err := coerce.ToType(v, fieldMetaddata.Type())
				if err != nil {
					return nil, fmt.Errorf("convert value [%+v] to type [%s] error: %s", v, fieldMetaddata.Type(), err.Error())
				}
				input[k] = v
			} else {
				//For the cases that metadata comes from iometadata, eg: subflow
				input[k] = v
			}
		} else {
			input[k] = v
		}

	}
	var err error

	stage.inputMapper, err = mf.NewMapper(input)

	if err != nil {
		return nil, err
	}

	if config.Output == nil {
		//If the output label is not defined use the default mapper ie. `$id`
		stage.outputMapper = NewDefaultOperationOutputMapper(stage)

	} else {
		//If the output Label is defined use the new one.
		stage.outputAttrs = config.Output
		stage.outputMapper = NewOperationOutputMapper(stage)
	}

	return stage, nil
}
func (stage *Stage) ID() string {
	return stage.id
}

func isExpr(v interface{}) bool {
	switch t := v.(type) {
	case string:
		if len(t) > 0 && t[0] == '$' {
			return true
		}
	default:
		if _, ok := mapper.GetObjectMapping(t); ok {
			return true
		}
	}
	return false
}
