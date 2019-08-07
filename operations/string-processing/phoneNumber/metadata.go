package phoneNumber

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Text string `md:"text"`
}
type Output struct {
	Number string `md:"number"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"text": i.Text,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Text, err = coerce.ToString(values["text"])
	return err
}
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"number": o.Number,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Number, err = coerce.ToString(values["number"])
	return err
}
