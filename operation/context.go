package operation

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

// Context describes the execution context for an Operation.
// It provides access to attributes, task and Flow information.
type Context interface {
	// ActivityHost gets the "host" under with the activity is executing
	//OperationHost() Host

	//Name the name of the activity that is currently executing
	Name() string

	// GetInput gets the value of the specified input attribute
	GetInput(name string) interface{}

	// SetOutput sets the value of the specified output attribute
	SetOutput(name string, value interface{}) error

	// GetInputObject gets all the activity input as the specified object.
	GetInputObject(input data.StructValue) error

	// SetOutputObject sets the activity output as the specified object.
	SetOutputObject(output data.StructValue) error

	//TODO : do we need this?
	// GetSharedTempData get shared temporary data for activity, lifespan
	// of the data dependent on the activity host implementation
	//GetSharedTempData() map[string]interface{}

	// Logger the logger for the activity
	Logger() log.Logger
}

type Host interface {
	// ID returns the ID of the Activity Host
	ID() string

	// Name the name of the Activity Host
	Name() string

	// IOMetadata get the input/output metadata of the operation host
	IOMetadata() *metadata.IOMetadata

	//TODO : I'm not sure if we need this
	// Reply is used to reply to the activity Host with the results of the execution
	//Reply(replyData map[string]interface{}, err error)

	// Return is used to indicate to the activity Host that it should complete and return the results of the execution
	//Return(returnData map[string]interface{}, err error)

	// Scope returns the scope for the Host's data
	//Scope() data.Scope
}
