* **Package** - collection of code that does similar things
  * Main package is required to run directly from terminal
  * Package clause - gives the name of the package that this file's code is a part of
  * Common Packages
    * fmt - format print output
    * reflect - to provide type for value or variable
    * math
    * strings
* **Import** - import packages needed before writing code
  * Packages can be imported indvidually with import clause or in a list inside of ()
* **Function** - group of one or more lines thant you can call (run) from other places in your program
  * * How to call a function
  * `package.function_name()`
    * pass in arguments in function requires
  * Main function runs first
* Typical go layout
  * Package clause
  * Import statements
  * Actual code
* Common Variable Types
  * **Strings** with Strings library
    |Escape Sequence  |Value  |
    |---------|---------|
    |`\n`    |     new line    |
    |`\t`    |      new tab   |
    |`\"`     |   add dobule quote      |
    |`\\`     |      single backslash   |
    * String literals are written in double quotes
  * **Runes** - single character
    * Rune literals are written in single quotes
    * Are written in ascii so returns numeric code
  * **Booleans** - true or false
  * **Numbers** - integers or booleans
  * **Math operations and comparisons**
    * +,-,*,/
    * <,>,>=,<=,==,!=
  * variable - a piece of storage containing a value
    * short variable decleration - implicit decleration
      * `count := 4`
    * explciit decleration
      * `var {variable_name} {type}`
      * variable types
        * int,float64,string,bool
      * any variable without assigned value is the variable types zero value
        * 0,false,""
* Naming Rules
  * must begin with letter
  * **exported** - any name of function, variable or type that begins with a capital letter is exported so it can be accessed from packages outside of the current one
  * **unexported** - lower case variable/function/type name can only be accessed from within the current package
  * use camel case or upper case for variable names
    * underscores are legal but bad practice
  * obvious variables should be abbreviated like index = i
**Conversions**
  * convert types from one type to another
    * float64(), int(), [strtocov library](https://golang.org/pkg/strconv/#FormatFloat)
* Go basic commands
  * `go build` - converts source code to binary files
    * can run with optionl args
      * `./hello.go --flag value`
  * `go run` - Compiles and runs a program, without saving an executable file.
  * `go fmt` - Formates source files using Go standard formatting
  * `go version` - current go version

  