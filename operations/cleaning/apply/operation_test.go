package apply

import (
	"fmt"
	"testing"

	_ "github.com/project-flogo/catalystml-flogo/operations/count"
	"github.com/project-flogo/fps/support/test"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["str"] = "<p>The hotspot is here</p>"

	p := Params{MapOrArray: "map"}

	pmap := make(map[string]interface{})
	pmap["blah"] = 1
	pmap["grrr"] = 2
	imap := make(map[string]interface{})
	imap["s0"] = "moo moo moo baa maa neigh neigh"
	imap["s1"] = "$item"

	// data := []string{"moo", "baa", "maa", "neigh"}
	data := make(map[interface{}]interface{})
	data["blah"] = "moo"
	data["grrr"] = "neigh"
	data["aho"] = "maa"
	data["sigh"] = "baa"
	fn := Config{Operation: "count", Params: pmap, Input: imap}

	inputs["data"] = data
	inputs["func"] = fn

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	output, err := opt.Eval(inputs)
	fmt.Println(output)
	assert.NotNil(t, output)
	// assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}

func TestArray(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["str"] = "<p>The hotspot is here</p>"

	p := Params{MapOrArray: "array"}

	pmap := make(map[string]interface{})
	pmap["blah"] = 1
	pmap["grrr"] = 2
	imap := make(map[string]interface{})
	imap["s0"] = "moo moo moo baa maa neigh neigh"
	imap["s1"] = "$item"

	data := []string{"moo", "baa", "maa", "neigh"}
	fn := Config{Operation: "count", Params: pmap, Input: imap}

	inputs["data"] = data
	inputs["func"] = fn

	optInitConext := test.NewOperationInitContext(p, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)

	mcomp := make(map[string]interface{})
	mcomp["moo"] = 3
	mcomp["neigh"] = 2
	mcomp["baa"] = 1
	mcomp["maa"] = 1

	output, err := opt.Eval(inputs)
	fmt.Println(output)
	// fmt.Println(output, mcomp, reflect.DeepEqual(output, mcomp))
	assert.NotNil(t, output)
	// assert.True(t, reflect.DeepEqual(output, mcomp))
	assert.Nil(t, err)

}
