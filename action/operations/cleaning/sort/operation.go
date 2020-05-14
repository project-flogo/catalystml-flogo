package sort

import (
	"errors"
	"sort"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/operations/common"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	sortByKey bool
	params    *Params
	logger    log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	if "first" != operation.params.NilPosition {
		operation.params.NilPosition = "last"
	}

	switch operation.params.By[0].(type) {
	case string:
		operation.sortByKey = true
	case int:
		operation.sortByKey = false
	}

	dataFrame, ok := in.Data.(*common.DataFrame)

	operation.logger.Info("Starting Operation Sort.")
	operation.logger.Debug("DataFrame is...", in.Data)
	operation.logger.Debug("Sort By...", operation.params.By)
	operation.logger.Debug("Axis is...", operation.params.Axis)
	operation.logger.Debug("Ascending is...", operation.params.Ascending)
	operation.logger.Debug("NilPosition is...", operation.params.NilPosition)
	operation.logger.Debug("sortByKey is...", operation.sortByKey)

	if !ok {
		errors.New("Input data should be DataFrame type.")
	}

	var result *common.DataFrame
	sorter := common.NewDataFrameSorter(
		operation.params.Axis,
		operation.params.Ascending,
		"last" == operation.params.NilPosition,
		operation.sortByKey,
		operation.params.By,
		dataFrame,
	)

	operation.logger.Debug("Before sort : ", sorter)
	sort.Sort(sorter)
	operation.logger.Debug("After sorted : ", sorter)

	result = sorter.GetDataFrame()

	operation.logger.Info("Operation Sort completed.")
	operation.logger.Debug("The output of Operation Sort, As DataFrame : ", result)

	return result.AsIs(), err
}
