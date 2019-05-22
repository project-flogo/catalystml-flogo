package fps

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/app/resource"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/fps/pipeline"
)

func init() {
	action.Register(&FPSAction{}, &ActionFactory{})

}

var manager *pipeline.Manager
var actionMd = action.ToMetadata(&Settings{})
var logger log.Logger

type FPSAction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	ioMetadata  *metadata.IOMetadata
	definition  *pipeline.Definition
	inst        *pipeline.Instance
}

type ActionFactory struct {
	resManager *resource.Manager
}

func (f *ActionFactory) Initialize(ctx action.InitContext) error {
	f.resManager = ctx.ResourceManager()

	logger = log.ChildLogger(log.RootLogger(), "fps-logger")

	if manager != nil {
		return nil
	}

	mapperFactory := mapper.NewFactory(pipeline.GetDataResolver())
	//fmt.Println(mapperFactory)
	manager = pipeline.NewManager()
	err := resource.RegisterLoader("fps", pipeline.NewResourceLoader(mapperFactory, pipeline.GetDataResolver()))

	return err

}

func (f *ActionFactory) New(config *action.Config) (action.Action, error) {

	settings := &Settings{}
	err := metadata.MapToStruct(config.Settings, settings, true)
	if err != nil {
		return nil, err
	}

	fpsAction := &FPSAction{}

	if settings.FpsURI == "" {
		return nil, fmt.Errorf("pipeline URI not specified")
	}

	if strings.HasPrefix(settings.FpsURI, resource.UriScheme) {

		res := f.resManager.GetResource(settings.FpsURI)

		if res != nil {
			def, ok := res.Object().(*pipeline.Definition)
			if !ok {
				return nil, errors.New("unable to resolve fps: " + settings.FpsURI)
			}
			fpsAction.definition = def
		} else {
			return nil, errors.New("unable to resolve fps in else: " + settings.FpsURI)
		}
	} else {
		def, err := manager.GetPipeline(settings.FpsURI)
		if err != nil {
			return nil, err
		} else {
			if def == nil {
				return nil, errors.New("unable to resolve fps : " + settings.FpsURI)
			}
		}
		fpsAction.definition = def
	}

	fpsAction.ioMetadata = fpsAction.definition.Metadata()

	instId := ""

	instLogger := logger

	if log.CtxLoggingEnabled() {
		//instLogger = log.ChildLoggerWithFields(logger, log.String("pipelineName", fpsAction.definition.Name()), log.String("pipelineId", instId))
	}

	//note: single pipeline instance for the moment
	inst := pipeline.NewInstance(fpsAction.definition, instId, instLogger)
	fpsAction.inst = inst

	return fpsAction, nil
}

func (f *FPSAction) Info() *action.Info {
	fmt.Println("Implement me")
	return nil
}

func (f *FPSAction) Metadata() *action.Metadata {
	return actionMd
}

func (f *FPSAction) IOMetadata() *metadata.IOMetadata {
	return f.ioMetadata
}

func (f *FPSAction) Run(context context.Context, inputs map[string]interface{}, handler action.ResultHandler) error {

	go func() {

		defer handler.Done()
		retData, err := f.inst.Run(inputs)

		if err != nil {
			handler.HandleResult(nil, err)
		} else {
			handler.HandleResult(retData, err)
		}

	}()

	return nil
}
