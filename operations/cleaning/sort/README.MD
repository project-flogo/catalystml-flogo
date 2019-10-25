# sort

## Overview

### Description
sort a matrix/map based on given columns 

### Implementation details

Input
    Data - map[string]interface{} or [][]interface{}
        Optional=False
Params
    Ascending - [boolean]
    Axis - int : 0=vertical/column, 1=horizontal/row
        Optional=True
        Default=0
    By - []interface{} : elements 1. conlumn label (string) or 2. column number (int from 0)
	NaPosition - Puts NaNs at the beginning if first; last puts NaNs at the end.
		Allows {"first", "last"}, Default "last"

OutputType - map[string]interface{}


## Compliance to Spec

### Rough level of compliance  
100%