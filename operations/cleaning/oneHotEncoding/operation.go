package oneHotEncoding

import (
	"strconv"
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	p := &Params{}
	// default for separateOut is False, which is zero value of boolean

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}
	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	in := &Input{}

	in.FromMap(inputs)

	a.logger.Info("Starting One Hot Encoding operation.")
	a.logger.Debug("Input of One Hot Encoding operation.", in.Data)   //Make Debug
	a.logger.Debug("Params of One Hot Encoding operation.", a.params) //Make Debug

	m := in.Data.(map[string]interface{})

	var l int
	params, err := coerce.ToArray(a.params.InputColumns)
	if err != nil {
		return nil, err
	}
	for _, val := range params {
		arr, err := coerce.ToArray(m[val.(string)])
		if err != nil {
			return nil, err
		}
		l = len(arr)

		temp := removeDuplicate(arr)
		for _, newKey := range temp {
			var result []interface{}

			for _, data := range arr {
				if newKey == data {
					result = append(result, 1)
				} else {
					result = append(result, 0)
				}
			}

			//Checking for collisions - looping until no collision
			nk, _ := coerce.ToString(newKey)
			for isColInMap(m, nk) {
				nk = avoidCol(nk)
			}
			m[nk] = result
		}
		if a.params.KeepOrig == false {
			delete(m, val.(string))
		}
	}

	for _, col := range a.params.OutputColumns {
		if _, ok := m[col.(string)]; !ok {

			var tmparr []interface{}
			for i := 0; i < l; i++ {
				tmparr = append(tmparr, 0.)
			}
			m[col.(string)] = tmparr
		}
	}

	a.logger.Info("One Hot Encoding operation Completed.")
	a.logger.Debug("Output of hot encoding.", m) //Make Debug

	return m, nil
}

func removeDuplicate(arr []interface{}) (result []interface{}) {
	tempMap := make(map[string]interface{})

	for _, val := range arr {

		v, _ := coerce.ToString(val)
		if _, ok := tempMap[v]; !ok {
			tempMap[v] = "1"
		}
	}
	for key, _ := range tempMap {
		result = append(result, key)
	}
	return result
}

func avoidCol(in string) (out string) {
	sep := "_"
	var num int
	a := strings.Split(in, sep)

	if v, err := strconv.Atoi(a[len(a)-1]); err == nil {
		num = v + 1
	} else {
		num = 1
	}

	return a[0] + sep + strconv.Itoa(num)
}

func isColInMap(m map[string]interface{}, col string) bool {
	_, ok := m[col]
	return ok
}
