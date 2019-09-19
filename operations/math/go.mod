module github.com/project-flogo/catalystml-flogo/operations/math

require (
	github.com/project-flogo/catalystml-flogo/action v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/common v0.0.0
	github.com/project-flogo/core v0.9.4-0.20190829220729-31eb91f2b8a7
	github.com/stretchr/testify v1.3.0
)

replace github.com/project-flogo/catalystml-flogo/action => ../../action

replace github.com/project-flogo/catalystml-flogo/operations/common => ../common
