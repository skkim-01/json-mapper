package JsonMapper

import (
	"errors"
	"fmt"
	"io"
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
	fmt.Println("#DBG\t NewYamlMap.length:\t", len(res))
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
		fmt.Println(output)
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
	fmt.Println(v)
	return v
}
