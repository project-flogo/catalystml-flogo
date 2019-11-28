package pipeline

import (
	"github.com/project-flogo/core/data/resolve"
)

var pipelineRes = resolve.NewCompositeResolver(map[string]resolve.Resolver{
	".":   &resolve.ScopeResolver{},
	"env": &resolve.EnvResolver{},
})

func GetDataResolver() resolve.CompositeResolver {
	return pipelineRes
}

var resolverInfo = resolve.NewResolverInfo(false, true)

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
