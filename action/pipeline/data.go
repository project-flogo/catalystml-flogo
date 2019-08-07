package pipeline

type PipelineInput struct {
	Type      string `json:"type"`
	Dimension int    `json:"dimension"`
	Shape     []int  `json:"shape"`
	Label     string `json:"label"`
}

type PipelineOutput struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
