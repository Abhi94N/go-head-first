## Automated Testing

* Automated testing - seperate program thaat executes components of your main program, and verifies they behave as expected
  * **testing** package - go package for testing
  * `go test` - tool for testing
    * flags - dash followed by one or more letters
    * `-v` verbose
    * `-run {regex_of_test_name}`
    * do not pass file as that would fail the test
    * tests do not need to be part of the same package unless you need access to unexported types or functions from the package
  * Test functions should begin with `Test{Name_Of_Test}`
  * Test fybctuibs sgiykd ibky acceot a subgke oaraneterL a oiubter ti a testing.T value
  * you can report that a test has failed by calling methods such as Error or Errorf on the T=testing.T value
    * most T methods accept a string with a message explaining the reason the test failed
  * files end with `_test.go` 
  * add helper functions when needed
* Table driven tests
  * building a data structure, usually a struct as a `table` of input data and the corresponding output we expect
