package main

import (
	"fmt"

	"github.com/skkim-01/json-mapper/examples/data"
	JsonMapper "github.com/skkim-01/json-mapper/src"
)

func main() {
	fmt.Println("#DBG \t #T1 \t Create from bytes")
	fmt.Println("  - The test cases below all include this function, so this test is skipped.")
	fmt.Println("---")
	fmt.Println("")

	fmt.Println("#DBG \t #T2 \t Create from string")
	j1, err := JsonMapper.NewString(data.TEST_JSON_ONE)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(j1.PPrint())
	fmt.Println("---")
	fmt.Println("")

	fmt.Println("#DBG \t #T3 \t Create from object")
	j2, err := JsonMapper.NewObject(j1.Find(""))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(j2.PPrint())
	fmt.Println("---")
	fmt.Println("")

	fmt.Println("#DBG \t #T4 \t Create from json file")
	j3, err := JsonMapper.NewFile("../data/test_json_one.json")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(j3.PPrint())
}
