package cleaning

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/cleaning/addCol2Table"
	"github.com/project-flogo/cml/operations/cleaning/apply"
	"github.com/project-flogo/cml/operations/cleaning/dropCol"
	"github.com/project-flogo/cml/operations/cleaning/replaceValue"
	"github.com/project-flogo/cml/operations/cleaning/set"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&addCol2Table.Operation{}, addCol2Table.New)
	_ = operation.Register(&set.Operation{}, set.New)
}
