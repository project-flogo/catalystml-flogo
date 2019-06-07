package pipeline

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
)

type Status int

const (
	// StatusNotStarted indicates that the Pipeline has not started
	StatusNotStarted Status = 0

	// StatusActive indicates that the Pipeline is active
	StatusActive Status = 100

	// StatusDone indicates that the Pipeline is done
	StatusDone Status = 500
)

type ExecutionStatus int

const (
	// ExecStatusNotStarted indicates that the Pipeline execution has not started
	ExecStatusNotStarted ExecutionStatus = 0

	// ExecStatusActive indicates that the Pipeline execution is active
	ExecStatusActive ExecutionStatus = 100

	// ExecStatusStalled indicates that the Pipeline execution has stalled
	ExecStatusStalled ExecutionStatus = 400

	// ExecStatusCompleted indicates that the Pipeline execution has been completed
	ExecStatusCompleted ExecutionStatus = 500

	// ExecStatusCancelled indicates that the Pipeline execution has been cancelled
	ExecStatusCancelled ExecutionStatus = 600

	// ExecStatusFailed indicates that the Pipeline execution has failed
	ExecStatusFailed ExecutionStatus = 700
)

const (
	bitIsTimer  uint8 = 1
	bitIsTicker uint8 = 2
)

type ExecutionContext struct {
	pipeline *Instance

	stageId int
	status  ExecutionStatus

	pipelineInput  map[string]interface{}
	pipelineOutput map[string]interface{}

	currentInput  map[string]interface{}
	currentOutput map[string]interface{}

	updateTimers uint8
}

func (eCtx *ExecutionContext) ID() string {
	return eCtx.pipeline.id
}

func (eCtx *ExecutionContext) Name() string {
	return eCtx.pipeline.def.name
}

func (eCtx *ExecutionContext) currentStage() *Stage {
	//possibly keep pointer to state in ctx?
	return eCtx.pipeline.def.stages[eCtx.stageId]
}

func (eCtx *ExecutionContext) GetSetting(setting string) (value interface{}, exists bool) {
	stage := eCtx.currentStage()
	attr, found := stage.params[setting]
	if found {
		return attr, true
	}

	return nil, false
}

func (eCtx *ExecutionContext) GetInput(name string) interface{} {

	value, found := eCtx.currentInput[name]
	if found {
		return value
	}

	return nil
}

func (eCtx *ExecutionContext) GetOutput(name string) interface{} {
	value, found := eCtx.currentOutput[name]
	if found {
		return value
	}

	return nil
}

func (eCtx *ExecutionContext) GetInputObject(input data.StructValue) error {
	err := input.FromMap(eCtx.currentInput)
	return err
}

func (eCtx *ExecutionContext) Logger() log.Logger {
	return eCtx.pipeline.logger
}

func (eCtx *ExecutionContext) SetOutput(name string, value interface{}) error {

	if eCtx.currentOutput == nil {
		eCtx.currentOutput = make(map[string]interface{})
	}

	//todo coerce to type based on metadata
	eCtx.currentOutput[name] = value

	return nil
}

func (eCtx *ExecutionContext) SetOutputObject(output data.StructValue) error {
	eCtx.currentOutput = output.ToMap()
	return nil
}
