## Goroutines and Channels

* **Goroutines** - program works on several tasks at once
  * use **go** key word next to a function to run a go routine
  * main function is started as a goroutine and is the first goroutine to end so channels are used to delay end of main call so go routine can finish
  * cannot call **go func()** if function returns a value
* **Channels** - goroutines cordinate using channels wich allow you to send data to each other and synchronize
  * use **chan** key word to declare channel
    * `var myChannel chan float64`
    * `myChannel = make(chan float64)`
    * `myChannel := make(chan float64)`
    * each type has its own channel(int, float64, struct, etc...)
    * **<-** - to send a retrieve value from the channel
      * send a value into the channel(in a function) if channel is a parameter
      * pass value from channel into another function
      * only be used to communicate between goroutines
      * **blocking** - sending go routine has value before recieving channel attempts to use it
        * send operations block the send goroutine until another goroutine executes a recieve operation on the sam channel
        * recieve operation blocks the recieving goroutine until another goroutine execultes a send operation on the same channel
        * **synchronize** - coordinate goroutine timing
      * keep loops for sending and recieving operations of a channel seperate
* **Concurrency** - allows a program to pause one task and work on other tasks
* **Parallelism** -  - running tasks simultaneously
* 