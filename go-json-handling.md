# GO JSON Handiling 

## Alternatives to JSON. 

* YAML, TOML, INI 

* BSON, MessagePack, CBOR, Smile 

* XML 

* ProtoBuf 

* Custom/proprietary formats 


## Marshaling JSON 

* Creating JSON from a Go object is (usually) very straight forward: 

```
    func main() {
        x := map[string]string {
            "foo": "bar",
        }

        data, _ := json.Marshal(x)
        fmt.Println(string(data))
    }
```

## Marshaling JSON, #2 

```
    func main() {
        type person struct {
            Name string `json: "name"`
            Age int `json: "age"`
            Description string `json:"descr, omitempty"`
            secret string // unexported fields are never (un)marshaled
        }

        x := person {
            Name: "Bob", 
            Age: "32", 
            secret: "Shh!" 
        }
        data, _ := json.Marshal(x)
        fmt.Println(string(data))
    }
```

* unexported exclude as unmarshal in the json. 

## Unmarshalling JSON 

* Unmarshalling JSON is often a bit trickier 

```
    func main() {
        data := []byte(`{"foo":"bar"}`)
        var x interface{} 
        _ = json.Unmarshal(data, &x)
        spew.Dump(x)
    }
```

## Unmarshalling JSON 

* Avoid a amap whenever possible. interface{} says nothing. 

```
    func main() {
        type person struct {
            Name string `json:"name"`
            Age int `json:"age"`
            Description string `json:"descr, omitempty"`
            secret string // Unexported fields are never (un)marshaled 
        }

        data := []byte(`{"name": "Bob", "age": 32, "secret": "Shh!"}`)
        var x person 
        _ := json.Unmarshal(data, &x)
        spew.Dump(x)
    }
```


## Cases of "unknown input" 

* Input may be a string or Number 

123 
"123" 

* Input may be object, or array of objects 

{...}
[{...}, {...}]

* Input may be success, or an error 

{"success":true, "results":[...]}
{"success":false, "error":"..."}

