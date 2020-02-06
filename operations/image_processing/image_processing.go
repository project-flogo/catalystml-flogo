package image_processing

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/grayscale"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/img2tensor"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/resize"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/subsectiontoimage"
	"github.com/project-flogo/catalystml-flogo/operations/image_processing/tensor2image"
)

func init() {
	_ = operation.Register(&img2tensor.Operation{}, img2tensor.New)
	_ = operation.Register(&resize.Operation{}, resize.New)
	_ = operation.Register(&grayscale.Operation{}, grayscale.New)
	_ = operation.Register(&subsectiontoimage.Operation{}, subsectiontoimage.New)
	_ = operation.Register(&tensor2image.Operation{}, tensor2image.New)
}
