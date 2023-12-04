package main

import (
	"fmt"

	JsonMapper "github.com/skkim-01/json-mapper/src"
)

func main() {

	test_search()

	return

	fmt.Println("========= test case 1 ========")
	test_json_v1()

	fmt.Println()
	fmt.Println("========= test case 2 ========")
	test_json_v2()

	fmt.Println("========= yaml find ========")
	test_yaml_find()
	fmt.Println("")

	fmt.Println("========= yaml insert ========")
	test_yaml_insert()
	fmt.Println("")

}

func test_yaml_insert() {
	yamlMap, err := JsonMapper.NewYamlMap(yaml_adder_test)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("========= find: idx 2 / spec.mountOptions")
	// yamlMap.Find(2, "spec.mountOptions")

	// fmt.Println("========= find: idx 2 / spec.mountOptions.0")
	// yamlMap.Find(2, "spec.mountOptions.0")

	fmt.Println("========== yaml adder test")
	// fmt.Println(yamlMap.Find(0, "metadata.name"))
	// fmt.Println(yamlMap.Find(0, "data.config.yaml"))
	//fmt.Println(yamlMap.Find(0, "data.config.yaml.staticPasswords"))
	yamlMap.Insert(0, "data", "stirngs", tmp)
	fmt.Println(yamlMap.Prints())
}

var tmp string = `123qwe
asdqwe
zxcqwe
asdqwe
ccc
`

var yaml_adder_test string = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: dex
  namespace: auth
data:
  config.yaml: |
    logger:
      level: "debug"
      format: text
    oauth2:
      skipApprovalScreen: true
    enablePasswordDB: true
    staticPasswords:
    - email: user@example.com
      hash: $2y$12$4K/VkmDd1q1Orb3xAt82zu8gk7Ad6ReFR4LCP9UeYE90NLiN9Df72
      username: user
`

var test_yaml string = `
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-airflow-dags
spec:
  env:
  - name: EXTEND_VALUE
    value: true
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: sc-smb-airflow
  mountOptions:
    - dir_mode=0777
    - file_mode=0777
    - uid=50000
    - gid=0
    - noserverino  # required to prevent data corruption
  csi:
    driver: smb.csi.k8s.io
    readOnly: false
    volumeHandle: 10.175.41.91/airflow/dags##
    volumeAttributes:
      source: "//10.175.41.91/airflow/dags"
    nodeStageSecretRef:
      name: smbcreds
      namespace: default
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-airflow-dags
  namespace: airflow
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  volumeName: pv-airflow-dags
  storageClassName: sc-smb-airflow
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-airflow-logs
spec:
  capacity:
    storage: 20Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: sc-smb-airflow
  mountOptions:
    - dir_mode=0777
    - file_mode=0777
    - uid=50000
    - gid=0
    - noserverino  # required to prevent data corruption
  csi:
    driver: smb.csi.k8s.io
    readOnly: false
    volumeHandle: 10.175.41.91/airflow/logs##
    volumeAttributes:
      source: "//10.175.41.91/airflow/logs"
    nodeStageSecretRef:
      name: smbcreds
      namespace: default
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-airflow-logs
  namespace: airflow
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi
  volumeName: pv-airflow-logs
  storageClassName: sc-smb-airflow

`

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
		],
	"sub": {
		"id": "1000",
		"id2": "1111"
	}
}
`

func test_yaml_find() {
	yamlMap, err := JsonMapper.NewYamlMap(test_yaml)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(yamlMap.Print())

	// success
	// fmt.Println("========= find: idx 3 / spec.accessModes")
	// yamlMap.Find(3, "spec.accessModes")

	fmt.Println("========= find: idx 2 / spec.mountOptions")
	fmt.Println(yamlMap.Find(2, "spec.mountOptions").([]interface{}))

	fmt.Println("========= find: idx 2 / spec.mountOptions.0")
	fmt.Println(yamlMap.Find(2, "spec.mountOptions.0").(string))
}

func test_search() {
	j, e := JsonMapper.NewBytes([]byte(test_json_string))
	if nil != e {
		fmt.Println(e)
		return
	}

	// fmt.Println(j.PPrint())
	searchResult := j.Search("id", `{
		"gt" : "0000"
	}`)
	fmt.Println("### TEST SEARCH 1:", searchResult)

	searchResult = j.Search("id", `{
		"gt" : "2000"
	}`)
	fmt.Println("### TEST SEARCH 2:", searchResult)
}

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
