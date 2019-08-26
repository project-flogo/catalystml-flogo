package string-processing

import (
	"github.com/project-flogo/cml/operations/string-processing/phonenumber"
	"github.com/project-flogo/cml/operations/string-processing/geoencoding"
	"github.com/project-flogo/cml/operations/string-processing/split"
	"github.com/project-flogo/cml/action/operation"
)
	
func init() {
	_ = operation.Register(&phonenumber.Operation{})
	_ = operation.Register(&geoencoding.Operation{}, geoencoding.New)
	_ = operation.Register(&split.Operation{}, split.New)
}
