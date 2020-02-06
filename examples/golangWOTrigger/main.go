package main

import (
	"fmt"

	_ "github.com/project-flogo/catalystml-flogo/action"
	"github.com/project-flogo/catalystml-flogo/action/api"
	_ "github.com/project-flogo/catalystml-flogo/operations/cleaning"
	_ "github.com/project-flogo/catalystml-flogo/operations/image_processing"
	_ "github.com/project-flogo/catalystml-flogo/operations/nlp"
	_ "github.com/project-flogo/catalystml-flogo/operations/string_processing"
)

func main() {

	cfg := api.SetURISettings("file://cml.json")

	act, err := api.NewAction(cfg)

	if err != nil {
		return
	}

	input := "Natural Language Processing (NLP) is all about leveraging tools, techniques and algorithms to process and understand natural language-based data, which is usually unstructured like text, speech and so on. In this series of articles, we will be looking at tried and tested strategies, techniques and workflows which can be leveraged by practitioners and data scientists to extract useful insights from text data. We will also cover some useful and interesting use-cases for NLP. This article will be all about processing and understanding text data with tutorials and hands-on examples."

	inputs := make(map[string]interface{})
	inputs["paragraph"] = input

	out, err := api.Run(act, inputs)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Output...", out)
}
