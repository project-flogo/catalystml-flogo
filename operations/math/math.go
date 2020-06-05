package math

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/math/add"
	"github.com/project-flogo/catalystml-flogo/operations/math/addPairWise"
	"github.com/project-flogo/catalystml-flogo/operations/math/divPairWise"
	"github.com/project-flogo/catalystml-flogo/operations/math/divide"
	"github.com/project-flogo/catalystml-flogo/operations/math/mean"
	"github.com/project-flogo/catalystml-flogo/operations/math/multPairWise"
	"github.com/project-flogo/catalystml-flogo/operations/math/multi"
	"github.com/project-flogo/catalystml-flogo/operations/math/norm"
	"github.com/project-flogo/catalystml-flogo/operations/math/normalize"
	"github.com/project-flogo/catalystml-flogo/operations/math/scale"
)

func init() {
	_ = operation.Register(&mean.Operation{}, mean.New)
	_ = operation.Register(&norm.Operation{}, norm.New)
	_ = operation.Register(&normalize.Operation{}, normalize.New)
	_ = operation.Register(&scale.Operation{}, scale.New)
	_ = operation.Register(&multPairWise.Operation{}, multPairWise.New)
	_ = operation.Register(&divPairWise.Operation{}, divPairWise.New)
	_ = operation.Register(&addPairWise.Operation{}, addPairWise.New)
	_ = operation.Register(&multi.Operation{}, multi.New)
	_ = operation.Register(&add.Operation{}, add.New)
	_ = operation.Register(&divide.Operation{}, divide.New)
}
