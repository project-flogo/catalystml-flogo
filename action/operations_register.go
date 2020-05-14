package action

import (
	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/apply"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/binning"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/filter"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/ifin"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/ifnotin"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/oneHotEncoding"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/replaceValue"
	"github.com/project-flogo/catalystml-flogo/action/operations/cleaning/set"

	"github.com/project-flogo/catalystml-flogo/action/operations/image_processing/grayscale"
	"github.com/project-flogo/catalystml-flogo/action/operations/image_processing/img2tensor"
	"github.com/project-flogo/catalystml-flogo/action/operations/image_processing/resize"
	subsectiontoimage "github.com/project-flogo/catalystml-flogo/action/operations/image_processing/subSectionToImage"
	"github.com/project-flogo/catalystml-flogo/action/operations/image_processing/tensor2image"

	"github.com/project-flogo/catalystml-flogo/action/operations/math/divPairWise"
	"github.com/project-flogo/catalystml-flogo/action/operations/math/mean"
	"github.com/project-flogo/catalystml-flogo/action/operations/math/multPairWise"
	"github.com/project-flogo/catalystml-flogo/action/operations/math/norm"
	"github.com/project-flogo/catalystml-flogo/action/operations/math/normalize"
	"github.com/project-flogo/catalystml-flogo/action/operations/math/scale"

	"github.com/project-flogo/catalystml-flogo/action/operations/nlp/getstopwords"
	"github.com/project-flogo/catalystml-flogo/action/operations/nlp/postag"
	"github.com/project-flogo/catalystml-flogo/action/operations/nlp/segment"
	"github.com/project-flogo/catalystml-flogo/action/operations/nlp/stem"
	"github.com/project-flogo/catalystml-flogo/action/operations/nlp/tokenize"

	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/addCol2Table"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/cast"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/concatMap"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/dropCol"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/flatten"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/groupBy"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/join"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/map2table"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/pivot"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/reshape"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/table2map"
	"github.com/project-flogo/catalystml-flogo/action/operations/restructuring/transpose"

	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/concat"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/contains"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/count"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/date"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/decodestring"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/encodestring"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/index"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/lastindex"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/levenshteinDistance"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/levenshteinSimilarity"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/matchregex"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/repeat"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/replace"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/split"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/tolower"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/toupper"
	"github.com/project-flogo/catalystml-flogo/action/operations/string_processing/uuid"
)

func init() {
	_ = operation.Register(&apply.Operation{}, apply.New)
	_ = operation.Register(&replaceValue.Operation{}, replaceValue.New)
	_ = operation.Register(&set.Operation{}, set.New)
	_ = operation.Register(&ifnotin.Operation{}, ifnotin.New)
	_ = operation.Register(&ifin.Operation{}, ifin.New)
	_ = operation.Register(&oneHotEncoding.Operation{}, oneHotEncoding.New)
	_ = operation.Register(&set.Operation{}, binning.New)
	_ = operation.Register(&set.Operation{}, filter.New)

	_ = operation.Register(&img2tensor.Operation{}, img2tensor.New)
	_ = operation.Register(&resize.Operation{}, resize.New)
	_ = operation.Register(&grayscale.Operation{}, grayscale.New)
	_ = operation.Register(&subsectiontoimage.Operation{}, subsectiontoimage.New)
	_ = operation.Register(&tensor2image.Operation{}, tensor2image.New)

	_ = operation.Register(&mean.Operation{}, mean.New)
	_ = operation.Register(&norm.Operation{}, norm.New)
	_ = operation.Register(&normalize.Operation{}, normalize.New)
	_ = operation.Register(&scale.Operation{}, scale.New)
	_ = operation.Register(&multPairWise.Operation{}, multPairWise.New)
	_ = operation.Register(&divPairWise.Operation{}, divPairWise.New)

	_ = operation.Register(&stem.Operation{}, stem.New)
	_ = operation.Register(&tokenize.Operation{}, tokenize.New)
	_ = operation.Register(&getstopwords.Operation{}, getstopwords.New)
	_ = operation.Register(&segment.Operation{}, segment.New)
	_ = operation.Register(&postag.Operation{}, postag.New)

	_ = operation.Register(&addCol2Table.Operation{}, addCol2Table.New)
	_ = operation.Register(&cast.Operation{}, cast.New)
	_ = operation.Register(&concatMap.Operation{}, concatMap.New)
	_ = operation.Register(&dropCol.Operation{}, dropCol.New)
	_ = operation.Register(&flatten.Operation{}, flatten.New)
	_ = operation.Register(&groupBy.Operation{}, groupBy.New)
	_ = operation.Register(&join.Operation{}, join.New)
	_ = operation.Register(&map2table.Operation{}, map2table.New)
	_ = operation.Register(&pivot.Operation{}, pivot.New)
	_ = operation.Register(&reshape.Operation{}, reshape.New)
	_ = operation.Register(&table2map.Operation{}, table2map.New)
	_ = operation.Register(&transpose.Operation{}, transpose.New)

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
	_ = operation.Register(&index.Operation{}, encodestring.New)
	_ = operation.Register(&lastindex.Operation{}, decodestring.New)
	_ = operation.Register(&matchregex.Operation{}, date.New)
}
