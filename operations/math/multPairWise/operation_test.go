package multPairWise
import (
	"testing"
	"github.com/project-flogo/core/support/log"
	
	"github.com/stretchr/testify/assert"
)
func TestSingleDEval(t *testing.T) {
	//params := Params{}
	opt := &Operation{logger : log.RootLogger() }
	
	inputs := make(map[string]interface{})
	inputs["matrix0"] = []interface{}{-2.7771, -0.012105382}
	inputs["matrix1"] = []interface{}{0,1}
	

	//input := &Input{Array: []float32{1.0, 2.0, 3.6, 4}}


	_, err := opt.Eval(inputs)

	assert.Nil(t, err)

}