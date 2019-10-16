module main

go 1.12

require (
	github.com/project-flogo/catalystml-flogo/action v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/common v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/math v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/restructuring v0.0.0
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/activity/syncaction v0.0.0
	github.com/project-flogo/core v0.9.4-beta.2
	github.com/project-flogo/legacybridge v0.9.1
	github.com/project-flogo/ml v0.1.2
	github.com/project-flogo/stream v0.2.0
	github.com/project-flogo/stream/activity/aggregate v0.2.0
	github.com/skothari-tibco/csvtimer v0.9.1
)

replace github.com/project-flogo/contrib/activity/syncaction => github.com/skothari-tibco/contrib/activity/syncaction v0.0.0-20190916140408-6170cf1c5b6a

replace github.com/project-flogo/catalystml-flogo/operations/restructuring => ../../operations/restructuring

replace github.com/project-flogo/catalystml-flogo/operations/cleaning => ../../operations/cleaning

replace github.com/project-flogo/catalystml-flogo/operations/common => ../../operations/common

replace github.com/project-flogo/catalystml-flogo/operations/math => ../../operations/math

replace github.com/project-flogo/catalystml-flogo/action => ../../action
