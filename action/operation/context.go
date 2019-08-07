package operation

import (
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/support/log"
)

type InitContext interface {

	// Params
	Params() map[string]interface{}

	// MapperFactory gets the mapper factory associated with the operation host
	MapperFactory() mapper.Factory

	// Logger logger to using during initialization, operation implementations should not
	// keep a reference to this
	Logger() log.Logger
}
