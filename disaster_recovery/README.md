* **Deferring function calls** - use `defer` argument to defer a function until the end of function call such as a return function
* **Recursion** calls its own function within method until a conditional is met
  * anything after the function call inside the function definition will be called after the recursion condition has been met
* **Panic** - empty interface that accepts a single argument to statisfy its empty interface
  * allows you to greatly simplify error handling by not having to return the error and then log it but show the call stack instead
  * With Panic you no longer have to return the error function
  * Only used in little programs to just terminate the program once an error appears
  * Use for inaccessible files, netowrk failures, and bad user input
  * Panic should be used when the error is irreversible or something that should never happen
* **Call stack** - list of the function calls that are active at any given point
* **Stack trace** - when a program panics, a stack trace listst the call stack
* **Recover** - return an error instead of a panic which leads to a confusing stack trace
    * must defer recover function to run if used with panic as panic exits the function and the recover is never called
      * also must wrap recover in a function and defer that function
      * when deferring, the recover function will return params that are passed to panic function
    * recover will intercept any panic