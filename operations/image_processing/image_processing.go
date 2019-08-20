package image_processing

import (
	"github.com/project-flogo/cml/action/operation"
	gs "github.com/project-flogo/cml/operations/image_processing/gray_scale"
)

func init() {
	_ = operation.Register(&gs.Operation{}, gs.New)
}
