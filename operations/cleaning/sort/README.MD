# sort

## Overview

### Description
sort a matrix/map based on axis and given columns 

### Implementation details

Input
    Data - map[string]interface{} or [][]interface{}
        Optional = False
Params
    Ascending - boolean
        Default" = False
        Optional" = True
    Axis - int : 0 = vertical/column, 1 = horizontal/row
		Allowed = [0, 1]
        Default = 0
        Optional" = True
	NilPosition - string : Puts Nils at the beginning if first; last puts Nils at the end.
		Allowed = ["first", "last"]
		Default = "last"
        Optional" = True
    By - []interface{} : sort bu elements 1. conlumn label (string) or 2. column number (int from 0)
        Default" = False
        Optional" = True

OutputType - map[string]interface{}


## Compliance to Spec

### Rough level of compliance  
100%