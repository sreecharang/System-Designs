# Generics 

* We have variables, different types like this. 

```
numbers1 := []int{1, 2, 3, 4, 5}
numbers2 := []int32{1, 2, 3, 4, 5}
numbers3 := []int64{1, 2, 3, 4, 5}
numbers4 := []float32{1.1, 2.1, 3.1, 4.1, 5.1}
numbers5 := []float64{1.1, 2.1, 3.1, 4.1, 5.1}

```

* we will write separate functions, by summing each of the slices. 

* Here we will essential to create five of these function, in order to add all the values. For example: 


```
    func sumNumbers(numbers []float64) float64 {
        var result float64
        for i := range numbers {
            result += numbers[i]
        }
        return result 
    }
```

* Here we will get code duplication. 

* To avoid this generic will come in place, Generic number type would be expected as opposed to a specific number type such as int64 or int32. 

* we can could do that by adding the type parameter in square brackets next to the name of the function. 

* This is boiler plate for generic. 


```
    func sumNumbers(numbers []int32) int32 {
        var result int32 
        for i := range numbers {
            result += numbers[i]
        }
        return result 
    }
```

```
    func sumNumbers[T Number](numbers []T) T {
        var result T 
        for i := range numbers {
            result += numbers[i]
        }
        return result
    }
```

```
    type Number interface {
        int | int64 | float64
    }

```

* Here type number is an interface with int | int32 | int64 | float32 | float64 

* That's generics will help to eliminate the necessity of writing the same function multiple times to account for each specific type. 

## Final output of the function looks as: 

```
    type Number interface {
        int | int64 | float64
    }


    func sumNumbers[T Number](numbers []T) T {
        var result T
        for i := range numbers {
            result += numbers[i]
        }
        return result
    }

    func main() {
        numbers := []int{1, 2, 3, 4, 5}
        result := sumNumbers(numbers)
        fmt.Println("The sum of the number is: ", result)
    }
```

### Output 

```
    The sum of the number is:  15
```

