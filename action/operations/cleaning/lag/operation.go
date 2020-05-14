package lag

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	in := &Input{}
	err := in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	a.logger.Info("Starting Operation Lag.")
	a.logger.Debug("Input of Operation Lag.", in.Table, in.Lagnum, "  ", in.Col)

	var coldata []interface{}
	switch v := in.Table.(type) {
	case []interface{}:
		col, _ := coerce.ToInt(in.Col)
		coldata = v[col].([]interface{})
	case map[string]interface{}:
		col, _ := coerce.ToString(in.Col)
		coldata = v[col].([]interface{})
	}

	var vec []interface{}
	for i, _ := range coldata {
		var dat interface{}
		if i >= in.Lagnum && i < in.Lagnum+len(coldata) {
			dat = coldata[i-in.Lagnum]
		} else {
			dat = nil
		}
		vec = append(vec, dat)
	}

	switch v := in.Table.(type) {
	case []interface{}:
		in.Table = append(v, vec)

	case map[string]interface{}:
		c, _ := coerce.ToString(in.Col)
		newCol := fmt.Sprintf("%s_%d", c, in.Lagnum)
		in.Table = v
		in.Table.(map[string]interface{})[newCol] = vec
	}

	a.logger.Info("Operation Lag Completed.")
	a.logger.Debug("Output of Operation Lag.", in.Table)

	return in.Table, nil
}
