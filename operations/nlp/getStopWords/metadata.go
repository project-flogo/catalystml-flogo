package getstopwords

type Params struct {
	Lang    string `md:"lang",required=false,allowed=["en"]`
	Lib     string `md:"lib",required=false,allowed=["nltk","none"]`
	FileLoc string `md:"fileLoc",required=false`
	Merge   bool   `md:"merge",required=false`
}

