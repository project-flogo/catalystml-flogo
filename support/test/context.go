package test

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/fps/operation"
)

func NewOperationContext(md *operation.Metadata) *OperationContext {

	optContext := &OperationContext{name: "sample"}
	optContext.inputs = make(map[string]interface{})
	optContext.outputs = make(map[string]interface{})
	for name, tv := range md.Input {
		optContext.inputs[name] = tv.Value()
	}
	for name, tv := range md.Output {
		optContext.outputs[name] = tv.Value()
	}

	return optContext
}

type OperationContext struct {
	inputs  map[string]interface{}
	outputs map[string]interface{}
	name    string
	logger  log.Logger
}

func (eCtx *OperationContext) Name() string {
	return eCtx.name
}

func (eCtx *OperationContext) GetInput(name string) interface{} {

	value, found := eCtx.inputs[name]
	if found {
		return value
	}

	return nil
}

func (eCtx *OperationContext) GetOutput(name string) interface{} {
	value, found := eCtx.outputs[name]
	if found {
		return value
	}

	return nil
}

func (eCtx *OperationContext) GetInputObject(input data.StructValue) error {
	err := input.FromMap(eCtx.inputs)
	return err
}

func (eCtx *OperationContext) Logger() log.Logger {

	return log.RootLogger()
}

func (eCtx *OperationContext) SetOutput(name string, value interface{}) error {

	if eCtx.outputs == nil {
		eCtx.outputs = make(map[string]interface{})
	}

	//todo coerce to type based on metadata
	eCtx.outputs[name] = value

	return nil
}

func (eCtx *OperationContext) SetOutputObject(output data.StructValue) error {
	eCtx.outputs = output.ToMap()
	return nil
}

func (eCtx *OperationContext) SetInputObject(input data.StructValue) error {
	eCtx.inputs = input.ToMap()
	return nil
}

func NewOperationInitContext(params interface{}, f mapper.Factory) operation.InitContext {

	var settingVals map[string]interface{}

	if sm, ok := params.(map[string]interface{}); ok {
		settingVals = sm
	} else {
		settingVals = metadata.StructToMap(params)
	}

	if f == nil {
		f = mapper.NewFactory(resolve.GetBasicResolver())
	}

	return &TestOperationInitContext{params: settingVals, factory: f}
}

type TestOperationInitContext struct {
	params  map[string]interface{}
	factory mapper.Factory
}

func (ic *TestOperationInitContext) Params() map[string]interface{} {
	return ic.params
}

func (ic *TestOperationInitContext) MapperFactory() mapper.Factory {
	return ic.factory
}

func (ic *TestOperationInitContext) Logger() log.Logger {
	return log.RootLogger()
}
