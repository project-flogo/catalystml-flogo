# Introduction

This is a Flogo Action implementation of the CatalystMl Specification.  It can also act as a standalone Go implementation of the CatalystMl. For more information about [Flogo Action](https://github.com/project-flogo/core/tree/master/action) . 
# Implementation Details :
This Flogo Action creates a pipeline of the operations specified in JSON spec and executes it sequentially. One of the challenging aspect of the pipeline is resolving the data in the spec. 
Mappers, Resolvers and Scope Interfaces provided by Project-Flogo core reposisotiry is used to simplify the resolution of data. For more information please visit [Project-flogo/core](https://github.com/project-flogo/core). 
## Detailed Implementation:
The implementation of CML specification in Flogo can be divided into three steps:
 * Configuration.
 * Initialization.
 * Execution. 
### Configuration.
    
   The CML JSON spec is unmarshalled into a [DefinitionConfig](pipeline/definition.go) struct .This struct is used to set up [Instance](pipeline/instance.go) of the pipleine. 
### Initialization.
During the initialization of pipeline instance  [Mappers](https://github.com/project-flogo/core/blob/master/data/mapper/mapper.go) are set up for Input and Output of the CML.  
[Operations](operation/operation.go) (defined in the CML spec) are also initialized. The Registered operations are fetched  and initialized  using
[factories](operation/registry.go). The mappers for input and output of each operation are also initialized.
### Execution.
After the initialization, when the action is called, the program iterates over each operation executing it. The input mappers of each operations is resolved before executing it. Only the inputs defined by the operation are sent over for the execution. There can be other variables in the pipeline scope that are not passed in execution of operation.
After the execution of operation, the output mappers are resolved and are added in the pipeline scope; even if not needed by further operations. For more information on how mappers and resolvers work Please visit [Mappers](https://github.com/project-flogo/core/blob/master/data/mapper/mapper.go)
, [Resolvers](https://github.com/project-flogo/core/blob/master/data/resolve/resolve.go), [Scope](https://github.com/project-flogo/core/blob/master/data/resolve/scope.go). The resolution of input and output is done using pipeline scope . The pipeline scope is nothing but the collection of all the variables, which are the output of each operations and input of CML, and its value
. After execution of all the operations the output of the CML is resolved and returned