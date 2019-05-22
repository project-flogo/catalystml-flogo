package fps

type Settings struct {
	Model  *ModelInfo `json:"model"`
	FpsURI string     `md:"fpsURI,required"`
}
