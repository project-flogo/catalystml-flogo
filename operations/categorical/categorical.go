package categorical

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/categorical/oneHotEncoding"
	
)

func init() {
	_ = operation.Register(&oneHotEncoding.Operation{}, oneHotEncoding.New)
}
