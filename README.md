# json-mapper

### Package Name : JsonMapper

#### type definitions
##### JsonMap struct
```go
  type JsonMap struct {
    m map[string]interface{}

    // for finder/remover
    splitKey []string
    cursor   int

    // for adder
    insertKey   string
    insertValue interface{}

    // for searcher
    searchOpt    map[string]interface{}
    searchKey    string
    searchResult []map[string]interface{}
  }
```
  
##### value type definitions
```go
  /*** reflect type definition ***/
  var IntType = reflect.TypeOf(1)
  var Float64Type = reflect.TypeOf(float64(1))
  var StringType = reflect.TypeOf(string(""))
  var BoolType = reflect.TypeOf(false)
  var SliceType = reflect.TypeOf([]interface{}(nil))
  var SliceMapType = reflect.TypeOf([]map[string]interface{}(nil))
  var JsonMapType = reflect.TypeOf((map[string]interface{})(nil))
```

##### split token
```go
  /*** Token for key depth ***/
  var SPLIT_TOKEN string = "."
```
  
#### APIs
##### NewBytes() : return JsonMap struct from json bytes
```go
  func NewBytes(b []byte) (*JsonMap, error)
```

##### JsonMap.PPrint() : return json string with indent
```go
  func (j *JsonMap) PPrint() string
```

##### JsonMap.Print() : return trimmed json string 
```go
  func (j *JsonMap) Print() string
```

##### JsonMap.Find() : find value from json key.
```go
  func (j *JsonMap) Find(k string) interface{}
```

##### JsonMap.Remove() : remove value from key. prevent remove root.
```go
  func (j *JsonMap) Remove(k string)
```
	
##### JsonMap.Insert() : insert/update value.
```go
  func (j *JsonMap) Insert(base, k string, v interface{})
```

##### JsonMap.Search() :  Retrieve all values ​​matching a condition.
search options: https://github.com/skkim-01/json-condition-parser/tree/main#type-definitions
```go
  func (j *JsonMap) Search(keyIncluded string, searchOpt string) []map[string]interface{}
```
