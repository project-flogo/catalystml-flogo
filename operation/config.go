package operation

type Config struct {
	Id        string                 `json:"id"`
	Operation string                 `json:"operation"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Input     map[string]interface{} `json:"input,omitempty"`
	Output    string                 `json:"output,required"`
}
