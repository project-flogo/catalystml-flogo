package addCol2Table

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

type Params struct {
	// Axis int `md:"axis"`
}

type Input struct {
	Matrix interface{} `md:"matrix"`
	Col    interface{} `md:"col"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	var m [][]interface{}
	tmpMat, err := coerce.ToArray(values["matrix"])
	if err != nil {
		return fmt.Errorf("unable to coerce matrix with ToArray: %s", err)
	}
	for _, row := range tmpMat {
		var n []interface{}
		tmpRow, err := coerce.ToArray(row)
		if err != nil {
			return fmt.Errorf("unable to coerce matrix with ToArray: %s", err)
		}
		for _, item := range tmpRow {
			n = append(n, item)
		}
		m = append(m, n)
	}
	i.Matrix = m

	tmpCol, _ := coerce.ToArray(values["col"])
	if err != nil {
		return fmt.Errorf("unable to coerce col with ToArray: %s", err)
	}
	i.Col = tmpCol
	return err
}
