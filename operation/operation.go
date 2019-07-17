package operation

import (
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/support/log"
)

// Activity is an interface for defining a custom Activity Execution
type Operation interface {

	// Metadata returns the metadata of the activity
	Metadata() *Metadata

	// Eval is called when an Activity is being evaluated.  Returning true indicates
	// that the task is done.
	Eval(ctx Context) error
}

type Factory func(ctx InitContext) (Operation, error)

type InitContext interface {

	// Params
	Params() map[string]interface{}

	// MapperFactory gets the mapper factory associated with the operation host
	MapperFactory() mapper.Factory

	// Logger logger to using during initialization, operation implementations should not
	// keep a reference to this
	Logger() log.Logger
}
