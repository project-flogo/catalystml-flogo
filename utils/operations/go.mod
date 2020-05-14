module github.com/project-flogo/catalystml-flogo/utils/operation

go 1.12

require (
	github.com/project-flogo/catalystml-flogo/action v0.1.6
	github.com/project-flogo/catalystml-flogo/operations/nlp v0.1.6 // indirect
	github.com/project-flogo/core v1.0.0
	github.com/project-flogo/operation/math v0.0.0
	github.com/stretchr/testify v1.4.0
)

replace github.com/project-flogo/operation/math => ../../action/example/operation
