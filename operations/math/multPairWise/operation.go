package multPairWise

import (
	"fmt"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {
	//To get the inputs in the desired types.
	input := &Input{}

	input.FromMap(inputs)

	a.logger.Info("Executing operation multPairWise ...")

	mtx0 := inputs["matrix0"].([]interface{})
	mtx1 := inputs["matrix1"].([]interface{})

	out,err:=mtxmultpairwise(mtx0,mtx1)
	if err!=nil{
		return nil,err
	}

	a.logger.Info("Output of multPairWise ", out)
	return out, nil
}

func mtxmultpairwise(mtx0 []interface{},mtx1 []interface{}) ([]interface{},error){
	//  Recursive function that either adds the two elements  
	//   or goes to the next level of the matrices

	if len(mtx0) != len(mtx1) {
		return nil, fmt.Errorf("matrices are not the same size")
	}

	var err error
	var mtxOut []interface{}

	for i:=0;i<len(mtx0); i++{
		switch v:=mtx0[i].(type){
		case []interface{}:
			var tmp []interface{}
			tmp,err=mtxmultpairwise(v,mtx1[i].([]interface{}))
			mtxOut=append(mtxOut,tmp)
		case int:
			mtxOut=append(mtxOut,v*mtx1[i].(int))
		case int32:
			mtxOut=append(mtxOut,mtx0[i].(int32)*mtx1[i].(int32))
		case int64:
			mtxOut=append(mtxOut,mtx0[i].(int64)*mtx1[i].(int64))
		case float32:
			mtxOut=append(mtxOut,mtx0[i].(float32)*mtx1[i].(float32))
		case float64:
			mtxOut=append(mtxOut,mtx0[i].(float64)*mtx1[i].(float64))
		}
	}
	return mtxOut,err
}