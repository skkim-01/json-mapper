package JsonMapper

import (
	"fmt"
	"reflect"
	"strconv"
)

// adder_search_root : add recursive function start point
func (j *JsonMap) adder_search_root() {
	var currentKey string
	currentKey = j.splitKey[j.cursor]

	// goto: switch type for next key
	// 	 - this section : find basekey & update exist key's value
SWITCH_TYPE:
	switch reflect.TypeOf(j.m[currentKey]) {
	case JsonMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			j.m[j.insertKey] = j.insertValue
			return
		}
		j.adder_search_map_r(j.m[currentKey].(map[string]interface{}))
		return

	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			j.m[j.insertKey] = j.insertValue
			return
		}
		// slice must set return value
		j.m[currentKey] = j.adder_search_slice_r(j.m[currentKey].([]interface{}))
		return

		// just for slice map
	case SliceMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			j.m[j.insertKey] = j.insertValue
			return
		}
		// slice must set return value
		j.m[currentKey] = j.adder_search_slice_r(j.m[currentKey].([]interface{}))
		return

		// for update value
	case Float64Type, StringType, BoolType, IntType:
		j.m[j.insertKey] = j.insertValue
		return

	default:
		// set next key
		if j.cursor < len(j.splitKey)-1 {
			j.cursor = j.cursor + 1
			currentKey = currentKey + "." + j.splitKey[j.cursor]
			goto SWITCH_TYPE
		} else {
			j.m[j.insertKey] = j.insertValue
		}
	}
}

// adder_search_slice_r : adder recursive function for slice
// 		- slice must have slice return value for set new slice
func (j *JsonMap) adder_search_slice_r(sub []interface{}) []interface{} {
	// slice key is always integer
	var currentKey int
	var err error

	currentKey, err = strconv.Atoi(j.splitKey[j.cursor])
	if nil != err {
		fmt.Println("### ERROR ### :", err)
		return sub
	}

	// index out of range: always append value
	if currentKey >= len(sub) || currentKey < 0 {
		sub = append(sub, j.insertValue)
		return sub
	}

	switch reflect.TypeOf(sub[currentKey]) {
	case JsonMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub[currentKey] = j.insertValue
			return sub
		}
		j.adder_search_map_r(sub[currentKey].(map[string]interface{}))
		return sub

		// --- must not hit this section. json doesn't allow netsted array.
	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub[currentKey] = j.insertValue
			return sub
		}
		// slice must set return value
		sub[currentKey] = j.adder_search_slice_r(sub[currentKey].([]interface{}))
		return sub
		// --- end of section

	case Float64Type, StringType, BoolType, IntType:
		sub[currentKey] = j.insertValue
		return sub

	default:
		// it must be error. must not hit this section.
		fmt.Println("### ERROR ### default: undefined error")
	}
	return sub
}

// adder_search_map_r : adder recursive function for map
func (j *JsonMap) adder_search_map_r(sub map[string]interface{}) {
	var currentKey string
	currentKey = j.splitKey[j.cursor]

	// goto: switch type for next key
SWITCH_TYPE:
	switch reflect.TypeOf(sub[currentKey]) {
	case JsonMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub[j.insertKey] = j.insertValue
			return
		}
		j.adder_search_map_r(sub[currentKey].(map[string]interface{}))
		return

	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub[j.insertKey] = j.insertValue
			return
		}
		sub[currentKey] = j.adder_search_slice_r(sub[currentKey].([]interface{}))
		return

	case Float64Type, StringType, BoolType, IntType:
		sub[j.insertKey] = j.insertValue
		return

	default:
		// set next key
		if j.cursor < len(j.splitKey)-1 {
			j.cursor = j.cursor + 1
			currentKey = currentKey + "." + j.splitKey[j.cursor]
			goto SWITCH_TYPE
		} else {
			sub[j.insertKey] = j.insertValue
		}
	}
}
