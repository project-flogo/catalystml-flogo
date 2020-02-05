package pipeline

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"

	"github.com/project-flogo/catalystml-flogo/action/cmlmapper"
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

func NewDefaultOperationOutputMapper(stage *Stage) mapper.Mapper {

	defMapper := make(map[string]interface{})

	//defMapper[stage.output] = "$" + stage.output
	t := stage.output

	if len(t) > 0 {

		if strings.Contains(t, "[") {
			defMapper[t] = cmlmapper.NewExpression(t)

		} else {

			defMapper[t] = t
		}

	}

	return &defaultOperationOutputMapper{mappings: defMapper}
}

// Newmapper gets the mapper from the mappings.
func (n *NewDefaultMapperFactory) NewMapper(mappings map[string]interface{}) (mapper.Mapper, error) {

	if len(mappings) == 0 {
		return nil, nil
	}

	defMapper := make(map[string]interface{})
	for key, value := range mappings {
		if value != nil {
			switch t := value.(type) {
			case string:
				if len(t) > 0 && t[0] == '$' {

					if strings.Contains(t, "[") {
						defMapper[key] = cmlmapper.NewExpression(t)

					} else {

						defMapper[key] = t
					}

				} else {

					defMapper[key] = t

				}
			default:
				defMapper[key] = t
			}

		}
	}

	return &defaultOperationOutputMapper{mappings: defMapper}, nil
}

type defaultOperationOutputMapper struct {
	mappings map[string]interface{}
}

// Apply the mapper using the scope
func (m *defaultOperationOutputMapper) Apply(scope data.Scope) (map[string]interface{}, error) {

	output := make(map[string]interface{}, len(m.mappings))

	for name := range m.mappings {
		if name == "function" {
			mf := GetMapperFactory()

			funcInputMapper, _ := mf.NewMapper(m.mappings["function"].(map[string]interface{})["input"].(map[string]interface{}))

			m.mappings["function"].(map[string]interface{})["input"], _ = funcInputMapper.Apply(scope)
			output["function"] = m.mappings["function"]

		}
		switch t := m.mappings[name].(type) {
		case string:
			if t[0] != '$' {

				output[name] = t
			} else {
				value, ok := scope.GetValue(m.mappings[name].(string)[1:])

				if ok {
					output[name] = value
				}
			}

		case []cmlmapper.DerefernceStruct:

			val, err := cmlmapper.Resolve(t, scope)
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
