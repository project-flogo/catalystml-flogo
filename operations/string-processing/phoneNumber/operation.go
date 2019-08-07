package phoneNumber

import (
	"errors"
	"regexp"

	"github.com/project-flogo/cml/action/operation"
)

func init() {
	_ = operation.Register(&Operation{})
}

type Operation struct {
}

func (a *Operation) Eval(ctx operation.Context) error {
	in := &Input{}

	ctx.GetInputObject(in)

	if in == nil {
		return errors.New("Input is not defined")
	}

	re := regexp.MustCompile(`[0-9]+`)

	//submatch := re.FindStringSubmatch(in.Text)
	match := re.FindAllString(in.Text, -1)

	if len(match) < 0 {
		// No match found
		ctx.Logger().Info("No match found")
		return nil
	}

	out := &Output{}

	out.Number = match[0]

	ctx.Logger().Info("Setting Output...", out)

	ctx.SetOutputObject(out)

	return nil
}
