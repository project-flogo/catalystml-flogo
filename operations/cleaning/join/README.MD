# join

## Overview

### Description
join two data objects like the Left Join command in SQL

### Implementation details

Input
    left - map[string][]interface{} : map key is column name, map value are column values. Optional=False
    right - map[string][]interface{} : map key is column name, map value are column values. Optional=False
    leftindex - []string : index column names.
    rightindex - []string : index column names.
Params
    on - []string - either index or the column name.
        Optional=True
        Default= array of “index”
    how - [string]
        Options: left, right, inner, outer
        Optional=False

OutputType - [matrix or map]

## Compliance to Spec

### Rough level of compliance  
100%