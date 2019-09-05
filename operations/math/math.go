package math

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/math/norm"
	"github.com/project-flogo/cml/operations/math/normalize"
	"github.com/project-flogo/cml/operations/math/scale"
)

func init() {
	_ = operation.Register(&norm.Operation{}, norm.New)
	_ = operation.Register(&normalize.Operation{}, normalize.New)
	_ = operation.Register(&scale.Operation{}, scale.New)
}
