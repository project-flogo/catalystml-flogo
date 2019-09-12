module github.com/project-flogo/cml/operations/nlp

require (
	github.com/bbalet/stopwords v1.0.0
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/kljensen/snowball v0.6.0
	github.com/mingrammer/commonregex v1.0.0 // indirect
	github.com/project-flogo/cml/action v0.0.0
	github.com/project-flogo/core v0.9.0-rc.1
	github.com/reiver/go-porterstemmer v1.0.1
	golang.org/x/text v0.3.2 // indirect
	gonum.org/v1/gonum v0.0.0-20190821101010-d61003946d0d // indirect
	gopkg.in/jdkato/prose.v2 v2.0.0-20190814032740-822d591a158c
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
)

replace github.com/project-flogo/cml/action => ../../action
