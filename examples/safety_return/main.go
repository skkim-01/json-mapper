package main

import (
	"fmt"

	"github.com/skkim-01/json-mapper/examples/data"
	JsonMapper "github.com/skkim-01/json-mapper/src"
)

func main() {
	fmt.Println("#DBG \t #T1 \t Create source object")
	jsource, err := JsonMapper.NewString(data.TEST_JSON_TWO)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("#DBG \t #T2 \t Find()")
	fmt.Println("\033[32mFind() returns interface{} type\033[0m")
	// The code below throws a runtime error.
	// tmp := jsource.Find("vmap")
	// var tmap map[string]interface{}
	// tmap = make(map[string]interface{})
	// tmap = tmp		// compile error: type assertion

	fmt.Println("#DBG \t #T3 \t FindBool()")
	bValue := jsource.FindBool("vbool", false)
	fmt.Println("\033[32mvbool:", bValue, "\033[0m")

	fmt.Println("#DBG \t #T4 \t FindString()")
	strValue := jsource.FindString("vstr", "")
	fmt.Println("\033[32mvstr:", strValue, "\033[0m")

	fmt.Println("#DBG \t #T5 \t FindInt()")
	nValue := jsource.FindInt("vint64", 0)
	fmt.Println("\033[32mvint64:", nValue, "\033[0m")

	fmt.Println("#DBG \t #T6 \t FindInt32()")
	i32Value := jsource.FindInt32("vint64", 0)
	fmt.Println("\033[32mvint64:", i32Value, "\033[0m")

	fmt.Println("#DBG \t #T7 \t FindInt64()")
	i64Value := jsource.FindInt64("vint64", 0)
	fmt.Println("\033[32mvint64:", i64Value, "\033[0m")

	fmt.Println("#DBG \t #T8 \t FindFloat32()")
	f32Value := jsource.FindFloat32("vfloat64", 0)
	fmt.Println("\033[32mvfloat64:", f32Value, "\033[0m")

	fmt.Println("#DBG \t #T9 \t FindFloat64()")
	f64Value := jsource.FindFloat64("vfloat64", 0)
	fmt.Println("\033[32mvfloat64:", f64Value, "\033[0m")

	fmt.Println("#DBG \t #T10 \t FindMap()")
	mapValue := jsource.FindMap("vmap", nil)
	fmt.Println("\033[32mvmap:", mapValue, "\033[0m")

	fmt.Println("#DBG \t #T10 \t FindSilce()")
	slMap := jsource.FindSilce("vslice", nil)
	fmt.Println("\033[32mvslice:", slMap, "\033[0m")
}
