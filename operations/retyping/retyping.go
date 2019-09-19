package retyping

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/retyping/map2table"
	"github.com/project-flogo/catalystml-flogo/operations/retyping/table2map"
)

func init() {
	_ = operation.Register(&table2map.Operation{}, table2map.New)
	_ = operation.Register(&map2table.Operation{}, map2table.New)
}
