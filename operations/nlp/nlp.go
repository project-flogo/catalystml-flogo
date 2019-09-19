package nlp

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/removeStop"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/stem"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/tokenize"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/getstopwords"
)

func init() {
	_ = operation.Register(&removeStop.Operation{}, removeStop.New)
	_ = operation.Register(&stem.Operation{}, stem.New)
	_ = operation.Register(&tokenize.Operation{}, tokenize.New)
	_ = operation.Register(&getstopwords.Operation{}, getstopwords.New)
}

