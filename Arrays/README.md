* **Array** - collection values that all share the same type
* **Elements** - values an array holds
* **index** - an elements position
* Format: `var anArray [*]type`
* Arrays also have **Zero** values when array size is determined
  * Zero values make it safe to manipulate array elements can use `array[0]++` to increment the value for that array index
* **Array literals** - declare arrays with values
  * `var array [3]int = [3]int{9,18,27}`
* **Panic** - an error that occurs when a programm is running. For arrays a runtime error fo index out of range will occur when accessing index outside of an array
* can use standard for loop or use for range
  * `for i = 0; i < len(array); i++`
  * `for..range` - keyword iterates over array
    * must specify index and array as go requires you to use every variable you declare
    * Use a blank identifier when you are not using a variable
    ```go
    for index,note := range notes {
          fmt.Println(index,note)
        }
    ```

