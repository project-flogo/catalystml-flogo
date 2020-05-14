package stem

import (
	"github.com/kljensen/snowball"
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/reiver/go-porterstemmer"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)
	if p.Algo == "" {
		p.Algo = "Porter"
	}

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Starting Operation Stem.")
	a.logger.Debug("Input for Operation Stem.", input.Str)
	var out string

	if a.params.Algo == "Porter" {
		a.logger.Info("Starting Porter stemmer...")
		word := []rune(input.Str)
		stem := porterstemmer.Stem(word)
		out = string(stem)

	} else if a.params.Algo == "Snowball" {
		a.logger.Info("Starting Snowball stemmer...")
		stemmed, err := snowball.Stem(input.Str, "english", true)
		if err != nil {
			return nil, err
		}
		out = stemmed

	}
	a.logger.Info("Operation Stem Words completed.")
	a.logger.Debug("Output of Stem word operation. ", out)

	return out, nil
}
