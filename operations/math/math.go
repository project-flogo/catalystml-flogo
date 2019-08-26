package math

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/math/mean"
	"github.com/project-flogo/cml/operations/math/norm"
	"github.com/project-flogo/cml/operations/math/normalize"
)

func init() {
	_ = operation.Register(&mean.Operation{}, mean.New)
	_ = operation.Register(&norm.Operation{}, norm.New)
	_ = operation.Register(&normalize.Operation{}, normalize.New)
}
