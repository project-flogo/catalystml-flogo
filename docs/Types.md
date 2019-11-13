# Types

The CatalystML specification defines certain supported data types and structures that can be passed between operations as well as as input/output.  Here we defined how these types and structures as they are handled within this specification.  All the handled types used are defined using golang (slice, map, byte, etc.)

* data structures:
    * "list" - defined as a slice (mutable array) interface ([]interface{})
    * "map" - defined as a map with string keys and interface values (map[string]interface{}
    * "image" - is handled as a slice of bytes ([]byte)
    * "datetime" - 
    * "any" - is treated as an interface{} that is then tested against each of the other types 
* data types:
    * "string" - defined as a string
    * "int32" - defined as int32
    * "int64" - defined as int64
    * "float32" - defined as float32
    * "float64" - defined as float64
    * "boolean" - defined as a boolean