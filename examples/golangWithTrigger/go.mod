module main

go 1.12

require (
	github.com/kr/pretty v0.2.0 // indirect
	github.com/project-flogo/catalystml-flogo/action v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/common v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/math v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/restructuring v0.0.0
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/activity/syncaction v0.0.0
	github.com/project-flogo/core v0.10.1
	github.com/project-flogo/legacybridge v0.9.1
	github.com/project-flogo/ml v0.1.2
	github.com/project-flogo/stream v0.2.0
	github.com/project-flogo/stream/activity/aggregate v0.2.0
	github.com/rogpeppe/go-internal v1.5.2 // indirect
	github.com/skothari-tibco/csvtimer v0.9.1
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200204104054-c9f3fb736b72 // indirect
	golang.org/x/lint v0.0.0-20200130185559-910be7a94367 // indirect
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5 // indirect
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
