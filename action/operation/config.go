package operation

// The struct config of operation.
type Config struct {
	Operation string                 `json:"operation"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Input     map[string]interface{} `json:"input,omitempty"`
	Output    string                 `json:"output,required"`
}
