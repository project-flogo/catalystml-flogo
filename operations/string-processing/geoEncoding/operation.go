package geoencoding

import (
	"fmt"
	"errors"

	"github.com/codingsince1985/geo-golang/google"
	"github.com/project-flogo/core/data/metadata"
	"github.com/codingsince1985/geo-golang"
	"github.com/project-flogo/cml/action/operation"

)

type Operation struct {
	params *Params
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p}, nil
}

func (a *Operation) Eval(ctx operation.Context)  error {
	in := &Input{}

	ctx.GetInputObject(in)

	if in == nil {
		return true, errors.New("Input is not defined")
	}

	out := &Output{}

	
	geocoder := google.Geocoder(a.params.ApiKey)
	
	if geocoder == nil {
		return true, errors.New("Error in getting geocoder")
	}

	ctx.Logger().Info(try(geocoder, in.Address))

	ctx.Logger().Info("Setting Output...", out)

	ctx.SetOutputObject(out)

	return true, nil
}

 func try(geocoder geo.Geocoder, addr string) string {
	location, _ := geocoder.Geocode(addr)
	
	if location != nil {
		return fmt.Sprintf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} 
	
	return fmt.Sprintf("got <nil> location")
	

	
}