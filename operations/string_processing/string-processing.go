package string_processing

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/concat"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/contains"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/count"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/date"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/decodestring"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/encodestring"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/index"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/indexany"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/lastindex"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/levenshteinDistance"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/levenshteinSimilarity"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/matchregex"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/repeat"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/replace"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/replaceall"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/replaceregex"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/split"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/tolower"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/toupper"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/trim"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/trimleft"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/trimprefix"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/trimright"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/trimsuffix"
	"github.com/project-flogo/catalystml-flogo/operations/string_processing/uuid"
)

func init() {
	_ = operation.Register(&count.Operation{}, count.New)
	_ = operation.Register(&replace.Operation{}, replace.New)
	_ = operation.Register(&replace.Operation{}, indexany.New)
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
	_ = operation.Register(&encodestring.Operation{}, encodestring.New)
	_ = operation.Register(&decodestring.Operation{}, decodestring.New)
	_ = operation.Register(&date.Operation{}, date.New)
	_ = operation.Register(&replaceall.Operation{}, replaceall.New)
	_ = operation.Register(&replaceregex.Operation{}, replaceregex.New)
	_ = operation.Register(&trim.Operation{}, trim.New)
	_ = operation.Register(&trimleft.Operation{}, trimleft.New)
	_ = operation.Register(&trimright.Operation{}, trimright.New)
	_ = operation.Register(&trimprefix.Operation{}, trimprefix.New)
	_ = operation.Register(&trimsuffix.Operation{}, trimsuffix.New)
}
