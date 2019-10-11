package main

import (
	"context"
	"fmt"

	_ "github.com/project-flogo/catalystml-flogo/action"
	_ "github.com/project-flogo/catalystml-flogo/operations/cleaning"
	_ "github.com/project-flogo/catalystml-flogo/operations/nlp"
	_ "github.com/project-flogo/catalystml-flogo/operations/string_processing"
	"github.com/project-flogo/core/action"
)

func main() {

	factory := action.GetFactory("github.com/project-flogo/catalystml-flogo/action")

	var act action.Action

	act, _ = factory.New(&action.Config{Settings: map[string]interface{}{"catalystMlURI": "file://cml.json"}})

	input := "Natural Language Processing (NLP) is all about leveraging tools, techniques and algorithms to process and understand natural language-based data, which is usually unstructured like text, speech and so on. In this series of articles, we will be looking at tried and tested strategies, techniques and workflows which can be leveraged by practitioners and data scientists to extract useful insights from text data. We will also cover some useful and interesting use-cases for NLP. This article will be all about processing and understanding text data with tutorials and hands-on examples."

	out, _ := act.(action.SyncAction).Run(context.Background(), map[string]interface{}{"paragraph": input})

	fmt.Println("Output...", out)
}
