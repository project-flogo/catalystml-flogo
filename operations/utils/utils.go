package utils

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/utils/runcml"
)

func init() {
	_ = operation.Register(&runcml.Operation{}, runcml.New)
}