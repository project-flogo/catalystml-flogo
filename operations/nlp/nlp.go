package nlp

import (
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/cml/operations/nlp/removeStop"
	"github.com/project-flogo/cml/operations/nlp/stem"
	"github.com/project-flogo/cml/operations/nlp/tokenize"
)

func init() {
	_ = operation.Register(&removeStop.Operation{}, removeStop.New)
	_ = operation.Register(&stem.Operation{}, stem.New)
	_ = operation.Register(&tokenize.Operation{}, tokenize.New)
}

