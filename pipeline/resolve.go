package pipeline

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/resolve"
)

var pipelineRes = resolve.NewCompositeResolver(map[string]resolve.Resolver{
	".":         &resolve.ScopeResolver{},
	"env":       &resolve.EnvResolver{},
	"property":  &resolve.PropertyResolver{},
	"operation": &OperationResolver{}})

func GetDataResolver() resolve.CompositeResolver {
	return pipelineRes
}

var resolverInfo = resolve.NewResolverInfo(false, true)

type OperationResolver struct {
}

func (r *OperationResolver) GetResolverInfo() *resolve.ResolverInfo {
	return resolverInfo
}

func (r *OperationResolver) Resolve(scope data.Scope, itemName, valueName string) (interface{}, error) {

	value, exists := scope.GetValue("_O." + itemName + "." + valueName)
	if !exists {
		return nil, fmt.Errorf("failed to resolve activity attr: '%s', not found in opeartion '%s'", valueName, itemName)
	}

	return value, nil

}

func resolveParamsValue(resolver resolve.CompositeResolver, params string, value interface{}) interface{} {

	strVal, ok := value.(string)

	if ok && len(strVal) > 0 && strVal[0] == '$' {
		v, err := resolver.Resolve(strVal, nil)

		if err == nil {

			logger.Debugf("Resolved params [%s: %s] to : %v", params, value, v)
			return v
		}
	}

	return value
}
