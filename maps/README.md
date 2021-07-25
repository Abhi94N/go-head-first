* **map** - a collection where each value is accessed via key
  * declaration: `var myMap map[string]float64`
  * `myMap = make(map[string]float64)` - actually makes the map
* **map literals** - create key values in advance
  * `myMap := map[string]float64{"a": 1.2, "b": 5.6}`
* **Zero Values** - the value is either empty string or 0 if not set
  * zero value for the map itself is nil
    * maps also return a boolean value for zero values
      * `value , ok := map[key]`
* **delete** function
  * `delete (map,key)
* **for range** - allows you to loop through key and value
  * `key,value := range map`
  * handles amps in random order so use the sort package to sort out keys such as `sort.Strings()`
* 