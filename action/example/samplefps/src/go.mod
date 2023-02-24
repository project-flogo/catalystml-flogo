module main

go 1.12

require (
	github.com/kr/pretty v0.2.0 // indirect
	github.com/project-flogo/catalystml-flogo/action v0.2.0
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/rest v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v1.0.0
	github.com/project-flogo/flow v0.9.4
	github.com/project-flogo/operation/math v0.0.0
	go.uber.org/zap v1.12.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace github.com/project-flogo/catalystml-flogo/action => ../../../

replace github.com/project-flogo/operation/math => ../../operation

replace github.com/project-flogo/catalystml-flogo/operations/string_processing => ../../../../operations/string_processing

replace github.com/project-flogo/catalystml-flogo/operations/common => ../../../../operations/common

replace github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0-20191112152814-dc69f9aae0fd => ../../../../operations/cleaning
