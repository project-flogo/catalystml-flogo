# binning

## Overview

### Description
Select a numerical column list and convert it into bins.

### Implementation details
Input
    Data - map[string]interface{} or [][]interface{}
        Optional = False
Params
	Quantile - int : number of bin by even quantiles
        Default = 0
        Optional = True
	Bins - []float64 : array of bin boundaries by data value
        Optional = True
	Labels - []string : used as labels for the resulting bins
        Optional = True
	Column - string : target column for binning
        Optional = True
	Retbins - bool : Whether to return the (bins, labels) or not
        Default = False
        Optional = True
	Precision - int : precision of bin boundaries
        Optional = True
	Duplicates - string : action on duplicate data raise error (still output data) or drop data
		Allowed = ["raise", "drop"]
        Default = "raise"
        Optional = True

OutputType - map[string]interface{}
## Compliance to Spec

### Rough level of compliance  

100%