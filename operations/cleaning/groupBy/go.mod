module github.com/project-flogo/cml/operations/cleaning/groupBy

go 1.12

require (
	github.com/project-flogo/cml/action v0.0.0-00010101000000-000000000000
	github.com/project-flogo/cml/operations/common v0.0.0-00010101000000-000000000000
	github.com/project-flogo/core v0.9.3
	github.com/stretchr/testify v1.4.0
)

replace github.com/project-flogo/cml/action => ../../../action

replace github.com/project-flogo/cml/operations/common => ../../common
