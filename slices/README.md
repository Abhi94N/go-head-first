* **Slices** - view of arrays that can add in length and in element
  * Declaration: `var array []type`
  * Creation: `make([]type, 7)`
  * Can also create slice literals which is better to use than arrays
    * `:= []type{value1, value2}`
  * Slice is a view of an array so if array changes, so will the slice
  * zero values will be `nil` value
* **Underlying array** - that actually holds the slice's data
  * `make` - creates an underlying array for you by slice
  * Cannot increase the size of the underlying array but can use the append function to create a new slice which copies elements to a new underlying array
* **Slice Operator** - slices an underlying array with a starting and ending index 
  * `underlyingArray[0:2]`
  * `underlyingArray[:]` - beginning index to last index
  * `underlyingArray[i:]` - beginning index at a specific value to last index
  *  `underlyingArray[:j]` - beginning index to a specific index

* **functions**
  * `append(slice, value1,value2...,valuex)`
* **variadic function** - can call varying number of arguments. this requires using an elipses, `var ...type` passed as a parameter
  * any number of variadic arguments or nonvariadic arguments can be called but only variadic args can be omitted
    * to pass a slice into a **variadic parameter** - pass like this: `var...`
* Other info
  * Use `os.Args` to get command line arguments
  