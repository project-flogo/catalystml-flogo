package cleaning

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/apply"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/binning"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/filter"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifnotin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/interpolateMissing"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/oneHotEncoding"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/replaceValue"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/set"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/sort"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&set.Operation{}, set.New)
	_ = operation.Register(&ifnotin.Operation{}, ifnotin.New)
	_ = operation.Register(&ifin.Operation{}, ifin.New)
	_ = operation.Register(&oneHotEncoding.Operation{}, oneHotEncoding.New)
	_ = operation.Register(&binning.Operation{}, binning.New)
	_ = operation.Register(&filter.Operation{}, filter.New)
	_ = operation.Register(&sort.Operation{}, sort.New)
	_ = operation.Register(&interpolateMissing.Operation{}, interpolateMissing.New)

}
