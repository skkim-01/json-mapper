package JsonMapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// NewBytes : new jmap from json bytes
func NewBytes(b []byte) (*JsonMap, error) {
	j := &JsonMap{
		m: make(map[string]interface{}),
	}
	err := FromJson(b, &j.m)
	return j, err
}

// PPrint : pretty print
func (j *JsonMap) PPrint() string {
	b, _ := ToJson(j.m)
	str, _ := prettyPrint(b)
	return str
}

// Print : print
func (j *JsonMap) Print() string {
	b, _ := ToJson(j.m)
	return string(b)
}

// Find : find value from json key
func (j *JsonMap) Find(k string) interface{} {
	if k == "" {
		return j.m
	}

	j.splitKey = strings.Split(k, SPLIT_TOKEN)
	j.cursor = 0

	v := j.finder_search_root()

	return v
}

// Remove : remove value from key. prevent remove root.
func (j *JsonMap) Remove(k string) {
	if k == "" {
		fmt.Println("### ERROR ### cannot delete root")
		return
	}

	j.splitKey = strings.Split(k, SPLIT_TOKEN)
	j.cursor = 0

	j.remover_search_root()
}

// Insert : insert/update, when insert root, set [base == ""]
func (j *JsonMap) Insert(base, k string, v interface{}) {
	j.splitKey = strings.Split(base, SPLIT_TOKEN)
	j.splitKey = append(j.splitKey, k)
	j.cursor = 0

	j.insertKey = k
	// type cast : []map[string]interface{} -> []interface{}
	vTmp := make([]interface{}, 0)
	switch reflect.TypeOf(v) {
	case SliceMapType:
		for i := range v.([]map[string]interface{}) {
			vTmp = append(vTmp, v.([]map[string]interface{})[i])
		}
		j.insertValue = vTmp
		break

	default:
		j.insertValue = v
	}

	j.adder_search_root()
}

// ToJson : object(struct) to json bytes
func ToJson(_o interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(_o)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// FromJson : json bytes to object(struct)
func FromJson(_byte []byte, _o interface{}) error {
	err := json.Unmarshal(_byte, &_o)
	return err
}
