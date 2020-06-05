package utils

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/utils/runcml"
	"github.com/project-flogo/catalystml-flogo/operations/utils/toLog"
)

func init() {
	_ = operation.Register(&runcml.Operation{}, runcml.New)
	_ = operation.Register(&toLog.Operation{}, toLog.New)
}
