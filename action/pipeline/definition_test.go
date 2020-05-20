package pipeline

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/stretchr/testify/assert"
)

func init() {
	operation.Register(&test.TestOperation{}, test.TestNew)
}

var def = `{
	"input": [
    	{
			"type": "array",
        	"label":"dataIn"
    	}
	],
	"structure": [
    {
      "operation":"test",
      "input":{
        "data":"$dataIn"
      },
      "output":"datatemp"
    },
    {
      "operation":"test",
      "input":{
        "data":"$datatemp['Age']",
        "value":100,
        "minvalue":18
      },
      "output":"datatemp['Age']"
	}
	]
}
`

func TestSimple(t *testing.T) {

	var defConfig *DefinitionConfig

	err := json.Unmarshal([]byte(def), &defConfig)
	assert.Nil(t, err)

	pDef, err := NewDefinition(defConfig, mapper.NewFactory(resolve.GetBasicResolver()), resolve.GetBasicResolver())
	assert.Nil(t, err)
	assert.NotNil(t, pDef)
}

var multipleInput = `{
	"input": [
    	{
			"type": "array",
        	"label":"dataIn"
    	}
	],
	"structure": [
    {
      "operation":"test",
      "input":[
		  {
        	"data":"$dataIn"
		  }, 
		  {
			"data":"$dataIn"
		  }
	  ],
      "output":["datatemp","datatemp2"]
    },
    {
      "operation":"test",
      "input":{
        "data":"$datatemp['Age']",
        "value":100,
        "minvalue":18
      },
      "output":"datatemp['Age']"
	}
	]
}
`

func TestMultipleInput(t *testing.T) {

	var defConfig *DefinitionConfig

	err := json.Unmarshal([]byte(multipleInput), &defConfig)
	assert.Nil(t, err)

	pDef, err := NewDefinition(defConfig, mapper.NewFactory(resolve.GetBasicResolver()), resolve.GetBasicResolver())
	assert.Nil(t, err)
	assert.NotNil(t, pDef)
}

var multipleParams = `{
	"input": [
    	{
			"type": "array",
        	"label":"dataIn"
    	}
	],
	"structure": [
    {
      "operation":"test",
      "input":[
		  {
        	"data":"$dataIn"
		  }, 
		  {
			"data":"$dataIn"
		  }
	  ],
      "output":["datatemp","datatemp2"]
    },
    {
      "operation":"test",
      "input":{
        "data":"$datatemp['Age']",
        "value":100,
        "minvalue":18
      },
      "output":"datatemp['Age']"
	},
	{
		"operation":"test",
		"input":{"data":"$datatemp"},
		"params":[
			{
		  	"col": ["bank","cardType","entry_type","transaction"],
		  	"separateOut":false
			},
			{
				"col": ["bank","cardType","entry_type","transaction"],
				"separateOut":false
			}
		],
		"output":"datatemp"
	  }
	]
}
`

func TestMultipleParams(t *testing.T) {

	var defConfig *DefinitionConfig

	err := json.Unmarshal([]byte(multipleParams), &defConfig)
	assert.Nil(t, err)

	pDef, err := NewDefinition(defConfig, mapper.NewFactory(resolve.GetBasicResolver()), resolve.GetBasicResolver())
	assert.Nil(t, err)
	assert.NotNil(t, pDef)
}

var multipleParamsInput = `{
	"input": [
    	{
			"type": "array",
        	"label":"dataIn"
    	}
	],
	"structure": [
    {
      "operation":"test",
      "input":[
		  {
        	"data":"$dataIn"
		  }, 
		  {
			"data":"$dataIn"
		  }
	  ],
      "output":["datatemp","datatemp2"]
    },
    {
      "operation":"test",
      "input":{
        "data":"$datatemp['Age']",
        "value":100,
        "minvalue":18
      },
      "output":"datatemp['Age']"
	},
	{
		"operation":"test",
		"input":[{"data":"$datatemp"}, {"data": "$datatemp"} ],
		"params":[
			{
		  	"col": ["bank","cardType","entry_type","transaction"],
		  	"separateOut":false
			},
			{
				"col": ["bank","cardType","entry_type","transaction"],
				"separateOut":false
			}
		],
		"output":["datatemp","datatemp2"]
	  }
	]
}
`

func TestMultipleParamsInput(t *testing.T) {

	var defConfig *DefinitionConfig

	err := json.Unmarshal([]byte(multipleParamsInput), &defConfig)
	assert.Nil(t, err)

	pDef, err := NewDefinition(defConfig, mapper.NewFactory(resolve.GetBasicResolver()), resolve.GetBasicResolver())
	assert.Nil(t, err)
	assert.NotNil(t, pDef)
}

var multipleMisMatchInputOutput = `{
	"input": [
    	{
			"type": "array",
        	"label":"dataIn"
    	}
	],
	"structure": [
    {
      "operation":"test",
      "input":[
		  {
        	"data":"$dataIn"
		  }, 
		  {
			"data":"$dataIn"
		  }
	  ],
      "output":["datatemp","datatemp2"]
    },
    {
      "operation":"test",
      "input":{
        "data":"$datatemp['Age']",
        "value":100,
        "minvalue":18
      },
      "output":"datatemp['Age']"
	},
	{
		"operation":"test",
		"input":[{"data":"$datatemp"}, {"data": "$datatemp"} ],
		"params":[
			{
		  	"col": ["bank","cardType","entry_type","transaction"],
		  	"separateOut":false
			},
			{
				"col": ["bank","cardType","entry_type","transaction"],
				"separateOut":false
			}
		],
		"output":"datatemp"
	  }
	]
}
`

func TestMultipleMisMatchInputOutput(t *testing.T) {

	var defConfig *DefinitionConfig

	err := json.Unmarshal([]byte(multipleMisMatchInputOutput), &defConfig)
	assert.Nil(t, err)

	pDef, err := NewDefinition(defConfig, mapper.NewFactory(resolve.GetBasicResolver()), resolve.GetBasicResolver())
	assert.NotNil(t, err)
	assert.Nil(t, pDef)
}
