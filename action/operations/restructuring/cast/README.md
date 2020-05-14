# cast

## Overview

### Description
Convert the base datatype of a data structure or datatype from one base type to another (i.e. [int32,int32] to [flat64,float64]).  It is useful to note that maps of arrays are allowed and handle by the specification.   Allowed types for toType: int64,float64,string,int32,float32,boolean.

### Implementation details
The main data structure manipulation is done with relfect switch statements.  But changing the base datatype is gone with the cast library used in Hugo:  github.com/spf13/cast.

## Compliance to Spec

### Rough level of compliance  
100%

