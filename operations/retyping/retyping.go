package retyping

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/retyping/map2table"
	"github.com/project-flogo/catalystml-flogo/operations/retyping/table2map"
	"github.com/project-flogo/catalystml-flogo/operations/retyping/cast"
)

func init() {
	_ = operation.Register(&table2map.Operation{}, table2map.New)
	_ = operation.Register(&map2table.Operation{}, map2table.New)
	_ = operation.Register(&cast.Operation{}, cast.New)
}
