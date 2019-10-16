package groupBy

type Params struct {
	Index     []string            `md:"index"`
	Aggregate map[string][]string `md:"aggregate"`
	Level     int                 `md:"level"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = ToDataFrame(values["data"])

	return err
}

func ToDataFrame(val interface{}) (interface{}, error) {
	return val, nil
}
