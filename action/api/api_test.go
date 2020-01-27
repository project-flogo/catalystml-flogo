package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAction(t *testing.T) {

	cmlOpt := SetURISettings("file://../example/samplefps/samplecml.json")

	cmlAct, err := NewAction(cmlOpt)

	assert.NotNil(cmlAct)

	assert.Nil(err)

}
