# Go Routine Channel

## Generators 

* Generate a stream of data on a channel. 

* In concept of a pipeline. 

* Generators are going to function, generate infinite of data which comes from another function


* This code will helps to generate multiple numbers infinitely max of 500,000,000 


```
    func repeatFunc[T any, K any](done <- chan K, fn func() T) <-chan T {
        stream := make(chan T)
        go func() {
            defer close(stream)
            for {
                select{
                case <- done:
                    return
                case stream <- fn():
                }
            }
        }()
        return stream
    }

    func main() { 
        done := make(chan int)
        defer close(done)

        randNumFetcher := func() int { return rand.Intn(500000000)}

        for rando := range repeatFunc(done, randNumFetcher) {
            fmt.Println(rando)
        }
    }
```

## Output 

```
    337972534
    365975293
    29102738
    189378798
    19647910
    ...
    ...
```

* The above program will works with any function is not void. (No return value)

* We are returning the stream, Go routine will continously send the data to stream. 


## Pipelines and Pipelines Stages. 

* Need to add more here 


<p align="center">
  <img src="./pdf/go-routines-pipelines.png" alt="go routine pipelines" width="400">
</p>


* Find Prime will take more time, we need to add Find Prime instances concurrently. That will make the processing find prime much faster. 

```
/workspaces/System-Designs/generators (main) $ go run main.go 
64295783
332125033
206600791
282137437
144595499
469175467
64716143
377074541
298479791
28737581
53.322659849s
```

* Here it will take about 54 seconds. 

