package api

import (
	"context"

	"github.com/project-flogo/core/action"
)

type Option func(*action.Config)

//Initiliaize the action.
func NewAction(option ...Option) (action.Action, error) {

	aConfig := new(action.Config)

	//Set the Config options for the action
	for _, opt := range option {
		opt(aConfig)
	}

	factory := action.GetFactory("github.com/project-flogo/catalystml-flogo/action")

	act, err := factory.New(aConfig)

	if err != nil {
		return nil, err
	}

	return act, nil

}

//Set `catalystMlURI` config of action
func SetURISettings(path string) Option {
	//Return Option.
	return func(a *action.Config) {
		a.Settings = make(map[string]interface{})
		a.Settings["catalystMlURI"] = path
	}
}

//Run Sync Action
func Run(act action.Action, inputs map[string]interface{}) (map[string]interface{}, error) {

	return act.(action.SyncAction).Run(context.Background(), inputs)

}
