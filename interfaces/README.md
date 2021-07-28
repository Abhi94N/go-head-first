**Interfaces** - defined as a set of methods that certain values that are expected to have
  * any type that has all methods listed in an interface definition is said to have satisifed the interface
    * can be used any where 
    * type can have methods in addition to interface
    * method names, parameter types and return value types must match as defined in interface
    * a type can satisfy multiple interfaces and an interface can have multiple types that it satisfied
        ```go
        type myInterface interface {
            methodWithParameters()
            methodWithParameter(float64)
            methodWithREturnValue() string
          }
        ```
    * **Concrete types** - specifies what its values can do and what tyhe are
      * holds type and value's data
    * **Interface types** - don't specify value but states underlying type
      * only describe what value can do
      * what methods it has
  * Its ok to assign a type that has other methods toa  avirable with an interface type as long as you don't call those other methods
  * When assigning type to an interface with a pointer reciever, pass in the address instead otherwise interface method leads to implementation failure
    * `var t Toogleable = &s`
  * **Type Assertion** - when you have a value of a concrete type to a variable with an interface type, type assertion returns the concrety type back
    * have two return values one is the optional bool value that checks if there er errors
  
  ```go
    var noiseMaker NoiseMaker = Robot("Optimus Prime")
    //() returns back the original concret type
    var robot Robot = noiseMaker.(Robot)
    //Note when in diffrent package, mustl call .(package.Type)
  ```
* **error** is an interface but does not need to be imported because like `int` and `string`, `error` type is a **predeclared identifier** and is part of the universal block
    ```go
    type error interface {
        Error() string
      }
    ```
* **fmt.Stringer** is another predeclared interface for string operations

* **empty interface** or `type interface{}`
  * doesn't have any methods that satisfy it and so every type satisfies it
  * **Println** function has an empty interface so it can take any value and print it out


NOTE: `go doc {pkg} {function}` - for docs 
