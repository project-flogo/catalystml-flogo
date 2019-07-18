package operation

/*
import (
	"encoding/json"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/metadata"
)

type Metadata struct {
	*metadata.IOMetadata
	Params map[string]data.TypedValue
}

func (md *Metadata) MarshalJSON() ([]byte, error) {
	var mdParams []*data.Attribute
	var mdInputs []*data.Attribute
	var mdOutputs []*data.Attribute

	for _, v := range md.Params {
		if attr, ok := v.(*data.Attribute); ok {
			mdParams = append(mdParams, attr)
		}
	}
	for _, v := range md.Input {
		if attr, ok := v.(*data.Attribute); ok {
			mdInputs = append(mdInputs, attr)
		}
	}
	for _, v := range md.Output {
		if attr, ok := v.(*data.Attribute); ok {
			mdOutputs = append(mdOutputs, attr)
		}
	}

	return json.Marshal(&struct {
		Params []*data.Attribute `json:"params,omitempty"`
		Input  []*data.Attribute `json:"input,omitempty"`
		Output []*data.Attribute `json:"output,omitempty"`
	}{
		Params: mdParams,
		Input:  mdInputs,
		Output: mdOutputs,
	})
}
func ToMetadata(mdStructs ...interface{}) *Metadata {

	var params map[string]data.TypedValue
	var input map[string]data.TypedValue
	var output map[string]data.TypedValue

	for _, mdStruct := range mdStructs {
		typedMap := metadata.StructToTypedMap(mdStruct)
		name := metadata.GetStructName(mdStruct)

		switch name {
		case "params":
			params = typedMap
		case "input":
			input = typedMap
		case "output":
			output = typedMap
		}
	}

	return &Metadata{Params: params, IOMetadata: &metadata.IOMetadata{Input: input, Output: output}}
}*/
