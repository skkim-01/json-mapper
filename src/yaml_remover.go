package JsonMapper

import (
	"fmt"
	"reflect"
	"strconv"
)

// finder_search_root : finder recursive function start point
func (j *YamlObject) remover_search_root() {
	var currentKey string
	currentKey = j.splitKey[j.cursor]

	// goto: switch type for next key
SWITCH_TYPE:
	switch reflect.TypeOf(j.m[currentKey]) {
	case YamlMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			delete(j.m, currentKey)
			return
		}
		j.remover_search_map_r(j.m[currentKey].(map[interface{}]interface{}))
		return

	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			delete(j.m, currentKey)
			return
		}
		j.m[currentKey] = j.remover_search_slice_r(j.m[currentKey].([]interface{}))
		return

	case Float64Type, StringType, BoolType, IntType:
		delete(j.m, currentKey)
		return

	default:
		// set next key
		if j.cursor < len(j.splitKey)-1 {
			j.cursor = j.cursor + 1
			currentKey = currentKey + "." + j.splitKey[j.cursor]
			goto SWITCH_TYPE
		}
	}
}

// finder_search_slice_r : finder recursive function for slice
func (j *YamlObject) remover_search_slice_r(sub []interface{}) []interface{} {
	// slice key is always integer
	var currentKey int
	var err error

	currentKey, err = strconv.Atoi(j.splitKey[j.cursor])
	if nil != err {
		fmt.Println("### ERROR ### :", err)
		return sub
	}

	if currentKey >= len(sub) || currentKey < 0 {
		fmt.Println("### ERROR ### : index is out of range")
		return sub
	}

	switch reflect.TypeOf(sub[currentKey]) {
	case YamlMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub = j.removeSliceElement(sub, currentKey)
			return sub
		}
		j.remover_search_map_r(sub[currentKey].(map[interface{}]interface{}))
		return sub

		// --- must not hit this section. json doesn't allow netsted array.
	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			sub = j.removeSliceElement(sub, currentKey)
			return sub
		}
		sub[currentKey] = j.remover_search_slice_r(sub[currentKey].([]interface{}))
		return sub

	case Float64Type, StringType, BoolType, IntType:
		sub = j.removeSliceElement(sub, currentKey)
		return sub
		// --- end of section

	default:
		// it must be error, must not hit
		fmt.Println("### LOG ### default: undefined error")
		return nil
	}
}

// finder_search_map_r : finder recursive function for map
func (j *YamlObject) remover_search_map_r(sub map[interface{}]interface{}) {
	var currentKey string
	currentKey = j.splitKey[j.cursor]

	// goto: switch type for next key
SWITCH_TYPE:
	switch reflect.TypeOf(sub[currentKey]) {
	case YamlMapType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			delete(sub, currentKey)
			return
		}
		j.remover_search_map_r(sub[currentKey].(map[interface{}]interface{}))
		return

	case SliceType:
		j.cursor = j.cursor + 1
		if j.cursor >= len(j.splitKey) {
			delete(sub, currentKey)
			return
		}
		sub[currentKey] = j.remover_search_slice_r(sub[currentKey].([]interface{}))
		return

	case Float64Type, StringType, BoolType, IntType:
		delete(sub, currentKey)
		return

	default:
		// set next key
		if j.cursor < len(j.splitKey)-1 {
			j.cursor = j.cursor + 1
			currentKey = currentKey + "." + j.splitKey[j.cursor]
			goto SWITCH_TYPE
		}
	}
}

// delete slice element
func (j *YamlObject) removeSliceElement(s []interface{}, index int) []interface{} {
	copy(s[index:], s[index+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	return s
}
