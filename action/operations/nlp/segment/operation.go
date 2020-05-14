package segment

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
	"gopkg.in/jdkato/prose.v2"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Starting operation Segment.")
	a.logger.Debug("Input for Operation Segment.", input.Str)

	doc, err := prose.NewDocument(input.Str)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, sent := range doc.Sentences() {
		out = append(out, sent.Text)
	}
	a.logger.Info("Operation Tokenize Segment.")
	a.logger.Debug("Output of Operation Segment.", out)

	return out, nil
}
