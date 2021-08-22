# goroutine

```go
func main() {

   for i := 0; i < 100; i++ {
      go func(i int) {
         fmt.Println(i+1)
      }(i)
   }

   fmt.Println("main")
   time.Sleep(time.Second)

}
```

### when the Goroutine is finished ?

the goroutine  function is finished , then the gorountine is finished.

the`main` functinon is finished to run, because the `gorouitne ` inside the `main` function is already finished.



### Goroutine Scheduling :

`GMP`  it is the implementation of the runtime level of the Go language, and is a set of scheduling system implemented by the Go language itself. Different from the operating system scheduling OS threads.



### math/rand

```go
func f()  {
   // that can help to print different random every time we run the code
   rand.Seed(time.Now().UnixNano()) 
   for i := 0; i < 5; i++ {
      r1 := rand.Int()
      r2 := rand.Intn(10)
      fmt.Println(r1, r2)
   }
}
```

`goroutine `

### Sync.WaitGroup

in Go Concurrencey, it's easy to make by using `sync.WaitGroup`  for running `gorountine` .

```go
var wg sync.WaitGroup

func f1(i int)  {
   defer  wg.Done() 
   time.Sleep(time.Second * time.Duration(rand.Intn(3)))
   fmt.Println(i)
}

func main() {
   for i := 0; i < 10; i++ {
      wg.Add(1) // running one goroutine and add +1
      go f1(i)
   }

   wg.Wait() //waiting for all added goroutine finish
}
```

if we run this code many times , the order of number in output is also different. that is mean there are 10 `goroutines `  concurrently running .



### GOMAXPROCS

```go
var wg sync.WaitGroup

func a()  {
   defer wg.Done()
   for i := 0; i < 10; i++ {
      fmt.Printf("A:%d\n", i)
   }
}

func b()  {
   defer wg.Done()
   for i := 0; i < 10; i++ {
      fmt.Printf("B:%d\n", i)
   }
}

func main() {
   runtime.GOMAXPROCS(1)
   fmt.Println(runtime.NumCPU())
   wg.Add(2)
   go a()
   go b()
   wg.Wait()
}
```

`M:N`  share `M` goroutine to `N` OS.

### The relationship between Operating System Threads and Goroutines in Go:

- One operating system thread corresponds to multiple goroutines .
- Go programs can use multiple operating system threads at the same time
- There is a many-to-many relationship between goroutine and OS threads, that is `m:n`





# Channel

If `goroutine` is the concurrent execution body of Go program, `channel` is the connection between them. A `channel` is a communication mechanism that allows one goroutine to send a specific value to another `goroutine`.

`Channel` in Go language is a special type. The `channel` is like a conveyor belt or queue, always following the First In First Out (FIFO) rule to ensure the order of sending and receiving data. Each 	`channel` is a specific type of conduit, that is  when you declare the channel, you need to specify the element type for it.



```go
var b chan int //channel type format
```



```go
var b chan int
func main() {
   fmt.Println(b)      // nil
   b := make(chan int) // you need to initialize the channel
   fmt.Println(b)      // 0xc000094060
```



The declared `channel`needs to be initialized with the `make` function before it can be used.



## channel Operation:

`channel` has three Operations `send`, `receive` and `close`.

`<-`

1. send : `ch1 <- 1`
2. receive : `<- ch1`
3. close : `close(ch1)`



The thing to note about closing the channel is that the channel needs to be closed only when the receiver `goroutine` is notified that all the data has been sent. The channel can be recycled by the garbage collection mechanism. It is not the same as closing the file. It is necessary to close the file after the operation is completed, but it is not necessary to close the channel.



## What happens after closing a channel

1. after closing a channel and you need send to new a channel again , that can cause a  `panic`
2. Receiving a closed channel will keep getting the value until the channel is empty
3. Performing a receive operation on a channel that is closed and has no value will get the zero value of the corresponding type
4. already closed channel can not close again, that can cause a `panic`.

## take  value for Loop channel

After finished to send data to channel, we need to close the function by using `close`.

```go
var wg sync.WaitGroup

func f1(ch1 chan int)  {
   defer wg.Done()
   for i := 0; i < 100; i++ {
      ch1 <- i
   }
   close(ch1)
}

func f2(ch1, ch2 chan int)  {
   defer wg.Done()
   for {
      x, ok := <- ch1
      if !ok {
         break
      }
      ch2 <- x * x
   }
   close(ch2)
}

func main() {
   a := make(chan int, 50)
   b := make(chan int, 100)
   wg.Add(2)
   go f1(a)
   go f2(a, b)
   wg.Wait()
   for ret := range b {
      fmt.Println(ret)
   }
}
```

 