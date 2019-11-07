package math

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/addCol2Table"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/cast"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/concatMap"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/dropCol"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/flatten"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/groupBy"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/join"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/map2table"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/pivot"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/reshape"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/table2map"
	"github.com/project-flogo/catalystml-flogo/operations/restructuring/transpose"
)

func init() {
	_ = operation.Register(&addCol2Table.Operation{}, addCol2Table.New)
	_ = operation.Register(&cast.Operation{}, cast.New)
	_ = operation.Register(&concatMap.Operation{}, concatMap.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&flatten.Operation{}, flatten.New)
	_ = operation.Register(&groupBy.Operation{}, groupBy.New)
	_ = operation.Register(&join.Operation{}, join.New)
	_ = operation.Register(&map2table.Operation{}, map2table.New)
	_ = operation.Register(&pivot.Operation{}, pivot.New)
	_ = operation.Register(&reshape.Operation{}, reshape.New)
	_ = operation.Register(&table2map.Operation{}, table2map.New)
	_ = operation.Register(&transpose.Operation{}, transpose.New)
}
