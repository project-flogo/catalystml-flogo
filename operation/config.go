package operation

type Config struct {
	Id     string                 `json:"id"`
	Ref    string                 `json:"ref"`
	Params map[string]interface{} `json:"params,omitempty"`
	Input  map[string]interface{} `json:"input,omitempty"`
	Output map[string]interface{} `json:"output,omitempty"`
}
