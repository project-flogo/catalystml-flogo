package levenshteinDistance

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	// Axis int `md:"axis"`
}

type Input struct {
	S0 string `md:"s0"`
	S1 string `md:"s1"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	tmpS0, err := coerce.ToString(values["s0"])
	if err != nil {
		return fmt.Errorf("unable to coerce s0 with ToString: %s", err)
	}
	i.S0 = tmpS0

	tmpS1, err := coerce.ToString(values["s1"])
	if err != nil {
		return fmt.Errorf("unable to coerce s0 with ToString: %s", err)
	}
	i.S1 = tmpS1

	return err
}
