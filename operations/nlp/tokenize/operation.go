package tokenize

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
	"gopkg.in/jdkato/prose.v2"
)

type Operation struct {
	// params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Debug("Executing operation...", input.Str)

	doc, err := prose.NewDocument(input.Str)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, tok := range doc.Tokens() {
		out = append(out, tok.Text)
		// fmt.Println(tok.Text, tok.Tag)
	}
	a.logger.Info("Resulting array...", out)

	return out, nil
}
