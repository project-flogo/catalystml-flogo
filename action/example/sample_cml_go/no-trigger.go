package main
import (
	"context"
	"fmt"
	
	"github.com/project-flogo/core/action"
	_ "github.com/project-flogo/operation/math"
	_ "github.com/project-flogo/catalystml-flogo/action"
	
)

func main() {
	
	factory := action.GetFactory("github.com/project-flogo/catalystml-flogo/action")

	var act action.Action

	act, _ = factory.New(&action.Config{Settings: map[string]interface{}{"catalystMlURI": "file://../samplefps/samplecml.json"}})
	
	out, _ := act.(action.SyncAction).Run(context.Background(),map[string]interface{}{"input": "3"})

	fmt.Println("Output...", out)
}