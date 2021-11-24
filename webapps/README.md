 ## Web Apps

* `net/http` package
* `Browser`
  * send request to web page which goes to server
  * Server gets he appropriate page and sends it back to the browser in a response
* `http.ResponseWriter` - object to write responses when request is called and send it to the browser
* `*http.Request` - represents the request from the browser
* **resource Path** - tells tge server which of its many resources you want to act on
* `http.HanndleFunc(resource_path, handlerFunc)` - looks for resource path and call the hanlder function
* **First-Class functions** - functions can be assigned to variables and then called from those variables as long as they are of type `func()`
  * when declaring func variables, it must provide parameters and return values if it has any check out [hello.go](./src/hello.go)
* 