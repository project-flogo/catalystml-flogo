# interpolateMissing
Fill in missing values

## Overview

### Description
Select a numerical column list and convert it into bins.

### Implementation details
    Input
        data - [array, map] data to be filled in
        col - [int, string, array] the column(s) to be filled in
    Params
        how - [string] mean: substitutes mean of included values for Missing, linear: interpolates linearly from adjacent values
            Optional - True
            Default- "mean"
            allowed: ["mean","linear"]
        edges - [string] for linear how to handle edges
            Optional - True
            allowed: []
    OutputType - [array, map]

OutputType - map[string]interface{}
## Compliance to Spec

### Rough level of compliance  

100%