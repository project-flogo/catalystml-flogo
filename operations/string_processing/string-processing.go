package string_processing

import (
	// "github.com/project-flogo/cml/operations/string_processing/phonenumber"
	// "github.com/project-flogo/cml/operations/string_processing/geoencoding"
	"github.com/project-flogo/cml/operations/string_processing/count"
	"github.com/project-flogo/cml/operations/string_processing/replace"
	"github.com/project-flogo/cml/operations/string_processing/levenshteinDistance"
	"github.com/project-flogo/cml/operations/string_processing/levenshteinSimilarity"
	"github.com/project-flogo/cml/operations/string_processing/tolower"
	"github.com/project-flogo/cml/operations/string_processing/toupper"
	"github.com/project-flogo/cml/action/operation"
)
	
func init() {
	// _ = operation.Register(&phonenumber.Operation{})
	// _ = operation.Register(&geoencoding.Operation{}, geoencoding.New)
	_ = operation.Register(&count.Operation{}, count.New)
	_ = operation.Register(&replace.Operation{}, replace.New)
	_ = operation.Register(&levenshteinDistance.Operation{}, levenshteinDistance.New)
	_ = operation.Register(&levenshteinSimilarity.Operation{}, levenshteinSimilarity.New)
	_ = operation.Register(&tolower.Operation{}, tolower.New)
	_ = operation.Register(&toupper.Operation{}, toupper.New)
}
