package sort

import (
	"errors"
	"sort"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/common"

	//	"github.com/project-flogo/core/data/coerce"
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
	operation.logger.Info("DataFrame is...", in.Data)
	operation.logger.Info("Sort By...", operation.params.By)
	operation.logger.Info("Axis is...", operation.params.Axis)
	operation.logger.Info("Ascending is...", operation.params.Ascending)
	operation.logger.Info("NilPosition is...", operation.params.NilPosition)
	operation.logger.Info("sortByKey is...", operation.sortByKey)

	if !ok {
		errors.New("Input data should be DataFrame type.")
	}

	var result interface{}
	switch operation.params.Axis {
	case 0:
		result, err = operation.sortByCol(dataFrame)
	case 1:
		result, err = operation.sortByRow(dataFrame)
	}

	if nil != err {
		return nil, err
	}

	operation.logger.Info("Operation Sort completed.")
	operation.logger.Info("The output of Operation Sort.", result)

	return result, err
}

func (operation *Operation) sortByRow(dataFrame *common.DataFrame) (*common.DataFrame, error) {
	tuples := common.TupleSorter{
		Ascending: operation.params.Ascending,
		NilLast:   "last" == operation.params.NilPosition,
		ByKey:     operation.sortByKey,
		SortBy:    operation.params.By,
		Tuples:    make([]common.SortableTuple, 0),
	}

	for _, key := range dataFrame.GetKeys() {
		data := append(dataFrame.GetColumn(key), key)
		tuples.Tuples = append(tuples.Tuples, common.SortableTuple{
			Data: data,
		})
	}

	sort.Sort(tuples)

	newDataFrame := common.NewDataFrame()
	for _, sTuple := range tuples.Tuples {
		newDataFrame.AddColumn(sTuple.Data[len(sTuple.Data)-1].(string), sTuple.Data[:len(sTuple.Data)-1])
	}

	return newDataFrame, nil
}

func (operation *Operation) sortByCol(dataFrame *common.DataFrame) (*common.DataFrame, error) {
	tuples := common.TupleSorter{
		Ascending: operation.params.Ascending,
		NilLast:   "last" == operation.params.NilPosition,
		ByKey:     operation.sortByKey,
		SortBy:    operation.params.By,
		Tuples:    make([]common.SortableTuple, 0),
	}

	keys := dataFrame.GetKeys()

	newDataFrame, _ := common.ProcessDataFrame(dataFrame, func(tuple map[string]interface{}, newDataFrame *common.DataFrame, lastTuple bool) error {

		tuples.Tuples = append(tuples.Tuples, common.NewSortableTuple(tuple, keys))

		if lastTuple {
			sort.Sort(tuples)
			for _, sTuple := range tuples.Tuples {
				newTuple := make(map[string]interface{})

				for key, index := range sTuple.KeyToIndex {
					newTuple[key] = sTuple.Data[index]
				}
				common.TupleAppendToDataframe(newTuple, newDataFrame)
			}
		}
		return nil
	})

	return newDataFrame, nil
}
