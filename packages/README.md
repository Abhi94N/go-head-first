* Go workspace directory
  * Workspace - directory named **go** in the current user's home directory
    * **bin** - holds compiled executable programs
    * **pkg** - holds compiled binary package files
        * [what-is-the-use-of-pkg-directory-in-go](https://stackoverflow.com/questions/47369621/what-is-the-use-of-pkg-directory-in-go)
    * **src** - holds src code
      * has sub directories for each package
        * each package has one or more source codes
        * specify package name in each file 
        * Capitalized functions in packages can be exported
      * Not required to store code for executable in a sub directory but it is best practice(where main is)
* **Packages**
  * Package names should be lower cased
  * name should be abbreviated if obvious
  * should be only one to two words
    * don't use undercore or capitalization
  * don't use name that users can use as local variable
  * You do not need a package qualifier if function is called inside the same package****
  * Packages also export **constants** - variables that never change
    * uses `const` instead of `var` keyword
    * like Packages, package qualifier is needed for const to work outside of package
* **Nested Packages**
  * only require the name of the package but must be brought in with directory format
    * `package/subpackage`
* **Important Commands**
  * `go build file` - converts go file to exc
  * `go install {directory with main}` - converts into executable
    * NOTE: must be called main.go for file to be a executable
  * `go get github.com/Abhi94N/greetings {Some optional function or const}` - installs package from git 
  * `go doc {package}` - shows details of package\
  * `go doc -http=:6060` - shows documentation locally with third party packages you crated
* Update **GOPATH** to change directories if you have multiple workspaces
  * use `export` for unix and `set` for windows
  * Go looks for the **go** directory in your root and sets that as the workspace directory by default but you can change it with **GOPATH**
  * **github.com** exists under source to as a package qualifier to identify packages of different origin
* **Publish a package**
    * Create a gitrepo with name of package and pass src files and subdirectories
    * Each github repo can be its own published package
    * use go get to download and install the package
