package lag

import (
	"github.com/project-flogo/catalystml-flogo/operations/common"
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Table  interface{} `md:"table"`
	Lagnum int         `md:"lagnum"`
	Col    interface{} `md:"col"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	switch v := values["table"].(type) {
	case []interface{}:
		i.Table, err = common.ToInterfaceArray(v)
	case map[string]interface{}:
		i.Table, err = coerce.ToObject(v)
	}
	if err != nil {

		return err
	}
	i.Lagnum, _ = coerce.ToInt(values["lagnum"])
	i.Col, _ = values["col"]

	return err
}
