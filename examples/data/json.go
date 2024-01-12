package data

var TEST_JSON_ONE string = `{
	"root.int" : 0,
	"child.1.map" : {
		"child.1.subint" : 1,
		"child.1.subslice" : [
			{ "id": "5001", "type": "None" }
		],
		"child.1.submap" : {
			"child.1.1.substring" : "child_1_1_value"
		}
	}
}
`

var TEST_JSON_TWO string = `{
	"vint": 10,
	"vint64": 1024248012043,
	"vfloat32": 3.141592,
	"vfloat64": 3.141582141582141582,
	"vstr": "It is string value// map[child.1.map:map[child.1.subint:1 child.1.submap:map[child.1.1.substring:child_1_1_value] child.1.subslice:[map[id:5001 type:None]]] root.bool:true root.int:1]",
	"vbool": true,
	"vmap" : {
		"child.1.submap" : {
			"child.1.1.substring" : "child_1_1_value"
		}
	},
	"vslice": [
		{ "id": "5001", "type": "None" },
		{ "id": "5002", "type": "type T" },
		{ "id": "5003", "type": "type C" }
	]
}
`
