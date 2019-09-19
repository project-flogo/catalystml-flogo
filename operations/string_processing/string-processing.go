package string_processing

import (
	// "github.com/project-flogo/catalystml-flogo/operations/string_processing/phonenumber"
	// "github.com/project-flogo/catalystml-flogo/operations/string_processing/geoencoding"
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/count"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/levenshteinDistance"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/levenshteinSimilarity"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/repeat"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/replace"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/split"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/tolower"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/toupper"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/uuid"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/concat"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/contains"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/index"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/lastindex"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/matchregex"
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
	_ = operation.Register(&repeat.Operation{}, repeat.New)
	_ = operation.Register(&split.Operation{}, split.New)
	_ = operation.Register(&uuid.Operation{}, uuid.New)
	_ = operation.Register(&concat.Operation{}, concat.New)
	_ = operation.Register(&contains.Operation{}, contains.New)
	_ = operation.Register(&index.Operation{}, index.New)
	_ = operation.Register(&lastindex.Operation{}, lastindex.New)
	_ = operation.Register(&matchregex.Operation{}, matchregex.New)
}
