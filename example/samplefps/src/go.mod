module main

go 1.12

require (
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/rest v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.2
	github.com/project-flogo/flow v0.9.1
	github.com/project-flogo/fps v0.0.0
	github.com/project-flogo/operation/math v0.0.0
	github.com/skothari-tibco/csvtimer v0.9.1

)

replace github.com/project-flogo/fps => ../../../

replace github.com/project-flogo/operation/math => ../../operation
