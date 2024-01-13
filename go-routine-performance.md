

# Go-Concurrency

* Multiple lines of code need to access the Critical section of the code. 

* Current go routines, with no synchronized place to ensure that the shared resources access the one go routine at a time. 

* Then we open door for data in consistency and Race condition. 



* To solve this, Syncrnoized Go routine acces the value. 

* Simple way is guard the execution Mutual Exclusion (mutex)   


## Mutex 

* Mutex is a lock, Lock of critical section, other go routine need to wait till the lock to be unlocked. 

* This will create a bottle neck to program. 

* Code follows as it is: 

<pre>
<code>
    func process(data int) int{
        time.Sleep(time.Second * 2)
        return data * 2
    }

    func processData(wg *sync.WaitGroup, result *[]int, data) { // In Confinement, result parameter will get replaced
    
         
        defer wg.Done() 
        processData := process(data)

        lock.Lock()   // In Confinement, This will be removed.
        

        *result = append (*result, processData) // In Confinement, append parameter will get replaced.

        lock.Unlock() // In Confinement, This will be removed. 
    }

    func main() {
        start := time.Now()

        var wg sync.WaitGroup
        input := []int{1, 2, 3, 4, 5}

        result := []int{} // In Confinement, This part will get replaced. 

        for _, data := range input {
            wg.Add(1) 
            go processData(&wg, &result , data) // In Confinement, result parameter will get changed. 
        }

        wg.Wait() 
        fmt.Println(time.Since(start))
	    fmt.Println(result)
    }
</code>
</pre>

### Output 

```
    /workspaces/System-Designs/routines (main) $ go run main.go 
    2.001128458s
    [8 10 2 4 6]
```

## Confinement 

* Confinement will helps to program, to eliminate the Mutex for access the shared resources. 

* ProcessData Mutex

```
    func processData(wg *sync.WaitGroup, result *[]int, data) { 
    
         
        defer wg.Done() 
        processData := process(data)

        lock.Lock()   // In Confinement, This will be removed.
        

        *result = append (*result, processData) // In Confinement, append parameter will get replaced.

        lock.Unlock() // In Confinement, This will be removed. 
    }
```

* Multiple Go routines, Can able to see the result slide on the Mutex section. 

* Mutate the result slice. 



* Confinement, is basically the idea of ensuring information only available one concurrent go routine. 

* For example, Each individual go routine only needs to access a particular index of result, we are storing only index of slice. 


* ProcessData Confinement 

```
    input := []int{1, 2, 3, 4, 5}

    result := make([]int, len(input))

    for i, data := range input {
        wg.Add(1) 
        go processData(&wg, &result[i], data) 
    }

    func processData(wg *sync.WaitGroup, resultDest *int, data int){

        defer wg.Done()
        
        processData := process(data)
  
        *resultDest = processData
    }

```

* In code, we giving result slice size on length of the input. 

* Because, We need to specific indexes that we can assign the process to. 

* &result[i] is pointer to the result address. 

* We are giving everything, we just providing the index of specific memory address. so it alter only one index address. 

* Basically, Go routine can't access other indexes of slice. Go routine confine to specific index of the result. 

* Confinement, is an idea ensuring information only avaialble one concurrent go routine. 

* We don't need to use the synchronization for race condition, Since the goroutine can bale access the specific index of the result slice. 

Code snippet as follows: 

```

    func process(data int) int{
        time.Sleep(time.Second * 2)
        return data * 2
    }

    func processData(wg *sync.WaitGroup, resultDest *int, data int){

        defer wg.Done()
        
        processData := process(data)

        // lock.Lock()
        *resultDest = processData
        // lock.Unlock()
    }

    func main() {

        start := time.Now()

        var wg sync.WaitGroup

        input := []int{1, 2, 3, 4, 5}
        // result := []int{}
        result := make([]int, len(input))

        for i, data := range input {
            wg.Add(1) 
            go processData(&wg, &result[i], data) 
        }

        wg.Wait()

        fmt.Println(time.Since(start))
        fmt.Println(result)

    }

```

### Output 

```
    /workspaces/System-Designs/routines (main) $ go run main.go 
    2.000110163s
    [2 4 6 8 10] 
```

* Reference [reference](https://www.youtube.com/watch?v=Bk1c30avsuU&t=286s)