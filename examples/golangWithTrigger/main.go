package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/project-flogo/contrib/trigger/timer"
	"github.com/skothari-tibco/csvtimer"

	cml "github.com/project-flogo/catalystml-flogo/action"

	_ "github.com/project-flogo/catalystml-flogo/operations/restructuring"
	_ "github.com/project-flogo/catalystml-flogo/operations/cleaning"
	_ "github.com/project-flogo/catalystml-flogo/operations/common"
	_ "github.com/project-flogo/catalystml-flogo/operations/math"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
)

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	f, err := os.OpenFile("text.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("ATM, Age, AmEx, Amount, Bank0, Bank1, Bank10, Bank2, Bank3, Bank4, Bank5, Bank6, Bank7, Bank8, Bank9, INTERNET, Master, Master_Debit, POS, Visa, Visa_Debit, chip, gender, magnetic, v1, v2, v3, v4, v3*v4*Master, v3*v4*Visa, fraud\n"); err != nil {
		log.Println(err)
	}

	trg := app.NewTrigger(&csvtimer.Trigger{}, nil)
	h, _ := trg.NewHandler(&csvtimer.HandlerSettings{RepeatInterval: "50", Header: true, FilePath: "init_data.csv"})

	h.NewAction(RunActions)

	cmlAct, _ := app.NewIndependentAction(&cml.Action{}, map[string]interface{}{"catalystMlURI": "file://cml.json"})

	actions = map[string]action.Action{"cml": cmlAct}
	return app
}

var actions map[string]action.Action

func RunActions(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	trgOut := &csvtimer.Output{}

	trgOut.FromMap(inputs)
	// fmt.Println(inputs)
	data := inputs["data"]
	// fmt.Println(data)

	// return nil, nil
	// inputs["paragraph"] = "Natural Language Processing (NLP) is all about leveraging tools, techniques and algorithms to process and understand natural language-based data, which is usually unstructured like text, speech and so on. In this series of articles, we will be looking at tried and tested strategies, techniques and workflows which can be leveraged by practitioners and data scientists to extract useful insights from text data. We will also cover some useful and interesting use-cases for NLP. This article will be all about processing and understanding text data with tutorials and hands-on examples."
	inputs["dataIn"] = []interface{}{data}

	out, err := api.RunAction(ctx, actions["cml"], inputs)
	// out, err := api.RunAction(ctx, actions["cml2"], inputs)

	d := out["output"].(map[string]interface{})
	// fmt.Println(d)
	// var arr []interface{}
	arr := []string{"ATM", "Age", "AmEx", "Amount", "Bank0", "Bank1", "Bank10", "Bank2", "Bank3", "Bank4", "Bank5", "Bank6", "Bank7", "Bank8", "Bank9", "INTERNET", "Master", "Master_Debit", "POS", "Visa", "Visa_Debit", "chip", "gender", "magnetic", "v1", "v2", "v3", "v4", "v3*v4*Master", "v3*v4*Visa", "fraud"}

	var mtx [][]interface{}
	for i := 0; i < len(d["gender"].([]interface{})); i++ {
		var a []interface{}
		for _, val := range arr {
			if d[val] != nil {
				a = append(a, d[val].([]interface{})[i])
			} else {
				a = append(a, []interface{}{0.}...)
			}
		}
		mtx = append(mtx, a)
	}
	// fmt.Println(mtx)
	s := ""
	for i := 0; i < len(mtx[0]); i++ {
		var sval string
		switch v := mtx[0][i].(type) {
		case int:
			sval = fmt.Sprintf("%d", v)
		case float32:
			sval = fmt.Sprintf("%f", v)
		case float64:
			sval = fmt.Sprintf("%f", v)

		}
		// fmt.Println("sval ", sval)
		var st string
		if i != 0 {
			st = "," + sval
		} else {
			st = sval
		}
		s += st
	}
	s += "\n"

	f, err := os.OpenFile("text.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(s); err != nil {
		log.Println(err)
	}

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"code": 200, "data": out}, nil
}
