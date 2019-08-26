package cleaning

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/cleaning/apply"
	"github.com/project-flogo/cml/operations/cleaning/dropCol"
	"github.com/project-flogo/cml/operations/cleaning/groupBy"
	"github.com/project-flogo/cml/operations/cleaning/join"
	"github.com/project-flogo/cml/operations/cleaning/pivot"
	"github.com/project-flogo/cml/operations/cleaning/replaceValue"
	"github.com/project-flogo/cml/operations/cleaning/transpose"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&groupBy.Operation{}, groupBy.New)
	_ = operation.Register(&join.Operation{}, join.New)
	_ = operation.Register(&pivot.Operation{}, pivot.New)
	_ = operation.Register(&transpose.Operation{}, transpose.New)
}
