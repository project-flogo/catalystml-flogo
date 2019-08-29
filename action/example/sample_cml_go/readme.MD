# Run Actions using Go-API.

This short tutorial walks through how to run Flogo Actions using the Go-API provided by the flogo core.
First Import the action you need to run .

```go 
import (
    ...
    "github.com/project-flogo/core/api"
    "github.com/project-flogo/flow"
    ...
)
```

Then we need to initalize the Flogo app using Go-API

```go
app := api.NewApp()
```

Set the action settings such as `flowURI` in case of Flow; `pipelineURI`, `groupBy`, `outputChannel` in case of Stream and others depending on your action .
```go

actionSettings := make(map[string]interface{})

actionSettings["flowURI"] = "file://action.json"
```

Then initalize the action using  `NewIndependentAction` function on the app.

```go
flowAction , err := app.NewIndependentAction(&flow.Action{},actionSettings)

```

To run the action we call `RunAction` function on `api`. This needs to be called in the ActionHandler of the Trigger. 

```go

output, err := api.RunAction(ctx, flowAction,inputsForAction)

```

Full Example.
```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/project-flogo/contrib/trigger/rest"
	
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
    "github.com/project-flogo/flow"
    _ "github.com/project-flogo/contrib/activity/log"
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


	flowAct, _ := app.NewIndependentAction(&flow.Action{}, map[string]interface{}{"flowURI": "file://sampleflow.json"})
	
	actions =  map[string]action.Action{"flowAction":flowAct}

	return app
}

var actions map[string]action.Action

func RunActions(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	
	trgOut := &rest.Output{}
	
	trgOut.FromMap(inputs)
	
	
	out, err := api.RunAction(ctx,actions["flowAction"],inputs)
	
	if err != nil {
		return nil, err
	}

	return out, nil
}

```