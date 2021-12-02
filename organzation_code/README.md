* organized in `package` - collection of source files in the same directory compiled together
  * functions, types, variables, and constants are diefined in one source file are visible to all other source files in the same package
* Repository contains one or more `modules` - collection of related Go packages that are released together
  * `go.mod` declares `module path` - the import path prefix for all packages within the module
    * module contains the packages in directory containing its go.mod file as well as subdirectories of that directory, up to the next directory consisting of go mod