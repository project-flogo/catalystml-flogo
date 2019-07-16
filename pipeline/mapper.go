package pipeline

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"

	"github.com/project-flogo/fps/fpsmapper"
	"github.com/project-flogo/fps/operation"
)

var mapperFactory mapper.Factory

/*
	The DefaultOutputMapper is used when the output label is not mentioned
	for the output from the operation and OperationOutputMapper is used when the
	output label is defined
*/

func SetMapperFactory(factory mapper.Factory) {
	mapperFactory = factory
}

func GetMapperFactory() mapper.Factory {

	if mapperFactory == nil {
		mapperFactory = &NewDefaultMapperFactory{}
	}
	return mapperFactory
}

type NewDefaultMapperFactory struct {
}

func (n *NewDefaultMapperFactory) NewMapper(mappings map[string]interface{}) (mapper.Mapper, error) {

	if len(mappings) == 0 {
		return nil, nil
	}

	defMapper := make(map[string]interface{})
	for key, value := range mappings {
		if value != nil {
			switch t := value.(type) {
			case string:
				if len(t) > 0 {

					if strings.Contains(t, "[") {
						defMapper[key] = fpsmapper.NewExpression(t)

					} else {

						defMapper[key] = t
					}

				}
			default:
				defMapper[key] = t
			}

		}
	}

	return &defaultOperationOutputMapper{mappings: defMapper}, nil
}

func NewDefaultOperationOutputMapper(stage *Stage) mapper.Mapper {

	defMapper := make(map[string]interface{})

	for key, _ := range stage.opt.Metadata().Output {

		defMapper[stage.ID()] = key
	}
	return &defaultOperationOutputMapper{mappings: defMapper}
}

type defaultOperationOutputMapper struct {
	mappings map[string]interface{}
}

func (m *defaultOperationOutputMapper) Apply(scope data.Scope) (map[string]interface{}, error) {

	output := make(map[string]interface{}, len(m.mappings))

	for name := range m.mappings {

		switch t := m.mappings[name].(type) {
		case string:
			if mapper.IsLiteral(t) {
				output[name] = t
			} else {
				value, ok := scope.GetValue(t[1:])

				if ok {
					output[name] = value
				}
			}

		case []fpsmapper.DerefernceStruct:

			val, err := fpsmapper.Resolve(t, scope)
			if err != nil {
				return nil, err
			}
			output[name] = val

		default:

			output[name] = t
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
