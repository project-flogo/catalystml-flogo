package pipeline

import (
	"errors"
	"strconv"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/log"
)

var mf mapper.Factory
var resolver resolve.CompositeResolver
var position int

// NewTask returns a Task based on the TaskConfig.
func NewTask(config TaskConfig, mf mapper.Factory, resolver resolve.CompositeResolver) (Task, error) {
	mf = mf
	resolver = resolver
	taskImp := TaskImpl{}

	/*
		Need to check the type of params and inputs. Params can either be nil, single value or array.
		For each of the case, inputs can be either nil, single value or array. The stages are initlaized
		accordingly and added to the Task.
	*/
	if config.Params == nil {

		if config.Input == nil {

			stage, err := getStageWithInputObject(config.Operation, nil, nil, config.Output)

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

			stage, err := getStageWithInputObject(config.Operation, nil, config.Input, config.Output)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)

			return taskImp, nil

		}
		if inputType.String() == "array" {

			stages, err := getStagesWithInputArray(config.Operation, nil, config.Input, config.Output, false)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stages...)

			return taskImp, nil
		}

	}

	if config.Input == nil {
		stage, err := getStageWithInputObject(config.Operation, nil, nil, config.Output)

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

			stage, err := getStageWithInputObject(config.Operation, config.Params, config.Input, config.Output)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stage)

			return taskImp, nil

		} else if inputType.String() == "array" {

			stages, err := getStagesWithInputArray(config.Operation, config.Params, config.Input, config.Output, false)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stages...)

			return taskImp, nil

		}

	}

	if paramType.String() == "array" {
		inputType, err := data.GetType(config.Input)

		if err != nil {
			return nil, err
		}

		if inputType.String() == "object" {
			for key, _ := range config.Params.([]interface{}) {

				stage, err := getStageWithInputObject(config.Operation, config.Params.([]interface{})[key], config.Input, config.Output)

				if err != nil {
					return nil, err
				}

				taskImp.stages = append(taskImp.stages, stage)
			}

			return taskImp, nil

		} else if inputType.String() == "array" {

			stages, err := getStagesWithInputArray(config.Operation, config.Params, config.Input, config.Output, true)

			if err != nil {
				return nil, err
			}

			taskImp.stages = append(taskImp.stages, stages...)

			return taskImp, nil

		}

	}

	return nil, nil
}

type TaskImpl struct {
	stages []*Stage
}

func (t TaskImpl) Position() {
	position = position + 1
}

// Eval runs the Task. A Task is composed of single or multiple stages.
func (t TaskImpl) Eval(scope data.Scope, logger log.Logger) (data.Scope, error) {
	currentInput := make(map[string]interface{})
	var err error
	var stageOutput interface{}
	for key, stage := range t.stages {

		logger.Debugf("Operation Input Mapper for stage [%v]: [%v]", stage.name+"-"+strconv.Itoa(position)+"-"+strconv.Itoa(key), stage.inputMapper)

		if stage.inputMapper != nil {

			currentInput, err = stage.inputMapper.Apply(scope)
			if err != nil {
				return nil, err
			}

		}

		logger.Debugf("Starting operation [%v] with inputs: [%v]", stage.name+"-"+strconv.Itoa(position)+"-"+strconv.Itoa(key), currentInput)

		stageOutput, err = stage.opt.Eval(currentInput)

		if err != nil {
			return nil, err
		}

		scope.SetValue(stage.output, stageOutput)
		logger.Debugf("Setting output for [%v] operation outputs: [%v] ", stage.name+"-"+strconv.Itoa(position)+"-"+strconv.Itoa(key), stageOutput)

		_, err = stage.outputMapper.Apply(scope)

		logger.Debugf("Scope after operation [%v] : [%v]", stage.name+"-"+strconv.Itoa(position)+"-"+strconv.Itoa(key), scope)

		if err != nil {
			return nil, err
		}

	}

	return scope, nil

}

func (t TaskImpl) Name() string {
	return t.stages[0].name
}

// GetStageWithInputObject returns the stages when the input is object.
func getStageWithInputObject(config string, params interface{}, inputs interface{}, output interface{}) (*Stage, error) {

	stageConfig := &StageConfig{}

	cInput, err := coerce.ToObject(inputs)
	if err != nil {
		return nil, err
	}

	if params == nil {

		stageConfig.Config = &operation.Config{Operation: config, Output: output.(string), Input: cInput}

	} else {
		stageConfig.Config = &operation.Config{Operation: config, Params: params.(map[string]interface{}), Output: output.(string), Input: cInput}
	}

	stage, err := NewStage(stageConfig, mf, resolver)

	if err != nil {
		return nil, err
	}
	return stage, nil
}

// GetStagesWithInputArray returns the stages when the input is array.
func getStagesWithInputArray(config string, params interface{}, inputs interface{}, output interface{}, isParamsArray bool) ([]*Stage, error) {

	var stages []*Stage

	in, err := coerce.ToArray(inputs)

	if err != nil {
		return nil, err
	}
	out, err := coerce.ToArray(output)

	if err != nil {
		return nil, err
	}

	if len(in) != len(out) {
		return nil, errors.New("Mismatch number of Inputs and Outputs")
	}

	for key, _ := range inputs.([]interface{}) {
		stageConfig := &StageConfig{}

		if isParamsArray {
			stageConfig.Config = &operation.Config{Operation: config, Params: params.([]interface{})[key].(map[string]interface{}), Input: in[key].(map[string]interface{}), Output: out[key].(string)}
		} else {
			if params != nil {
				stageConfig.Config = &operation.Config{Operation: config, Params: params.(map[string]interface{}), Input: in[key].(map[string]interface{}), Output: out[key].(string)}
			} else {
				stageConfig.Config = &operation.Config{Operation: config, Input: in[key].(map[string]interface{}), Output: out[key].(string)}
			}

		}

		stage, err := NewStage(stageConfig, mf, resolver)

		if err != nil {
			return nil, err
		}

		stages = append(stages, stage)
	}

	return stages, nil

}
