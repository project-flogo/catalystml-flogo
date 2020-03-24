module github.com/project-flogo/catalystml-flogo/action

require (
	github.com/project-flogo/catalystml-flogo/operations/cleaning v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/common v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/math v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/nlp v0.0.0-20200218153131-0a290fa79171
	github.com/project-flogo/catalystml-flogo/operations/restructuring v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/string_processing v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/core v0.9.4
	github.com/stretchr/testify v1.4.0

)

replace github.com/project-flogo/catalystml-flogo/operations/cleaning => ../operations/cleaning
replace github.com/project-flogo/catalystml-flogo/operations/common => ../operations/common
//replace github.com/project-flogo/catalystml-flogo/operations/image_processing => ../operations/image_processing
replace github.com/project-flogo/catalystml-flogo/operations/math => ../operations/math
replace github.com/project-flogo/catalystml-flogo/operations/nlp => ../operations/nlp
replace github.com/project-flogo/catalystml-flogo/operations/restructuring => ../operations/restructuring
replace github.com/project-flogo/catalystml-flogo/operations/string_processing => ../operations/string_processing
