// Author: yangzq80@gmail.com
// Date: 2021-03-30
//
package json

import (
	"encoding/json"
	"fmt"
	"github.com/schollz/jsonstore"
	"testing"
)

func TestSave(t *testing.T) {
	ks := new(jsonstore.JSONStore)

	// set a key to any object you want
	type Human struct {
		Name   string
		Height float64
	}
	err := ks.Set("human:1", Human{"Dante", 5.4})
	if err != nil {
		panic(err)
	}

	// Saving will automatically gzip if .gz is provided
	if err = jsonstore.Save(ks, "humans.json.gz"); err != nil {
		panic(err)
	}

	// Load any JSON / GZipped JSON
	ks2, err := jsonstore.Open("humans.json")
	if err != nil {
		panic(err)
	}

	// get the data back via an interface
	var human Human
	err = ks2.Get("human:1", &human)
	if err != nil {
		panic(err)
	}
	fmt.Println(human.Name) // Prints 'Dante'

	ks.Set("a", "{\"id\": \"redis-two-datacent-shared-cluster\",   \"deploy\": \"同城双活集群部署\",   \"des\": \"同城双数据中心共享同一集群使用\",   \"name\": \"redis\",   \"version\": \"5.0.9\",   \"type\": \"Cache Database\",   \"files\": [     \"redis-server\",     \"redis-cli\",     \"redis.conf\"   ],   \"rules\": {     \"limits\": [       {         \"nodes\": \"{{collection('rules.nodeRules').find({'role':'master'}).nodes}}\",         \"min\": 3       }     ],     \"nodeRules\": [       {         \"datacentId\": \"datacent-1\",         \"name\": \"主数据中心\",         \"nodes\": [           {             \"id\": \"node-1\",             \"role\": \"master\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-2\",             \"role\": \"master\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-3\",             \"role\": \"master\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-4\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-5\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-6\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           }         ]       },       {         \"datacentId\": \"datacent-2\",         \"name\": \"同城数据中心\",         \"link\": {           \"datacentId\": \"datacent-1\",           \"des\": \"Replicates data\"         },         \"nodes\": [           {             \"id\": \"node-7\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-8\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-9\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-10\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-11\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           },           {             \"id\": \"node-12\",             \"role\": \"slave\",             \"ip\": \"{{nil}}\"           }         ]       }     ]   },   \"cmds\": {     \"deployCmds\": [       {         \"id\": \"start-instances\",         \"cmd\": \"{{config.path}}/redis-server {{config.path}}/redis.conf\",         \"target\": \"allNodes\"       },       {         \"id\": \"config-master\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster create {{node-1.ip}}:6379 {{node-2.ip}}:6379 {{node-3.ip}}:6379 --cluster-replicas 0\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-1-slave1\",         \"cmd\": \"{{config.path}}/redis-cli --cluster add-node {{node-4.ip}}:6379 {{node-1.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-1-slave2\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-7.ip}}:6379 {{node-1.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-1-slave3\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-8.ip}}:6379 {{node-1.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-2-slave1\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-5.ip}}:6379 {{node-2.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-2-slave2\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-9.ip}}:6379 {{node-2.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-2-slave3\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-10.ip}}:6379 {{node-2.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-3-slave1\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-6.ip}}:6379 {{node-3.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-3-slave2\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-11.ip}}:6379 {{node-3.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"config-slave-master-3-slave3\",         \"cmd\": \"{{config.path}}/redis-cli -h {{node-1.ip}} --cluster add-node {{node-12.ip}}:6379 {{node-3.ip}}:6379 --cluster-slave\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       }     ],     \"managedCmds\": [       {         \"id\": \"stop\",         \"cmd\": \"{{config.path}}/redis-cli -c -h {{node-1.ip}} shutdown\",         \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       },       {         \"id\": \"start\",         \"cmd\": \"{{config.path}}/redis-server {{config.path}}/redis.conf\",         \"target\": \"allNodes\"       },       {         \"id\": \"remove\",         \"cmd\": \"rm -rf {{config.path}}/*\",         \"target\": \"allNodes\"       },       {           \"id\": \"health\",           \"cmd\": \"./cluster-cli redis\",           \"target\": \"{{collection('rules.nodeRules.nodes').findOne({'id':'node-1'})}}\"       }     ]   },   \"config\": {     \"auth\": true,     \"password\": \"admin\",     \"path\": \"./redis\"   } }")
	jsonstore.Save(ks, "a.json.gz")
	var r string
	ks.Get("a", &r)
	fmt.Println(r)
}

func TestJsonToMap(t *testing.T) {

	var jsonStr = `
{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World"
}
`
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonMap["array"])

}
