package JsonMapper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type YamlMap struct {
	m []YamlObject
}

// yamltype
var YamlMapType = reflect.TypeOf((map[interface{}]interface{})(nil))

func NewYamlMap(strYaml string) (*YamlMap, error) {
	reader := strings.NewReader(strYaml)
	res, err := _decode(reader)
	if err != nil {
		return nil, err
	}
	y := &YamlMap{
		m: res,
	}
	return y, nil
}

func NewYamlMapFile(yamlFile string) (*YamlMap, error) {
	pFile, err := os.Open(yamlFile)
	if err != nil {
		return nil, err
	}
	res, err := _decode(pFile)
	if err != nil {
		return nil, err
	}
	y := &YamlMap{
		m: res,
	}
	return y, nil
}

func NewYamlObject(o interface{}) (*YamlMap, error) {
	yamlBytes, err := yaml.Marshal(o)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(yamlBytes)
	res, err := _decode(reader)
	if err != nil {
		return nil, err
	}
	y := &YamlMap{
		m: res,
	}
	return y, nil
}

func _decode(r io.Reader) ([]YamlObject, error) {
	yamlMaps := make([]YamlObject, 0)
	d := yaml.NewDecoder(r)
	for {
		tmp := make(map[interface{}]interface{}, 0)
		err := d.Decode(&tmp)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("#ERROR\t YamlMap._decode:\t", err)
			return nil, err
		}
		yamlMaps = append(yamlMaps, *newYamlObject(tmp))
	}
	return yamlMaps, nil
}

func (o *YamlMap) Prints() string {
	retv := ""
	for _, v := range o.m {
		retv += "---\n"
		output, err := v.Print()
		if err != nil {
			fmt.Println("#ERROR\t YamlMap.Print:\t", err)
			break
		}
		retv += output
	}
	return retv
}

func (o *YamlMap) Print(idx int) string {
	if retv, err := o.m[idx].Print(); err == nil {
		return retv
	} else {
		return err.Error()
	}
}

func (o *YamlMap) Find(idx int, k string) interface{} {
	if idx < 0 {
		return nil
	} else if idx >= len(o.m) {
		return nil
	} else if k == "" {
		return o.m[idx]
	}

	o.m[idx].splitKey = strings.Split(k, SPLIT_TOKEN)
	o.m[idx].cursor = 0

	v := o.m[idx].finder_search_root()
	return v
}

// Insert : insert/update, when insert root, set [base == ""]
func (o *YamlMap) Insert(idx int, base, k string, v interface{}) {
	if idx < 0 {
		return
	} else if idx >= len(o.m) {
		return
	}

	o.m[idx].splitKey = strings.Split(base, SPLIT_TOKEN)
	o.m[idx].splitKey = append(o.m[idx].splitKey, k)
	o.m[idx].cursor = 0

	o.m[idx].insertKey = k
	// type cast : []map[string]interface{} -> []interface{}
	vTmp := make([]interface{}, 0)
	switch reflect.TypeOf(v) {
	case SliceMapType:
		for i := range v.([]map[interface{}]interface{}) {
			vTmp = append(vTmp, v.([]map[interface{}]interface{})[i])
		}
		o.m[idx].insertValue = vTmp

	default:
		o.m[idx].insertValue = v
	}

	o.m[idx].adder_search_root()
}

// Remove
func (o *YamlMap) Remove(idx int, k string) {
	if idx < 0 {
		return
	} else if idx >= len(o.m) {
		return
	}

	// prevent remove root
	if k == "" {
		return
	}

	o.m[idx].splitKey = strings.Split(k, SPLIT_TOKEN)
	o.m[idx].cursor = 0
	o.m[idx].insertKey = k

	o.m[idx].remover_search_root()
}
