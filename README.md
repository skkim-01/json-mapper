# json-mapper

### Package Name : JsonMapper

#### type definitions
##### JsonMap struct
```
  type JsonMap struct {
    m map[string]interface{}

    // for finder/remover
    splitKey []string
    cursor   int

    // for adder
    insertKey   string
    insertValue interface{}
  }
```
  
##### value type definitions
```
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
```
  /*** Token for key depth ***/
  var SPLIT_TOKEN string = "."
```
  
#### APIs
##### NewBytes() : return JsonMap struct from json bytes
```
  func NewBytes(b []byte) (*JsonMap, error)
```

##### JsonMap.PPrint() : return json string with indent
```  
  func (j *JsonMap) PPrint() string
```

##### JsonMap.Print() : return trimmed json string 
```
  func (j *JsonMap) Print() string
```

##### JsonMap.Find() : find value from json key.
```
  func (j *JsonMap) Find(k string) interface{}
```

##### JsonMap.Remove() : remove value from key. prevent remove root.
```
  func (j *JsonMap) Remove(k string)
```
	
##### JsonMap.Insert() : insert/update value.
```
  func (j *JsonMap) Insert(base, k string, v interface{})
```
