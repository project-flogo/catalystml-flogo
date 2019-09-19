package image_processing

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/img2tensor"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/resize"
)

func init() {
	_ = operation.Register(&img2tensor.Operation{}, img2tensor.New)
	_ = operation.Register(&resize.Operation{}, resize.New)
}
