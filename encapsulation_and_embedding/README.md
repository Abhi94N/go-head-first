# go-head-first

* **Setter methods** - methods for structs to set properties of struct
  * usually Named `SetYear`
  * Setter methods need pointer recievers to update instance
  * use `error` package to deal with error handling in setter function and return the error if it exists
* **Getter methods** - methods for structs to get properties of a struct
  * getters allow for returning of unexported for a year
    * by convention getters are set to the name of the value it returns: `Year()`
* **Encapsulation** - hiding data in part of a program from code in another part
  * allows for data validation
* **Embedding** - moving structs or types as children properties of other type structs
  * promotes any exported properties to parent type struct with embedded type struct
* **Runes** - Characters
  * in package `unicode/utf8` consists of functions for runes(chars) like `RuneCountInString`