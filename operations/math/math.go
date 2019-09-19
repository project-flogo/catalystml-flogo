package math

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/math/mean"
	"github.com/project-flogo/catalystml-flogo/operations/math/norm"
	"github.com/project-flogo/catalystml-flogo/operations/math/normalize"
	"github.com/project-flogo/catalystml-flogo/operations/math/scale"
	"github.com/project-flogo/catalystml-flogo/operations/math/multPairWise"
)

func init() {
	_ = operation.Register(&mean.Operation{}, mean.New)
	_ = operation.Register(&norm.Operation{}, norm.New)
	_ = operation.Register(&normalize.Operation{}, normalize.New)
	_ = operation.Register(&scale.Operation{}, scale.New)
	_ = operation.Register(&multPairWise.Operation{}, multPairWise.New)
}
