package types

import (
	"bytes"
	"fmt"
	"image"

	"github.com/project-flogo/core/data"
)

func ValidateType(t string, val interface{}) error {

	// If the defined type is image
	//Check if you can decode the image.
	if t == "image" {
		_, _, err := image.Decode(bytes.NewReader(val.([]byte)))
		if err != nil {
			return fmt.Errorf("Type mismatch in input. Error in inferning image type %v", err)
		}
		return nil
	}

	definedType, _ := data.ToTypeEnum(t)
	givenType, _ := data.GetType(val)
	if definedType != givenType {
		return fmt.Errorf("Type mismatch in input. Defined type [%s] passed type [%s]", definedType, givenType)
	}

	return nil
}
