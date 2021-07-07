* **Formatting Output**
  * Println
  * Printf - format value
  * Sprintf - returns a string instead of printing it
    |Verb  |Output  |
    |---------|---------|
    | %f     | floating point number         |
    | %d     | decimal integer         |
    | %s     | String         |
    | %t     | Boolean        |
    | %v     | any value        |
    | %#v     | any value as formatted in Go language        |
    | %T     | Type of supplied value (int,string,etc)       |
    | %%     | A literal percent sign         |
* **Function**
  * declaration begins with **func** keyword
  * if function is declared in same file it is being called, just use the function name otherwise you must use **PackageName.Function**
  * Functions that are capitalized can be exported
  * Parameter - variable that is local to a function and set when the function is called
    * passed like **variable *type***
    * Provide return value after parameters if functions returns value
        ```go
        func paintNeeded(width float64, height float64) float64 {
                area := width * height
                return area / metersPerLiter
        }
        ```
    * Go requires you the return the value specified and the return must be the last line within function scope
    * To have multiple return values pass as follows **(int,bool,string)**
* **Error package**
  * **errors.New** creates an error
  * output with **var.Method()**
  * You can also use **fmt.Errorf** or **fmt.Println(err.Error())**

* **Pointers**
  * address of a variable that stores a value
  * `&variable` returns the hexadecimal address  for the value assigned by a variable
    * equivalent to the pointer when printing and performing arithmetic
  * Types
    * Use the **reflect** package to get type of pointers
  * `*pointer` - pointer variable can hold pointers of one type of value 
    * equivalent to the value when printing and performing arithmethic
    * when parameter is a pointer, you must pass the pointer address
  * 