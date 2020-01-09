package api

import (
	"context"

	"github.com/project-flogo/core/action"
)

type Option func(*action.Config)

func NewAction(option ...Option) (action.Action, error) {

	factory := action.GetFactory("github.com/project-flogo/catalystml-flogo/action")

	aConfig := new(action.Config)

	for _, opt := range option {
		opt(aConfig)
	}

	act, err := factory.New(aConfig)

	if err != nil {
		return nil, err
	}

	return act, nil

}

func SetURISettings(path string) Option {
	return func(a *action.Config) {
		a.Settings = make(map[string]interface{})
		a.Settings["catalystMlURI"] = path
	}
}

func Run(act action.Action, inputs map[string]interface{}) (map[string]interface{}, error) {

	return act.(action.SyncAction).Run(context.Background(), inputs)

}
