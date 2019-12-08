package interpolateMissing

import (
	"errors"
	//"fmt"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/common"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger     log.Logger
	params     *Params
	method     InterpolateMethod
	edgeMethod InterpolateMethod
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	var method InterpolateMethod
	switch p.How {
	case "linear":
		method = Linear{}
	case "mean":
		method = Mean{}
	default:
		method = Mean{}
	}

	var edgeMethod InterpolateMethod
	switch p.Edges {
	case "linear":
		edgeMethod = Linear{}
	}

	return &Operation{
		params:     p,
		logger:     ctx.Logger(),
		method:     method,
		edgeMethod: edgeMethod,
	}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	dataFrame, isDataFrame := in.Data.(*common.DataFrame)
	if !isDataFrame {
		errors.New("Input data should be DataFrame type.")
	}

	a.logger.Info("Starting Operation interpolateMissing.")
	a.logger.Debug("Input Data of Operation interpolateMissing = ", dataFrame)
	a.logger.Debug("How = ", a.params.How)
	a.logger.Debug("Edges = ", a.params.Edges)

	columnName, isString := in.Col.(string)
	if !isString {
		for _, columnName := range dataFrame.GetLabels() {
			a.interpolateMissing(dataFrame.GetColumn(columnName))
		}
	} else {
		a.logger.Debug("Input Col Operation interpolateMissing = ", columnName)
		a.interpolateMissing(dataFrame.GetColumn(columnName))
	}

	a.logger.Info("Operation interpolateMissing Completed")
	a.logger.Debug("Output of Operation interpolateMissing = ", dataFrame)

	return dataFrame.AsIs(), nil
}

func (a *Operation) interpolateMissing(
	array []interface{},
) {

	inInterval := false
	gaps := make([]Interval, 0)
	edges := make([]Interval, 0)
	var gap Interval
	gapIndex := 0
	for index, value := range array {
		wasInInterval := inInterval
		a.logger.Debug("index = ", index, ", value = ", value)
		if nil == value {
			if 0 == index || nil != array[index-1] {
				inInterval = true
			}
		} else {
			if 0 != index && nil == array[index-1] {
				inInterval = false
			}
		}

		if !wasInInterval && inInterval {
			gap = NewInterval()
			gap.SetStartBound(index - 1)
		} else if wasInInterval && !inInterval {
			gap.SetEndBound(index)
			if -1 == gap.GetStartBound() || -1 == gap.GetEndBound() {
				edges = append(edges, gap)

			} else {
				gaps = append(gaps, gap)
			}
			gapIndex++
		}

		a.logger.Debug("in gap : ", inInterval)
	}

	if inInterval {
		edges = append(edges, gap)
	}
	a.logger.Debug("gaps : ", gaps)
	a.logger.Debug("edges : ", edges)

	for _, gap := range gaps {
		gap.Interpolate(array, a.method)
	}

	if nil != a.edgeMethod {
		for _, edge := range edges {
			edge.InterpolateEdge(array, a.edgeMethod)
		}
	}
}

func NewInterval() Interval {
	return Interval{stratBound: -1, endBound: -1}
}

type Interval struct {
	stratBound int
	endBound   int
}

func (i *Interval) SetStartBound(startBound int) {
	i.stratBound = startBound
}

func (i *Interval) GetStartBound() int {
	return i.stratBound
}

func (i *Interval) SetEndBound(endBound int) {
	i.endBound = endBound
}

func (i *Interval) GetEndBound() int {
	return i.endBound
}

func (i *Interval) Interpolate(series []interface{}, how InterpolateMethod) {
	how.Process(series, i)
}

func (i *Interval) InterpolateEdge(series []interface{}, how InterpolateMethod) {
	how.ProcessEdge(series, i)
}

type InterpolateMethod interface {
	Process(series []interface{}, gap *Interval)
	ProcessEdge(series []interface{}, gap *Interval)
}

type Mean struct {
}

func (m Mean) Process(series []interface{}, gap *Interval) {
	mean := (series[gap.GetStartBound()].(float64) + series[gap.GetEndBound()].(float64)) / 2
	for i := gap.GetStartBound() + 1; i < gap.GetEndBound(); i++ {
		series[i] = mean
	}
}

func (m Mean) ProcessEdge(series []interface{}, gap *Interval) {
}

type Linear struct {
}

func (l Linear) Process(series []interface{}, gap *Interval) {
	gapLength := gap.GetEndBound() - gap.GetStartBound()
	delta := slope(float64(gapLength), series[gap.GetStartBound()].(float64), series[gap.GetEndBound()].(float64))
	for i := gap.GetStartBound() + 1; i < gap.GetEndBound(); i++ {
		series[i] = series[i-1].(float64) + delta
	}
}

func (l Linear) ProcessEdge(series []interface{}, gap *Interval) {
	if -1 == gap.GetStartBound() {
		gapLength := gap.GetEndBound()
		//fmt.Println(
		//	"gap : ", *gap,
		//	", len(*series) : ", len(series),
		//	", (*series)[gapLength] = ", series[gapLength],
		//	", (*series)[gapLength+1] = ", series[gapLength+1],
		//)
		if (gapLength+2) <= len(series) && nil != series[gapLength] && nil != series[gapLength+1] {
			delta := slope(1.0, series[gapLength].(float64), series[gapLength+1].(float64))
			for i := gapLength - 1; i >= 0; i-- {
				series[i] = series[i+1].(float64) - delta
			}
		}
	} else if -1 == gap.GetEndBound() {
		gapLength := len(series) - gap.GetStartBound()
		//fmt.Println(
		//	"gap : ", *gap,
		//	", len(*series) : ", len(series),
		//	", (*series)[gapLength] = ", series[gapLength],
		//	", (*series)[gapLength+1] = ", series[gapLength+1],
		//)
		if (gapLength+2) <= len(series) && nil != series[gap.GetStartBound()] && nil != series[gap.GetStartBound()-1] {
			delta := slope(1.0, series[gap.GetStartBound()-1].(float64), series[gap.GetStartBound()].(float64))
			for i := gap.GetStartBound() + 1; i < len(series); i++ {
				series[i] = series[i-1].(float64) + delta
			}
		}
	}
}

func slope(gapLen float64, y1 float64, y2 float64) float64 {
	return (y2 - y1) / gapLen
}
