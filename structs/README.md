* **Struct** - value constructed out of other values of different type
  * each struct has one or more fields
    ```go
    struct {
            field1 string
            field2 int
          }
    ```
  * declaration
    ```go
    var myStruct struct {
          number float64
          word string
          toggle bool
        }
    ```
  * access fields with dot operator
    * `myStruct.number`
* **Type Definitions** - allow you to create your own types
  * create a new **defined type** based on **underlying type**
  * use type keyword
    *     type myType struct {
          number float64
          word string
          toggle bool
        }
    ```
  * `types` are created outside of main at the package level scope
    * replaces type where car is a type
      * `var porsche car`
* set address of struct to pointer
  * `var pointer *myStruct = &value`
  * `(*pointer).myField` or `pointer.myField`
  * `pointer.myField = 9`
* NOTE: Whein exporting structs both the struct and the field names have to be capitalized
* Struct Literals but instead of an array it is a struct which is like a map but with fields of different value
  * `c := car{name: name, topSpeed: 225}`
* Annonymous fields when you have a type but not a Name

      ```go
      type Employee struct {
        Name   string
        Salary float64
        Address
      }
      ```
* **embedded struct** - inner struct that is stored within an outer struct using an anonymous field
  * Fields for embedded struct are **promoted** to outer struct
  * THis is allowed: `employee.City` - you do not have to inlude the name of the embedded struct

* can create types with underlying types can also be compared with its underlying literal types but not with other defined types even if they have the same underlying literal type
* NOTE: Go does not support overloading which is having multiple functions with the same name but with different parameters but this is not supported by Go
    ```go
      type Liters float64
      type Gallons float64

      carFuel := Liters(240.0)
      busFuel := Gallons(10.0)

    //to convert one type to another
      carFuel_2 := Gallons(carFuel * 0.264)
      busFuel_2 := Liters(busFuel * 3.785)
    
    //can do operation with its literal type
    fmt.Println(Liters(1.2) + 3.4)
    ```
* **Defining Methods** - different from a function as methods are called on recievers
  * Similar to class methods vs function methods 
  * **reciever parameter** - parameter before the function name and function parameters
    * `func (m MyType) sayHi() `
      * This is the type the method is associated wtih
      * basically replacement for `self` or `this`
      * can only define methods for types in same package as the type definition
      * you can use structs and embeds another package type as an anonymous fields
  * **reciever** - the value you're calling the method on
    * `value.sayHi()`
      * where value is of type `MyType` and is also the `reciever`
  * `pointer reciever parameters` - to update current instance value
    * pass param as pointers `func (n *Number) Double() `
    * `ten.Double()` Go converts nonpointer reference method calls to pointers and then runs to method
  * Cannot create multiple functions with same naeme in same package but can define multiple methods with the same name as long as each type is different
  