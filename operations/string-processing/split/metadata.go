package split

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Text      string `md:"text"`
	Separator string `md:"separator"`
}

type Params struct {
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"text":      i.Text,
		"separator": i.Separator,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Text, err = coerce.ToString(values["text"])
	i.Separator, err = coerce.ToString(values["separator"])
	return err
}
