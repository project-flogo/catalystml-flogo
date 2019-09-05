package getstopwords

import (
	"fmt"
	"io/ioutil"
	"strings"

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

	fmt.Println(p)

	if p.Lang == "" {
		p.Lang = "en"
	}

	if p.Lib == "" {
		p.Lib = "nltk"
	}

	if err != nil {
		return nil, err
	}

	return &Operation{params: p, logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	fmt.Println(a.params)
	lib := a.params.Lib
	lang := a.params.Lang
	path := a.params.FileLoc
	merge := a.params.Merge

	a.logger.Info("Starting GetStopWords Operation")
	pathprint := path
	if path == "" {
		pathprint = "None"
	} else {
		pathprint = path
	}
	a.logger.Infof("Lib: %s, Lang: %s, FileLoc: %s, Merge: %t ", lib, lang, pathprint, merge)

	var stoplist []string
	liblist := []string{}
	if lib == "nltk" && lang == "en" {
		liblist = []string{"wasn", "doesn", "most", "hasn't", "isn", "having", "or", "isn't", "will", "hadn't", "more", "here", "won't", "aren't", "between", "won", "such", "shan", "you'd", "be", "yourselves", "their", "above", "by", "why", "shouldn't", "i", "these", "was", "it", "so", "but", "now", "mustn't", "mightn", "aren", "for", "haven't", "she", "m", "has", "doing", "the", "don't", "she's", "it's", "needn", "against", "not", "ours", "to", "does", "re", "our", "my", "each", "o", "under", "am", "didn", "just", "do", "of", "further", "wasn't", "weren", "hadn", "nor", "hers", "were", "being", "which", "during", "then", "myself", "until", "down", "should've", "hasn", "doesn't", "ma", "didn't", "themselves", "that", "t", "with", "shan't", "how", "have", "him", "again", "who", "at", "they", "her", "only", "a", "itself", "can", "all", "shouldn", "on", "any", "is", "too", "me", "about", "its", "been", "ll", "once", "both", "his", "from", "where", "over", "whom", "you've", "into", "same", "wouldn", "s", "yours", "did", "if", "ain", "your", "and", "than", "out", "are", "them", "an", "few", "y", "ve", "ourselves", "in", "theirs", "herself", "you'll", "what", "because", "off", "you", "should", "while", "before", "below", "haven", "as", "some", "we", "those", "own", "through", "after", "he", "you're", "when", "couldn", "couldn't", "no", "mightn't", "d", "himself", "needn't", "up", "weren't", "yourself", "wouldn't", "had", "don", "that'll", "mustn", "this", "there", "other", "very"}
	}

	listfromfile := []string{}
	if path != "" {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		listfromfile = strings.Split(string(b), "\n")

	}

	if merge {
		if path != "" {
			stoplist = append(stoplist, listfromfile...)
		}
		if lib != "none" {
			stoplist = append(stoplist, liblist...)
		}
	} else {
		if lib == "none" && path == "" {
			return nil, fmt.Errorf("merge is false but neither a lib or a file is provided ")
		} else if lib != "none" && path != "" {
			return nil, fmt.Errorf("merge is false but both a lib and a file is provided ")
		} else if lib != "none" {
			stoplist = liblist
		} else if path != "" {
			stoplist = listfromfile
		}
	}

	if len(stoplist) == 0 {
		a.logger.Info("WARNING: length of stopword list is zero")
	}

	out := stoplist
	return out, nil
}
