package cleaning

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/cleaning/apply"
	"github.com/project-flogo/cml/operations/cleaning/dropCol"
	"github.com/project-flogo/cml/operations/cleaning/replaceValue"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
}
