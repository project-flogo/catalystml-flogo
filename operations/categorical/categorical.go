package categorical

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/categorical/oneHotEncoding"
)

func init() {
	_ = operation.Register(&oneHotEncoding.Operation{}, oneHotEncoding.New)
}
