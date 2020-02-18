package pipeline

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
)

type Task interface {
	Eval(scope data.Scope, logger log.Logger) (data.Scope, error)
	Name() string
	Position()
}
