module main

go 1.12

require (
	github.com/jung-kurt/gofpdf v1.15.1 // indirect
	github.com/project-flogo/catalystml-flogo/action v0.0.0
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0-20191112152814-dc69f9aae0fd // indirect
	github.com/project-flogo/catalystml-flogo/operations/image_processing v0.0.0-20191112152814-dc69f9aae0fd // indirect
	github.com/project-flogo/catalystml-flogo/operations/math v0.0.0-20191112152814-dc69f9aae0fd // indirect
	github.com/project-flogo/catalystml-flogo/operations/restructuring v0.0.0-20191112152814-dc69f9aae0fd // indirect
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/rest v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.4
	github.com/project-flogo/flow v0.9.4
	github.com/project-flogo/operation/math v0.0.0
	github.com/skothari-tibco/csvtimer v0.9.1
	golang.org/x/crypto v0.0.0-20191111213947-16651526fdb4 // indirect
	golang.org/x/exp v0.0.0-20191030013958-a1ab85dbe136 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	golang.org/x/mobile v0.0.0-20191031020345-0945064e013a // indirect
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5 // indirect
	golang.org/x/sys v0.0.0-20191110163157-d32e6e3b99c4 // indirect
	golang.org/x/tools v0.0.0-20191112005509-a3f652f18032 // indirect
	gonum.org/v1/netlib v0.0.0-20191031114514-eccb95939662 // indirect
	gonum.org/v1/plot v0.0.0-20191107103940-ca91d9d40d0a // indirect
)

replace github.com/project-flogo/catalystml-flogo/action => ../../../

replace github.com/project-flogo/operation/math => ../../operation

replace github.com/project-flogo/catalystml-flogo/operations/string_processing => ../../../../operations/string_processing

replace github.com/project-flogo/catalystml-flogo/operations/common => ../../../../operations/common

replace github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0-20191112152814-dc69f9aae0fd => ../../../../operations/cleaning
