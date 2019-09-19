package main

import (
	"context"
	"fmt"
	"os"

	"github.com/project-flogo/contrib/trigger/rest"
	
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
	_ "github.com/project-flogo/operation/math"
	cml "github.com/project-flogo/catalystml-flogo/action"
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

	trg := app.NewTrigger(&rest.Trigger{}, &rest.Settings{Port: 8080})
	h, _ := trg.NewHandler(&rest.HandlerSettings{Method: "GET", Path: "/blah/:num"})
	
	h.NewAction(RunActions)


	cmlAct, _ := app.NewIndependentAction(&cml.Action{}, map[string]interface{}{"catalystMlURI": "file://../samplefps/samplecml.json"})
	
	actions =  map[string]action.Action{"cml":cmlAct}

	return app
}

var actions map[string]action.Action

func RunActions(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	
	trgOut := &rest.Output{}
	
	trgOut.FromMap(inputs)
	
	inputs["input"] = "3"
	
	out, err := api.RunAction(ctx,actions["cml"],inputs)
	
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"code":200, "data":out["feat1"]}, nil
}