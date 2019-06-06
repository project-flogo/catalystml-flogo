package fps

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Model         *ModelInfo `json:"model"`
	CatalystMlURI string     `md:"catalystMlURI,required"`
}

type Input struct {
	Input map[string]interface{} `md:"input"`
}

type Output struct {
	Output map[string]interface{} `md:"output"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"input": i.Input,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.Input, err = coerce.ToObject(values["input"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Output, err = coerce.ToObject(values["output"])
	if err != nil {
		return err
	}

	return nil
}
