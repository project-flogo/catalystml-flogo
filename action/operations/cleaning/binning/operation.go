package binning

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/operations/common"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
	params *Params
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	dataFrame, ok := in.Data.(*common.DataFrame)

	a.logger.Info("Starting Operation Binning.")
	a.logger.Debug("Input of Operation Binning.", dataFrame)
	a.logger.Debug("Quantile...", a.params.Quantile)
	a.logger.Debug("Bins...", a.params.Bins)
	a.logger.Debug("Column...", a.params.Column)
	a.logger.Debug("Labels...", a.params.Labels)
	a.logger.Debug("Retbins...", a.params.Retbins)
	a.logger.Debug("Precision...", a.params.Precision)
	a.logger.Debug("Retbins...", a.params.Retbins)
	a.logger.Debug("Duplicate...", a.params.Duplicates)

	if !ok {
		errors.New("Input data should be DataFrame type.")
	}

	sorter := common.NewDataFrameSorter(
		0,
		true,
		true,
		true,
		[]interface{}{a.params.Column},
		dataFrame,
	)

	a.logger.Debug("Before sort : ", sorter)
	sort.Sort(sorter)
	a.logger.Debug("After sorted : ", sorter)

	sortedDF := sorter.GetDataFrame()
	targetArr := sortedDF.GetColumn(a.params.Column)

	var result interface{}
	if 0 < a.params.Quantile {
		result = a.qBinning(
			targetArr,
			a.params.Quantile,
			"drop" == strings.ToLower(a.params.Duplicates),
			a.params.Retbins,
			math.Pow10(a.params.Precision),
			a.params.Labels,
		)
	} else {
		result = a.binning(
			targetArr,
			a.params.Bins,
			"drop" == strings.ToLower(a.params.Duplicates),
			a.params.Retbins,
			math.Pow10(a.params.Precision),
			a.params.Labels,
		)
	}

	a.logger.Info("Operation Binning Completed")
	a.logger.Debug("Output of Operation Binning.", result)

	return result, nil
}

func (a *Operation) binning(
	array []interface{},
	bounds []float64,
	dropDuplicate bool,
	retbins bool,
	roundingFactor float64,
	labels []string,
) map[string]interface{} {
	useLabel := false
	if len(labels) == len(bounds)-1 {
		useLabel = true
	}
	result := make(map[string]interface{})
	result["bounds"] = bounds
	previousValue := 0.0
	previousBinNumber := 0
	upperBound := bounds[1]
	counter := 0
	for index, data := range array {
		value := data.(float64)
		if value > upperBound {
			counter++
			upperBound = bounds[counter+1]
		}

		var binNumber int
		duplicate := false
		if value == previousValue {
			binNumber = previousBinNumber
			duplicate = true
		} else {
			previousBinNumber = counter
			binNumber = counter
		}

		a.logger.Debug(">> i = ", index, ", value = ", value, ", binNumber = ", binNumber, ", previousBinNumber = ", previousBinNumber)

		if retbins && (!dropDuplicate || !duplicate) {
			var bins map[string]interface{}
			if nil == result["bins"] {
				bins = make(map[string]interface{}, len(bounds))
				result["bins"] = bins
			} else {
				bins = result["bins"].(map[string]interface{})
			}

			var label string
			if useLabel {
				label = labels[binNumber]
			} else {
				label = strconv.Itoa(binNumber)
			}

			if nil == bins[label] {
				bins[label] = make([]interface{}, 0)
			}
			bins[label] = append(bins[label].([]interface{}), value)
			a.logger.Debug("i = ", index, ", label = ", label, ", value = ", value)
		}
		previousValue = value
	}
	return result
}

func (a *Operation) qBinning(
	array []interface{},
	quantile int,
	dropDuplicate bool,
	retbins bool,
	roundingFactor float64,
	labels []string,
) map[string]interface{} {
	useLabel := false
	if len(labels) == quantile {
		useLabel = true
	}
	result := make(map[string]interface{})
	bounds := make([]interface{}, quantile+1)
	result["bounds"] = bounds
	deltaPercentile := 100.0 / float64(quantile)
	lowerBound := 0.0
	upperBound := deltaPercentile
	previousValue := 0.0
	previousPcnt := 0.0
	previousBinNumber := 0
	counter := 0
	for index, data := range array {
		value := data.(float64)
		percentile := 100 * (float64(index) + 0.5) / float64(len(array))
		if percentile > upperBound {
			lowerBound += deltaPercentile
			upperBound += deltaPercentile
			a.logger.Debug(">> i = ", index+1, ", PreviousPcnt = ", previousPcnt, ", Percentile = ", percentile)
			a.logger.Debug(">> previousValue = ", previousValue, ", value = ", value)
			a.logger.Debug(">> lowerBound = ", lowerBound, ", upperBound = ", upperBound)
			counter++
			bound := (value*(lowerBound-previousPcnt) + previousValue*(percentile-lowerBound)) / (percentile - previousPcnt)
			bounds[counter] = math.Round(bound*roundingFactor) / roundingFactor
		} else if 0 == index {
			bounds[counter] = value
		} else if len(array)-1 == index {
			bounds[counter+1] = value
		}

		var binNumber int
		duplicate := false
		if value == previousValue {
			binNumber = previousBinNumber
			duplicate = true
		} else {
			previousBinNumber = counter
			binNumber = counter
		}

		if retbins && (!dropDuplicate || !duplicate) {
			var bins map[string]interface{}
			if nil == result["bins"] {
				bins = make(map[string]interface{}, quantile)
				result["bins"] = bins
			} else {
				bins = result["bins"].(map[string]interface{})
			}

			var label string
			if useLabel {
				label = labels[binNumber]
			} else {
				label = strconv.Itoa(binNumber)
			}

			if nil == bins[label] {
				bins[label] = make([]interface{}, 0)
			}
			bins[label] = append(bins[label].([]interface{}), value)
			a.logger.Debug("i = ", index, ", label = ", label, ", value = ", value)
		}
		previousValue = value
		previousPcnt = percentile
	}
	return result
}
