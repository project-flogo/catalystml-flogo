# filter

## Overview

### Description
keep/remove rows with certain values

### Implementation details

Input
    Data - map[string]interface{} or [][]interface{}
        Optional = False
	Value - [int, float, string, NaN]
		Optional=True
	    Default=Nil
	filterType - [string]
	    Optional=True
	    Default=Remove
	    Acceptable values = “Remove”,”Keep”
Params
    Axis - int : 0 = vertical/column, 1 = horizontal/row
		Allowed = [0, 1]
        Default = 0
        Optional" = True
	Col - [string]
        Optional=False

OutputType - map[string]interface{} or [][]interface{}

## Compliance to Spec

### Rough level of compliance  
100%