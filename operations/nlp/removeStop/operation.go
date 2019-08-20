package removeStop

import (
	"fmt"

	"github.com/bbalet/stopwords"
	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)
	if p.Lang == "" {
		p.Lang = "en"
	}

	if err != nil {
		return nil, err
	}

	blah := []string{"wasn", "doesn", "most", "hasn't", "isn", "having", "or", "isn't", "will", "hadn't", "more", "here", "won't", "aren't", "between", "won", "such", "shan", "you'd", "be", "yourselves", "their", "above", "by", "why", "shouldn't", "i", "these", "was", "it", "so", "but", "now", "mustn't", "mightn", "aren", "for", "haven't", "she", "m", "has", "doing", "the", "don't", "she's", "it's", "needn", "against", "not", "ours", "to", "does", "re", "our", "my", "each", "o", "under", "am", "didn", "just", "do", "of", "further", "wasn't", "weren", "hadn", "nor", "hers", "were", "being", "which", "during", "then", "myself", "until", "down", "should've", "hasn", "doesn't", "ma", "didn't", "themselves", "that", "t", "with", "shan't", "how", "have", "him", "again", "who", "at", "they", "her", "only", "a", "itself", "can", "all", "shouldn", "on", "any", "is", "too", "me", "about", "its", "been", "ll", "once", "both", "his", "from", "where", "over", "whom", "you've", "into", "same", "wouldn", "s", "yours", "did", "if", "ain", "your", "and", "than", "out", "are", "them", "an", "few", "y", "ve", "ourselves", "in", "theirs", "herself", "you'll", "what", "because", "off", "you", "should", "while", "before", "below", "haven", "as", "some", "we", "those", "own", "through", "after", "he", "you're", "when", "couldn", "couldn't", "no", "mightn't", "d", "himself", "needn't", "up", "weren't", "yourself", "wouldn't", "had", "don", "that'll", "mustn", "this", "there", "other", "very"}
	fmt.Println(blah)

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	// I am changing the word parser to keep punctuation
	// 		I just added \pP to the default regex
	stopwords.OverwriteWordSegmenter(`[\pL\p{Mc}\p{Mn}\pP-_']+`)

	a.logger.Debug("Executing operation...", input.Str)
	a.logger.Info("Removing stopwords...", input.Str)
	out := stopwords.CleanString(input.Str, a.params.Lang, false)

	a.logger.Info("String without stopwords: ", out)

	return out, nil
}
