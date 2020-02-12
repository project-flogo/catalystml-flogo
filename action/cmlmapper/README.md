## CML Mapper .
CML mapper is a mapper private to CML
# Uses
It gets the required value from the underlying data
in the CML pipeline.


Eg. If  `math1` is defined as :
 [ 1  2    3
   4 "Abc" 6 ]
Using this mapper we can directly do $math1[0][2] within the CML Specification to get the desired value.

# Implementation Details

There are two main methods : NewExpression and Resolve. 

NewExpression takes in a string such as `math1[0][2]` and converts into an array of "dereferenceStruct" which helps us navigate variable "math1".

Resolve takes in the "dereferenceStruct" and "scope" (which contains the variable) and gives us the dersired value from the variable "math1"