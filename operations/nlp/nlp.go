package nlp

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/removeStop"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/stem"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/tokenize"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/getstopwords"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/segment"
	"github.com/project-flogo/catalystml-flogo/operations/nlp/postag"
)

func init() {
	_ = operation.Register(&removeStop.Operation{}, removeStop.New)
	_ = operation.Register(&stem.Operation{}, stem.New)
	_ = operation.Register(&tokenize.Operation{}, tokenize.New)
	_ = operation.Register(&getstopwords.Operation{}, getstopwords.New)
	_ = operation.Register(&segment.Operation{}, segment.New)
	_ = operation.Register(&postag.Operation{}, postag.New)
}

