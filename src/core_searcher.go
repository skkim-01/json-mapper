package JsonMapper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	JsonConditionParser "github.com/skkim-01/json-condition-parser"
)

// searcher_search_root : search recursive function start point
func (j *JsonMap) searcher_search_root() interface{} {
	j.searchResult = make([]map[string]interface{}, 0)

	for k, v := range j.m {
		switch reflect.TypeOf(v) {
		case JsonMapType:
			j.searcher_search_map_r(k, v.(map[string]interface{}))
			break
		case SliceType:
			j.searcher_search_slice_r(k, v.([]interface{}))
			break
		case Float64Type, StringType, BoolType, IntType:
			j.hit("", k, v)
			break
		default:
			// unnecessary case: full-scan
		}
	}

	return nil
}

// searcher_search_slice_r
func (j *JsonMap) searcher_search_slice_r(parentKey string, sub []interface{}) {
	for idx, v := range sub {
		switch reflect.TypeOf(v) {
		case JsonMapType:
			j.searcher_search_map_r(
				fmt.Sprintf("%v.%v", parentKey, idx),
				v.(map[string]interface{}),
			)
			break
		case SliceType:
			// --- must not hit this section. json doesn't allow netsted array.
			j.searcher_search_slice_r(
				fmt.Sprintf("%v.%v", parentKey, idx),
				v.([]interface{}),
			)
			break
		case Float64Type, StringType, BoolType, IntType:
			j.hit(parentKey, strconv.Itoa(idx), v)
			break
		default:
			// unnecessary case: full-scan
		}
	}
}

// searcher_search_map_r
func (j *JsonMap) searcher_search_map_r(parentKey string, sub map[string]interface{}) {
	for k, v := range sub {
		switch reflect.TypeOf(v) {
		case JsonMapType:
			j.searcher_search_map_r(
				fmt.Sprintf("%v.%v", parentKey, k),
				v.(map[string]interface{}),
			)
		case SliceType:
			j.searcher_search_slice_r(
				fmt.Sprintf("%v.%v", parentKey, k),
				v.([]interface{}),
			)
		case Float64Type, StringType, BoolType, IntType:
			j.hit(parentKey, k, v)
			break
		default:
			// unnecessary case: full-scan
		}
	}
}

func (j *JsonMap) operation(param interface{}) bool {
	// hit: true, else false
	return JsonConditionParser.ParseMap(param, j.searchOpt)
}

func (j *JsonMap) hit(parentKey, currentKey string, currentValue interface{}) {
	if strings.Contains(j.searchKey, currentKey) {
		if JsonConditionParser.ParseMap(currentValue, j.searchOpt) {
			// store k, v
			tmp := make(map[string]interface{})
			if len(parentKey) == 0 {
				tmp[fmt.Sprintf("%v", currentKey)] = currentValue
			} else {
				tmp[fmt.Sprintf("%v.%v", parentKey, currentKey)] = currentValue
			}
			// add parentkey
			tmp["parentKey"] = parentKey
			j.searchResult = append(j.searchResult, tmp)
		}
	}
}
