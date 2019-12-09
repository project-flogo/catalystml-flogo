# Engine versus Operation  

During the creation of this implementation it was decided to break the implementation into two parts.  The first  is the individual operations that perform the data transforms. The operations are handles as specific functions that are contributed individually through a given issue request to match the specification.  Each Operation is constructed of at least the following:

    * operation.go and operation_test.go files that define the function the operation performs and its tests.
    * metadata.go which defines the inputs and parameters of the operation
    * a README.md that describes how close the operation implementation is the the specification
    * a descriptor.json that defines additional data about the operation such as the version of specification the operation complies with.

The second portion of this implementaiton is the Engine that reads the structure JSON, handles the data in memory, gets inputs, writes outputs, and defines a operation object that is then extended for the individual operations.  The engine is implemented as a [flogo action](https://github.com/project-flogo/core/blob/master/action/action.go) that handles a given event a data to be manipulated making it a strong tool for this implementation.