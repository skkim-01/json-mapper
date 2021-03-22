package JsonMapper

import (
	"bytes"
	"encoding/json"
)

// prettyPrint : pretty-print byte to json string
func prettyPrint(src []byte) (string, error) {
	var dst bytes.Buffer
	err := json.Indent(&dst, src, "", "  ")
	if nil != err {
		return "", err
	}
	return dst.String(), nil
}
