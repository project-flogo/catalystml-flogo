package cleaning

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/apply"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/concatMap"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/ifnotin"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/oneHotEncoding"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/replaceValue"
	"github.com/project-flogo/catalystml-flogo/operations/cleaning/set"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&set.Operation{}, set.New)
	_ = operation.Register(&ifnotin.Operation{}, ifnotin.New)
	_ = operation.Register(&ifin.Operation{}, ifin.New)
	_ = operation.Register(&concatMap.Operation{}, concatMap.New)
	_ = operation.Register(&oneHotEncoding.Operation{}, oneHotEncoding.New)

}
