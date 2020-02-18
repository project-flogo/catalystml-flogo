module main

go 1.12

require (
	github.com/project-flogo/catalystml-flogo/action v0.0.0-20200205173424-7f21c14cfdb1
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing v0.0.0-20200205173424-7f21c14cfdb1
	github.com/project-flogo/catalystml-flogo/operations/nlp v0.0.0-20191016133650-3f39d84d181e
	github.com/project-flogo/catalystml-flogo/operations/string_processing v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.4
	github.com/project-flogo/flow v0.9.1

)

replace github.com/project-flogo/catalystml-flogo/operations/image_processing => ../../operations/image_processing

replace github.com/project-flogo/catalystml-flogo/action v0.0.0-20200205173424-7f21c14cfdb1 => ../../action

replace github.com/project-flogo/catalystml-flogo/operations/nlp v0.0.0-20191016133650-3f39d84d181e => ../../operations/nlp
