module main

go 1.12

require (
	github.com/project-flogo/cml/action v0.0.0
	github.com/project-flogo/cml/operations/cleaning v0.0.0
	github.com/project-flogo/cml/operations/nlp v0.0.0
	github.com/project-flogo/cml/operations/string_processing v0.0.0
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.2
	github.com/project-flogo/flow v0.9.1

)

replace github.com/project-flogo/cml/action => /Users/avanderg@tibco.com/working/go_working/cml/action

replace github.com/project-flogo/cml/operations/cleaning => /Users/avanderg@tibco.com/working/go_working/cml/operations/cleaning

replace github.com/project-flogo/cml/operations/nlp => /Users/avanderg@tibco.com/working/go_working/cml/operations/nlp

replace github.com/project-flogo/cml/operations/string_processing => /Users/avanderg@tibco.com/working/go_working/cml/operations/string_processing
