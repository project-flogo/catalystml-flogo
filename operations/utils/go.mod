module github.com/project-flogo/catalystml-flogo/operations/utils

go 1.12

require (
	github.com/project-flogo/catalystml-flogo/action v0.1.6
	github.com/project-flogo/catalystml-flogo/operations/common v0.1.6-hf.1
	github.com/project-flogo/catalystml-flogo/operations/image_processing v0.1.6-hf.1
	github.com/project-flogo/catalystml-flogo/operations/nlp v0.1.6-hf.1
	github.com/project-flogo/catalystml-flogo/operations/string_processing v0.1.6-hf.1
	github.com/project-flogo/core v0.9.4
	github.com/project-flogo/operation/math v0.0.0
	github.com/stretchr/testify v1.4.0
)

replace github.com/project-flogo/operation/math => ../../action/example/operation