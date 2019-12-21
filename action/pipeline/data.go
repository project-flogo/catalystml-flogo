package pipeline

type PipelineInput struct {
	Type      string      `json:"type"`
	Dimension int         `json:"dimension"`
	Shape     []int       `json:"shape"`
	Label     interface{} `json:"label"`
}

type PipelineOutput struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TaskConfig struct {
	Operation string      `json:"operation,required"`
	Params    interface{} `json:"params,omitempty"`
	Input     interface{} `json:"input,omitempty"`
	Output    interface{} `json:"output,required"`
}
