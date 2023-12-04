package JsonMapper

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type YamlObject struct {
	m map[interface{}]interface{}
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

func newYamlObject(data map[interface{}]interface{}) *YamlObject {
	o := &YamlObject{
		m: data,
	}
	return o
}

func (o *YamlObject) Print() (string, error) {
	bytesOutput, err := yaml.Marshal(o.m)
	if err != nil {
		fmt.Println("#ERROR\t yamlObject.print:\t", err)
		return "", err
	}
	return string(bytesOutput[:]), nil
}
