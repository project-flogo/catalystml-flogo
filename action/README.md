# Introduction

This is a Flogo Action implementation of the CatalystMl Specification.  It can also act as a standalone Go implementation of the CatalystMl. For more information about Flogo Action please visit `` . 

# Implementation Details :

This Flogo Action creates a pipeline of the operations specified in JSON spec and executes it sequentially. One of the challenging aspect of the pipeline is resolving the data in the spec. 
Mappers, Resolvers and Scope Interfaces provided by Project-Flogo core reposisotiry is used to simplify the resolution of data. For more information please visit [Project-flogo/core](https://github.com/project-flogo/core). 

## Detailed Implementation:

The implementation of CML specification in Flogo can be divided into three steps:

 * Configuration.
 * Initialization.
 * Execution. 

### Configuration.
    
   The CML JSON spec is unmarshalled into a [DefinitionConfig](pipeline/definition.go) struct . We use this struct to set up [Instance](pipeline/instance.go) of the 
pipleine. 

### Initialization.
During the initialization of pipeline instance we set up [Mappers](https://github.com/project-flogo/core/blob/master/data/mapper/mapper.go) for Input and Output of the CML.  We also
initialize the [Operations](operation/operation.go) defined in the CML spec. We do this by getting the registered operations and initializing those using
[factories](operation/registry.go). Here we also set up the input and output mappers for each operation.

### Execution.
After the initialization, when the action is called. We iterate over an array of initialized operations. We [resolve](https://github.com/project-flogo/core/blob/master/data/resolve/resolve.go) the input 
and output mappers for each operation and add the output in the pipeline [scope](https://github.com/project-flogo/core/blob/master/data/resolve/scope.go). We resolve the inputs for the operation using the pipeline
scope. The pipeline scope is nothing but the collection of all the variables, which are the output of each operations and input of CML, and its value
. After execution of each operations we resolve the output of CML and return that value.
