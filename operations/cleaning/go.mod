module github.com/project-flogo/cml/operations/cleaning

require (
	github.com/project-flogo/cml/action v0.0.0
	github.com/project-flogo/core v0.9.0-rc.1
	github.com/stretchr/testify v1.3.0
	github.com/project-flogo/cml/operations/common v0.0.0
	
)

replace github.com/project-flogo/cml/action => ../../action

replace github.com/project-flogo/cml/operations/common => ../common