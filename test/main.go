package main

import (
	"fmt"

	JsonMapper "github.com/skkim-01/json-mapper/src"
)

func main() {
	fmt.Println("========= test case 1 ========")
	test_json_v1()

	fmt.Println()
	fmt.Println("========= test case 2 ========")
	test_json_v2()
}

/*** test case 1 ***/
var test_json_string string = `
{
	"id": "0001",
	"type": "donut",
	"name": "Cake",
	"ppu": 0.55,
	"batters":
		{
			"batter":
				[
					{ "id": "1001", "type": "Regular" },
					{ "id": "1002", "type": "Chocolate" },
					{ "id": "1003", "type": "Blueberry" },
					{ "id": "1004", "type": "Devil's Food" }
				]
		},
	"topping":
		[
			{ "id": "5001", "type": "None" },
			{ "id": "5002", "type": "Glazed" }			
		]
}
`

// test_json_v1 : test with "test_json_string"
func test_json_v1() {
	j, e := JsonMapper.NewBytes([]byte(test_json_string))
	if nil != e {
		fmt.Println(e)
		return
	}

	// ### TEST CASE 0
	fmt.Println("### TEST CASE 1: Print source")
	fmt.Println(j.PPrint())

	// ### TEST CASE 1
	k1 := "id"
	v1 := j.Find(k1)
	fmt.Println()
	fmt.Println("### TEST CASE 1: find root")
	fmt.Println(k1, ":", v1)

	// ### TEST CASE 2
	k2 := "batters.batter.0.id"
	v2 := j.Find(k2)
	fmt.Println()
	fmt.Println("### TEST CASE 2: find sub")
	fmt.Println(k2, ":", v2)

	// ### TEST CASE 3
	newBatter := make(map[string]interface{}, 0)
	newBatter["id"] = "1005"
	newBatter["type"] = "json mapper"
	j.Insert("batters.batter", "-1", newBatter)
	fmt.Println("### TEST CASE 3: insert value at the end of slice")
	fmt.Println(newBatter)

	// ### TEST CASE 3
	j.Remove("id")
	j.Remove("type")
	j.Remove("name")
	j.Remove("ppu")
	j.Remove("topping")

	// WARNING : remove will change index
	fmt.Println()
	fmt.Println("### TEST CASE 3: remove slice")
	j.Remove("batters.batter.0") // remove id : 1001
	fmt.Println(j.Find("batters.batter"))
	j.Remove("batters.batter.1") // remove id : 1003
	fmt.Println(j.Find("batters.batter"))
	j.Remove("batters.batter.0") // remove id : 1002
	fmt.Println(j.Find("batters.batter"))
	j.Remove("batters")

	// ### TEST CASE 4 : clean map
	fmt.Println()
	fmt.Println("### TEST CASE 4: clean map")
	fmt.Println(j.PPrint())
}

/*** test case 2 ***/
var test_json_string_2 string = `
{
	"root.int" : 0,
	"root.string" : "root_value",
	"root.bool" : true,
	"child.1.map" : {
		"child.1.subint" : 1,
		"child.1.substring" : "child_1_value",
		"child.1.subslice" : [
			{ "id": "5001", "type": "None" },
			{ "id": "5002", "type": "Glazed" },
			{ "id": "5005", "type": "Sugar" }
		],
		"child.1.submap" : {
			"child.1.1.subint" : 2,
			"child.1.1.substring" : "child_1_1_value"
		}
	}
}
`

// test_json_v2 : key contains dot(.)
func test_json_v2() {
	j, e := JsonMapper.NewBytes([]byte(test_json_string_2))
	if nil != e {
		fmt.Println(e)
		return
	}

	fmt.Println("### TEST CASE 1: Find Values")
	fmt.Println("root.int :", j.Find("root.int"))
	fmt.Println("child.1.map.child.1.subint :", j.Find("child.1.map.child.1.subint"))
	fmt.Println("child.1.map.child.1.subslice.2.type :", j.Find("child.1.map.child.1.subslice.2.type"))
	fmt.Println("child.1.map.child.1.submap.child.1.1.subint :", j.Find("child.1.map.child.1.submap.child.1.1.subint"))

	fmt.Println()
	fmt.Println("### TEST CASE 2: Insert value")
	j.Insert("child.1.map.child.1.submap", "insert_value", "new value")
	j.Insert("", "root.new.int.value", 100)
	newValue := make(map[string]interface{}, 0)
	newValue["id"] = "1005"
	newValue["type"] = "json mapper"
	j.Insert("child.1.map.child.1.subslice", "-1", newValue)

	fmt.Println()
	fmt.Println("### TEST CASE 3: print inserted")
	fmt.Println(j.PPrint())

	fmt.Println()
	fmt.Println("### TEST CASE 4: remove")
	j.Remove("child.1.map.child.1.subslice.0")
	j.Remove("child.1.map.child.1.subslice.0")
	j.Remove("child.1.map.child.1.subslice.0")
	j.Remove("child.1.map.child.1.subslice.0")
	j.Remove("child.1.map.child.1.subslice.0") // out of index
	j.Remove("child.1.map.child.1.subslice")
	j.Remove("child.1.map.child.1.submap.child.1.1.subint")
	j.Remove("root.bool")
	j.Remove("root.int")
	j.Remove("root.string")

	fmt.Println()
	fmt.Println("### TEST CASE 5: print result")
	fmt.Println(j.PPrint())
}
