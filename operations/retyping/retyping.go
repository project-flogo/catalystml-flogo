package retyping

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/retyping/map2table"
	"github.com/project-flogo/cml/operations/retyping/table2map"
)

func init() {
	_ = operation.Register(&table2map.Operation{}, table2map.New)
	_ = operation.Register(&map2table.Operation{}, map2table.New)
}
