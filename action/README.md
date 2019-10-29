# Introduction

This is a Flogo Action implementation of the CatalystMl Specification.  It can also act as a standalone Go implementation of the CatalystMl. For more information about Flogo Action please visit `` . 


# Usage 
 
## Standalone Flogo App with Catalyst Ml.


## Flogo App with Flogo Flow and Catalyst Ml.

## Go App as Flogo App and Catalyst Ml.

## Standalone Go App with Catalyst Ml.


# Implementation Details :

This Flogo Action creates a pipeline of the operations specified in JSON spec and executes it sequentially. One of the challenging aspect of the pipeline is resolving the data in the spec. 
Mappers, Resolvers and Scope Interfaces provided by Project-Flogo core reposisotiry is used to simplify the resolution of data. For more information please visit. 

## Detailed Implementation:

The CML JSON spec is unmarshalled into a `DefinitionConfig` struct . We use this struct to set up `Instance` of the 
pipleine. During the initialization of pipeline instance we set up mappers for Input and Output of the CML. Then we proceed to
initialize the Operations defined in the CML spec. We do this by getting the registered operations and initializing those using
factories. Here we also set up the input and output mappers for each operation.

After the initialization, when the action is called. We iterate over an array of initialized operations. We resolve the input 
and output mappers for each operation and add the output in the pipeline scope. We resolve the inputs for the operation using the pipeline
scope. The pipeline scope is nothing but the collection of all the variables, which are the output of each operations and input of CML, and its value
. After execution of each operations we then proceed to resolve the output of CML and return that value.
