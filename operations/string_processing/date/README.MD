# date

## Overview

### Description
(year,month,day, etc.) extract date information from string

### Implementation details
Uses the time.Parse library/function from golang.

Input
    data - string : with date time information
        Optional = False
Params
	format - string : see https://golang.org/pkg/time/
        Optional" = False

OutputType - datetime object : datatime object includes properties, month etc.

## Compliance to Spec

### Rough level of compliance  
100%