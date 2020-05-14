package transpose

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {

	i.Data = values["data"]
	return nil

}
