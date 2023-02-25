module main

go 1.12

require (
	github.com/project-flogo/catalystml-flogo/action v0.2.0
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/common v0.1.6
	github.com/project-flogo/catalystml-flogo/operations/math v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/restructuring v0.0.0
	github.com/project-flogo/core v1.0.0
	github.com/skothari-tibco/csvtimer v0.9.1
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/lint v0.0.0-20200130185559-910be7a94367 // indirect
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/tools v0.0.0-20200205141839-4abfd4a1628e // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/project-flogo/contrib/activity/syncaction => github.com/skothari-tibco/contrib/activity/syncaction v0.0.0-20190916140408-6170cf1c5b6a

replace github.com/project-flogo/catalystml-flogo/operations/restructuring => ../../operations/restructuring

replace github.com/project-flogo/catalystml-flogo/operations/cleaning => ../../operations/cleaning

replace github.com/project-flogo/catalystml-flogo/operations/common => ../../operations/common

replace github.com/project-flogo/catalystml-flogo/operations/math => ../../operations/math

replace github.com/project-flogo/catalystml-flogo/action => ../../action
