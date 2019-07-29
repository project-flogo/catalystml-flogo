package operation

// Activity is an interface for defining a custom Activity Execution
type Operation interface {
	//Metadata() *Metadata

	Eval(inputs map[string]interface{}) (interface{}, error)
}

type Factory func(ctx InitContext) (Operation, error)
