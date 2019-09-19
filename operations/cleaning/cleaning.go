package cleaning

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/addCol2Table"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/apply"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/dropCol"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/groupBy"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifnotin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/join"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/concatMap"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/pivot"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/replaceValue"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/set"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/transpose"

)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&addCol2Table.Operation{}, addCol2Table.New)
	_ = operation.Register(&set.Operation{}, set.New)
	_ = operation.Register(&ifnotin.Operation{}, ifnotin.New)
	_ = operation.Register(&ifin.Operation{}, ifin.New)
	_ = operation.Register(&groupBy.Operation{}, groupBy.New)
	_ = operation.Register(&join.Operation{}, join.New)
	_ = operation.Register(&concatMap.Operation{}, concatMap.New)
	_ = operation.Register(&transpose.Operation{}, transpose.New)
	_ = operation.Register(&pivot.Operation{}, pivot.New)

}
