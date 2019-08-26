package removeStop

import (
	// "fmt"

	"github.com/bbalet/stopwords"
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)
	if p.Lang == "" {
		p.Lang = "en"
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

	// I am changing the word parser to keep punctuation
	// 		I just added \pP to the default regex
	stopwords.OverwriteWordSegmenter(`[\pL\p{Mc}\p{Mn}\pP-_']+`)

	a.logger.Debug("Executing operation...", input.Str)
	a.logger.Info("Removing stopwords...", input.Str)
	out := stopwords.CleanString(input.Str, a.params.Lang, false)

	a.logger.Info("String without stopwords: ", out)

	return out, nil
}
