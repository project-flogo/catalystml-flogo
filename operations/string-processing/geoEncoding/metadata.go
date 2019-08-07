package geoencoding

import "github.com/project-flogo/core/data/coerce"

type Params struct {
	ApiKey string `md:"apiKey"`
}

type Input struct {
	Address string `md:"address"`
}

type Output struct {
	GeoLocation interface{} `md:"geoLocation"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"address": i.Address,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Address, err = coerce.ToString(values["address"])
	return err
}
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"geoLocation": o.GeoLocation,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	o.GeoLocation = values["geoLocation"]
	
	return nil
}
