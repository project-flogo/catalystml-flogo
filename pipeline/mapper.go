package pipeline

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/fps/operation"
)

var mapperFactory mapper.Factory

func SetMapperFactory(factory mapper.Factory) {
	mapperFactory = factory
}

func GetMapperFactory() mapper.Factory {

	if mapperFactory == nil {
		mapperFactory = mapper.NewFactory(GetDataResolver())
	}
	return mapperFactory
}

func NewDefaultOperationOutputMapper(stage *Stage) mapper.Mapper {
	attrNS := stage.ID() + "."
	return &defaultOperationOutputMapper{attrNS: attrNS, metadata: stage.opt.Metadata()}
}

type defaultOperationOutputMapper struct {
	attrNS   string
	metadata *operation.Metadata
}

func (m *defaultOperationOutputMapper) Apply(scope data.Scope) (map[string]interface{}, error) {

	output := make(map[string]interface{}, len(m.metadata.Output))
	for name := range m.metadata.Output {

		value, ok := scope.GetValue(name)
		if ok {
			output[m.attrNS+name] = value
		}
	}

	return output, nil
}

type newOperationOutputMapper struct {
	metadata *operation.Metadata
}

func NewOperationOutputMapper(stage *Stage) mapper.Mapper {

	return &newOperationOutputMapper{metadata: stage.opt.Metadata()}
}

func (m *newOperationOutputMapper) Apply(scope data.Scope) (map[string]interface{}, error) {

	output := make(map[string]interface{}, len(m.metadata.Output))
	for name := range m.metadata.Output {

		value, ok := scope.GetValue(name)
		if ok {
			output[name] = value
		}
	}

	return output, nil
}
