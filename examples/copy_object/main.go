package main

import (
	"fmt"

	"github.com/skkim-01/json-mapper/examples/data"
	JsonMapper "github.com/skkim-01/json-mapper/src"
)

func main() {
	fmt.Println("#DBG \t #T1 \t Create source object")
	jsource, err := JsonMapper.NewString(data.TEST_JSON_ONE)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("#DBG \t #T2 \t Get source object")
	sourceObject := jsource.GetObject()

	fmt.Println("#DBG \t #T3 \t Get copy object")
	copyObject := jsource.CopyObject()

	fmt.Println("#DBG \t #T4 \t Update map")
	jsource.Insert("", "root.int", 1)
	jsource.Insert("", "root.bool", true)

	fmt.Println("#DBG \t #T5 \t Print objects")
	fmt.Println("     \t     \t - GetObject(): The \"map[string]interface{}\" value will be changed when update/delete value")
	fmt.Println("     \t     \t - CopyObject(): The \"map[string]interface{}\"value will not be changed when update/delete value")
	fmt.Println("### json string")
	fmt.Println(jsource.PPrint())
	fmt.Println("### source object")
	fmt.Println(sourceObject)
	fmt.Println("### copied object")
	fmt.Println(copyObject)
}
