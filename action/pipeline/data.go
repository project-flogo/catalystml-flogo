package pipeline

// Input of the CML Pipeline.
type PipelineInput struct {
	Type      string      `json:"type"`
	Dimension int         `json:"dimension"`
	Shape     []int       `json:"shape"`
	Label     interface{} `json:"label"`
}

// Output of CML Pipeline
type PipelineOutput struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// Config of Task.
type TaskConfig struct {
	Operation string      `json:"operation,required"`
	Params    interface{} `json:"params,omitempty"`
	Input     interface{} `json:"input,omitempty"`
	Output    interface{} `json:"output,required"`
}
