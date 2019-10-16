# pivot

## Overview
reshape a matrix or map by re-indexing like the pivot function in pandas

### Description
Reshape data (produce a “pivot” table) based on column values. Uses unique values from specified index / columns to form axes of the resulting DataFrame. It currently supports data aggregations such as Sum, Count, Mean, Min and Max.

### Implementation details

Input
    Data - [array or map] Column(s) to use for populating new frame’s values.

Params
    Index - []string : Name of columns which value to be used to make new frame’s index.
    Columns - []string : Name of columns which value to be used to make new frame’s columns.
    Aggregate - map[string]string : Map key is groupKey, map value is aggregate function. Currently support Sum, Count, Mean, Min, Max 

OutputType - [DataFrame]

## Compliance to Spec

### Rough level of compliance  
100%