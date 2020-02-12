package action

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/pipeline"
	_ "github.com/project-flogo/catalystml-flogo/operations/cleaning"
	_ "github.com/project-flogo/catalystml-flogo/operations/common"
	_ "github.com/project-flogo/catalystml-flogo/operations/image_processing"
	_ "github.com/project-flogo/catalystml-flogo/operations/math"
	_ "github.com/project-flogo/catalystml-flogo/operations/nlp"
	_ "github.com/project-flogo/catalystml-flogo/operations/restructuring"
	_ "github.com/project-flogo/catalystml-flogo/operations/string_processing"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/app/resource"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	action.Register(&Action{}, &ActionFactory{})

}

var manager *pipeline.Manager
var actionMd = action.ToMetadata(&Settings{}, &Input{}, &Output{})

type Action struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	definition  *pipeline.Definition
	inst        *pipeline.Instance
}

type ActionFactory struct {
	resManager *resource.Manager
}

func (f *ActionFactory) Initialize(ctx action.InitContext) error {
	f.resManager = ctx.ResourceManager()

	if manager != nil {
		return nil
	}

	pipeline.DefaultManager = f.resManager
	var err error

	if resource.GetLoader("cam") == nil {
		err = resource.RegisterLoader("cam", pipeline.NewResourceLoader(nil, pipeline.GetDataResolver()))
	}

	return err

}

// New returns the Action from the config.
func (f *ActionFactory) New(config *action.Config) (action.Action, error) {

	settings := &Settings{}
	err := metadata.MapToStruct(config.Settings, settings, true)
	if err != nil {
		return nil, err
	}

	catalystMlAction := &Action{}

	if settings.CatalystMlURI == "" {
		return nil, fmt.Errorf("pipeline URI not specified")
	}

	if strings.HasPrefix(settings.CatalystMlURI, resource.UriScheme) {
		res := f.resManager.GetResource(settings.CatalystMlURI)

		if res != nil {
			def, ok := res.Object().(*pipeline.Definition)
			if !ok {
				return nil, errors.New("unable to resolve fps: " + settings.CatalystMlURI)
			}
			catalystMlAction.definition = def
		} else {
			return nil, errors.New("unable to resolve fps in else: " + settings.CatalystMlURI)
		}
	} else {
		manager = pipeline.NewManager()

		def, err := manager.GetPipeline(settings.CatalystMlURI)
		if err != nil {
			return nil, err
		} else {
			if def == nil {
				return nil, errors.New("unable to resolve fps : " + settings.CatalystMlURI)
			}
		}
		catalystMlAction.definition = def
	}

	instId := ""

	if log.CtxLoggingEnabled() {
		//instLogger = log.ChildLoggerWithFields(logger, log.String("pipelineName", fpsAction.definition.Name()), log.String("pipelineId", instId))
	}

	//note: single pipeline instance for the moment
	inst := pipeline.NewInstance(catalystMlAction.definition, instId, log.ChildLogger(log.RootLogger(), "fps-logger"))
	catalystMlAction.inst = inst

	return catalystMlAction, nil
}

func (f *Action) Info() *action.Info {

	return nil
}

func (f *Action) Metadata() *action.Metadata {
	return actionMd
}

func (f *Action) IOMetadata() *metadata.IOMetadata {

	return f.definition.MetaData()
}

// Run runs the instance of the CML pipeline.
func (f *Action) Run(context context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	retData, err := f.inst.Run(inputs)

	if err != nil {
		return nil, err
	}

	return retData, nil
}
