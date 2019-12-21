package pipeline

import (
	"fmt"
	"strconv"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/log"
)

func NewTask(config TaskConfig, mf mapper.Factory, resolver resolve.CompositeResolver) (Task, error) {

	taskImp := TaskImpl{}

	if config.Params == nil {

		if config.Input == nil {

			stageConfig := &StageConfig{}

			stageConfig.Config = &operation.Config{Operation: config.Operation, Output: config.Output.(string)}

			stage, err := NewStage(stageConfig, mf, resolver)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)

			return taskImp, nil
		}

		inputType, err := data.GetType(config.Input)

		if err != nil {
			return nil, err
		}

		if inputType.String() == "object" {

			stageConfig := &StageConfig{}
			stageConfig.Config = &operation.Config{Operation: config.Operation, Input: config.Input.(map[string]interface{}), Output: config.Output.(string)}
			stage, err := NewStage(stageConfig, mf, resolver)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)

			return taskImp, nil

		}
		if inputType.String() == "array" {

			for key, _ := range config.Input.([]interface{}) {
				stageConfig := &StageConfig{}

				stageConfig.Config = &operation.Config{Operation: config.Operation, Input: config.Input.([]interface{})[key].(map[string]interface{}), Output: config.Output.([]interface{})[key].(string)}
				stage, err := NewStage(stageConfig, mf, resolver)

				if err != nil {
					return nil, err
				}

				taskImp.stages = append(taskImp.stages, stage)
			}
			fmt.Println("Len...", len(taskImp.stages))
			return taskImp, nil
		}

	}

	if config.Input == nil {

		stageConfig := &StageConfig{}

		stageConfig.Config = &operation.Config{Operation: config.Operation, Output: config.Output.(string)}

		stage, err := NewStage(stageConfig, mf, resolver)

		if err != nil {
			return nil, err
		}

		taskImp.stages = append(taskImp.stages, stage)

		return taskImp, nil
	}

	paramType, err := data.GetType(config.Params)

	if err != nil {
		return nil, err
	}

	if paramType.String() == "object" {

		inputType, err := data.GetType(config.Input)

		if err != nil {
			return nil, err
		}

		if inputType.String() == "object" {
			stageConfig := &StageConfig{}
			stageConfig.Config = &operation.Config{Operation: config.Operation, Params: config.Params.(map[string]interface{}), Input: config.Input.(map[string]interface{}), Output: config.Output.(string)}
			stage, err := NewStage(stageConfig, mf, resolver)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)
		} else if inputType.String() == "array" {
			for key, _ := range config.Input.([]interface{}) {
				stageConfig := &StageConfig{}

				stageConfig.Config = &operation.Config{Operation: config.Operation, Params: config.Params.(map[string]interface{}), Input: config.Input.([]interface{})[key].(map[string]interface{}), Output: config.Output.([]interface{})[key].(string)}
				stage, err := NewStage(stageConfig, mf, resolver)

				if err != nil {
					return nil, err
				}

				taskImp.stages = append(taskImp.stages, stage)
			}
		} else {
			stageConfig := &StageConfig{}

			stageConfig.Config = &operation.Config{Operation: config.Operation, Output: config.Output.(string)}
			stage, err := NewStage(stageConfig, mf, resolver)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)
		}

	} else if paramType.String() == "array" {
		inputType, err := data.GetType(config.Input)

		if err != nil {
			return nil, err
		}

		if inputType.String() == "object" {
			for key, _ := range config.Params.([]interface{}) {
				stageConfig := &StageConfig{}
				stageConfig.Config = &operation.Config{Operation: config.Operation, Params: config.Params.([]interface{})[key].(map[string]interface{}), Input: config.Input.(map[string]interface{}), Output: config.Output.(string)}
				stage, err := NewStage(stageConfig, mf, resolver)

				if err != nil {
					return nil, err
				}

				taskImp.stages = append(taskImp.stages, stage)
			}
		} else if inputType.String() == "array" {
			for key, _ := range config.Input.([]interface{}) {
				stageConfig := &StageConfig{}

				stageConfig.Config = &operation.Config{Operation: config.Operation, Params: config.Params.([]interface{})[key].(map[string]interface{}), Input: config.Input.([]interface{})[key].(map[string]interface{}), Output: config.Output.([]interface{})[key].(string)}
				stage, err := NewStage(stageConfig, mf, resolver)

				if err != nil {
					return nil, err
				}

				taskImp.stages = append(taskImp.stages, stage)
			}
		}
	}

	return taskImp, nil
}

type TaskImpl struct {
	stages []*Stage
}

func (t TaskImpl) Eval(scope data.Scope, logger log.Logger) (data.Scope, error) {
	currentInput := make(map[string]interface{})
	var err error

	for key, stage := range t.stages {

		logger.Debugf("Operation Input Mapper for stage [%v]: [%v]", stage.name+"-"+strconv.Itoa(key), stage.inputMapper)

		if stage.inputMapper != nil {

			currentInput, err = stage.inputMapper.Apply(scope)
			if err != nil {
				return nil, err
			}

		}

		logger.Debugf("Starting operation [%v] with inputs: [%v]", stage.name+"-"+strconv.Itoa(key), currentInput)
		stageOutput, err := stage.opt.Eval(currentInput)

		if err != nil {
			return nil, err
		}

		scope.SetValue(stage.output, stageOutput)
		logger.Debugf("Setting output for [%v] operation outputs: [%v] ", stage.name+"-"+strconv.Itoa(key), stageOutput)

		_, err = stage.outputMapper.Apply(scope)

		logger.Debugf("Scope after operation [%v] : [%v]", stage.name+"-"+strconv.Itoa(key), scope)

		if err != nil {
			return nil, err
		}

	}

	return scope, nil

}

func (t TaskImpl) Name() string {
	return t.stages[0].name
}
