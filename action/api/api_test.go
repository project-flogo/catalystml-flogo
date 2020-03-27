package api

import (
	"testing"

	_ "github.com/project-flogo/catalystml-flogo/action"
	"github.com/stretchr/testify/assert"
)

func TestNewAction(t *testing.T) {
	cmlOpt := SetURISettings("file://../example/samplefps/samplecml.json")
	cmlAct, err := NewAction(cmlOpt)
	assert.Nil(t, cmlAct)
	// since math operation is not loaded yet.
	assert.NotNil(t, err)
}
