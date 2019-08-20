package image_processing

import (
	"github.com/project-flogo/cml/action/operation"
	gs "github.com/project-flogo/cml/operations/image_processing/gray_scale"
	"github.com/project-flogo/cml/operations/image_processing/img2tensor"
	"github.com/project-flogo/cml/operations/image_processing/resize"
)

func init() {
	_ = operation.Register(&gs.Operation{}, gs.New)
	_ = operation.Register(&img2tensor.Operation{}, img2tensor.New)
	_ = operation.Register(&resize.Operation{}, resize.New)
}
