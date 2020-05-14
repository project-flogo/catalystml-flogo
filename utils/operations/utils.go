package utils

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/utils/operations/runcml"
	"github.com/project-flogo/catalystml-flogo/utils/operations/toLog"
)

func init() {
	_ = operation.Register(&runcml.Operation{}, runcml.New)
	_ = operation.Register(&runcml.Operation{}, toLog.New)
}
