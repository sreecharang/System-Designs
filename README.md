# System-Designs
System design for the application. [Systems Designs](./systems-design.md)

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

    func processData(wg *sync.WaitGroup, <span style="background-color: #01C8FE; border-radius: 4px;">result *[]int</span>, data) { // In Confinement, result parameter will get replaced
    
         
        defer wg.Done() 
        processData := process(data)

        <span style="background-color: #01C8FE; border-radius: 4px;">lock.Lock()</span>    // In Confinement, This will be removed.
        

        *result = <span style="background-color: #01C8FE; border-radius: 4px;">append</span> (*result, processData) // In Confinement, append parameter will get replaced.

        <span style="background-color: #01C8FE; border-radius: 4px;">lock.Unlock()</span>  // In Confinement, This will be removed. 
    }

    func main() {
        start := time.Now()

        var wg sync.WaitGroup
        input := []int{1, 2, 3, 4, 5}

        <span style="background-color: #01C8FE; border-radius: 4px;">result := []int{}</span> // In Confinement, This part will get replaced. 

        for _, data := range input {
            wg.Add(1) 
            go processData(&wg, <span style="background-color: #01C8FE; border-radius: 4px;">&result</span> , data) // In Confinement, result parameter will get changed. 
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

* 


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
