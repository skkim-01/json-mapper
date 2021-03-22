package JsonMapper

import (
	"bytes"
	"encoding/json"
)

// toJson : object(struct) to json bytes
func toJson(_o interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(_o)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// fromJson : json bytes to object(struct)
func fromJson(_byte []byte, _o interface{}) error {
	err := json.Unmarshal(_byte, &_o)
	return err
}

// prettyPrint : pretty-print byte to json string
func prettyPrint(src []byte) (string, error) {
	var dst bytes.Buffer
	err := json.Indent(&dst, src, "", "  ")
	if nil != err {
		return "", err
	}
	return dst.String(), nil
}
