package ifin

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["arr0"] = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2}
	inputs["arr1"] = []interface{}{2, 4, 6, 8, 10}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	fmt.Println(output)
	assert.NotNil(t, output)
	// assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}

func TestStr(t *testing.T) {

	inputs := make(map[string]interface{})
	inputs["arr0"] = []interface{}{"a", "the", "was", "help"}
	inputs["arr1"] = []interface{}{"wasn", "doesn", "most", "hasn't", "isn", "having", "or", "isn't", "will", "hadn't", "more", "here", "won't", "aren't", "between", "won", "such", "shan", "you'd", "be", "yourselves", "their", "above", "by", "why", "shouldn't", "i", "these", "was", "it", "so", "but", "now", "mustn't", "mightn", "aren", "for", "haven't", "she", "m", "has", "doing", "the", "don't", "she's", "it's", "needn", "against", "not", "ours", "to", "does", "re", "our", "my", "each", "o", "under", "am", "didn", "just", "do", "of", "further", "wasn't", "weren", "hadn", "nor", "hers", "were", "being", "which", "during", "then", "myself", "until", "down", "should've", "hasn", "doesn't", "ma", "didn't", "themselves", "that", "t", "with", "shan't", "how", "have", "him", "again", "who", "at", "they", "her", "only", "a", "itself", "can", "all", "shouldn", "on", "any", "is", "too", "me", "about", "its", "been", "ll", "once", "both", "his", "from", "where", "over", "whom", "you've", "into", "same", "wouldn", "s", "yours", "did", "if", "ain", "your", "and", "than", "out", "are", "them", "an", "few", "y", "ve", "ourselves", "in", "theirs", "herself", "you'll", "what", "because", "off", "you", "should", "while", "before", "below", "haven", "as", "some", "we", "those", "own", "through", "after", "he", "you're", "when", "couldn", "couldn't", "no", "mightn't", "d", "himself", "needn't", "up", "weren't", "yourself", "wouldn't", "had", "don", "that'll", "mustn", "this", "there", "other", "very"}

	opt := &Operation{logger: log.RootLogger()}

	output, err := opt.Eval(inputs)
	fmt.Println(output)
	assert.NotNil(t, output)
	// assert.Equal(t, output, " hotspot ", "two should be the same")
	assert.Nil(t, err)

}
