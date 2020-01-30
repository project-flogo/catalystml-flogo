## API's

These provide API's to directly run CML pipeline directly. 

First we set up config for the CML pipeline; Currently you can set the location of cml file Eg.

`fileCfg := api.SetURISettings("file://cml.json")`

Then we initilaize the CML pipeline using the above configuartion. Eg.

`act, err := api.NewAction(fileCfg)` 

You could also add different config within this method Eg . `NewAction(cfg1, cfg2)`

Once initilaized you can run the pipleine using `Run()` Eg.

`out, err := api.Run(act, inputs)`
