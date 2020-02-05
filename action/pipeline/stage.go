package pipeline

import (
	"fmt"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/catalystml-flogo/action/operation"
)

type Stage struct {
	opt operation.Operation

	params map[string]interface{}

	outputAttrs map[string]interface{}

	inputMapper  mapper.Mapper
	outputMapper mapper.Mapper
	output       string
	name         string
}

type StageConfig struct {
	*operation.Config
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

// NewStage gets the satage from the stage config.
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
		initCtx := &initContextImpl{params: config.Config.Params, mFactory: mf, name: config.Operation}
		pa, err := f(initCtx)
		if err != nil {
			return nil, fmt.Errorf("unable to create stage '%s' : %s", config.Operation, err.Error())
		}
		opt = pa
	}

	stage := &Stage{}
	if config.Output == "" {
		return nil, fmt.Errorf("Output not defined for operation %s", config.Operation)
	}
	stage.output = config.Output
	stage.opt = opt
	stage.name = config.Operation

	input := make(map[string]interface{})
	mf = GetMapperFactory()

	for k, v := range config.Input {

		input[k] = v
	}
	var err error

	// Set the input mapper.
	stage.inputMapper, err = mf.NewMapper(input)

	if err != nil {
		return nil, err
	}

	// Set the output mapper.
	stage.outputMapper = NewDefaultOperationOutputMapper(stage)

	return stage, nil
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
