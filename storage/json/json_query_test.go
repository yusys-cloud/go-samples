// Author: yangzq80@gmail.com
// Date: 2021-04-12
//
package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"testing"
)

var (
	jsonStr2 = `[
    {
        "k": "code:1383028319401807872",
        "v": {
            "code": "iostat -d -m -c -x 2",
            "des": "-d 磁盘使用状态 -m 兆 -c 部分cpu 2秒 -x 设备使用率（%util）",
            "name": "磁盘io"
        }
    },
    {
        "k": "code:1383032203256008704",
        "v": {
            "code": "sysctl -a",
            "name": "Linux内核所有参数"
        }
    }
]`
	jsonStr = `
{
   "name":"computers",
   "description":"List of computer products",
   "vendor":{
      "name":"Star Trek",
      "email":"info@example.com",
      "website":"www.example.com",
      "items":[
         {
            "id":1,
            "name":"MacBook Pro 13 inch retina",
            "price":1350
         },
         {
            "id":2,
            "name":"MacBook Pro 15 inch retina",
            "price":1700
         },
         {
            "id":3,
            "name":"Sony VAIO",
            "price":1200
         },
         {
            "id":4,
            "name":"Fujitsu",
            "price":850
         },
         {
            "id":5,
            "name":"HP core i5",
            "price":850,
            "key": 2300
         },
         {
            "id":6,
            "name":"HP core i7",
            "price":950
         },
         {
            "id":null,
            "name":"HP core i3 SSD",
            "price":850
         }
      ],
      "prices":[
         2400,
         2100,
         1200,
         400.87,
         89.90,
         150.10
     ],
     "names":[
        "John Doe",
        "Jane Doe",
        "Tom",
        "Jerry",
        "Nicolas",
        "Abby"
     ]
   }
}
`
)

func TestFind(t *testing.T) {
	name := gojsonq.New().FromString(jsonStr).Find("name.first")
	println(name.(string)) // Tom

	jq := gojsonq.New().FromString(jsonStr).From("age")
	fmt.Printf("%#v\n", jq.Count())
}

func TestSearch(t *testing.T) {

	jq := gojsonq.New().FromString(jsonStr).From("vendor.items").Where("name", "contains", "i7")

	out := jq.Get()

	fmt.Println(out)
}

func TestSearchStr2(t *testing.T) {

	jq := gojsonq.New().FromString(jsonStr2).Where("v.name", "contains", "io")

	out := jq.Get()

	fmt.Println(out)
}

func TestOrder(t *testing.T) {
	jq := gojsonq.New().FromString(jsonStr).
		From("vendor.items").SortBy("price", "asc")

	out := jq.Only("id", "price")

	fmt.Println(out)
}

func TestJSONQ_Where_multiple_where_expecting_result(t *testing.T) {
	jq := gojsonq.New().FromString(jsonStr).
		From("vendor.items").
		Where("price", "=", 1700).
		Where("id", "=", 2)
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple Where expecting data")
}

func assertJSON(t *testing.T, v interface{}, expJSON string, tag ...string) {
	bb, err := json.Marshal(v)
	if err != nil {
		t.Errorf("failed to marshal: %v", err)
	}
	eb := []byte(expJSON)
	if !bytes.Equal(bb, eb) {
		if len(tag) > 0 {
			t.Errorf("Tag: %s\nExpected: %v\nGot: %v", tag[0], expJSON, string(bb))
		} else {
			t.Errorf("Expected: %v\nGot: %v", expJSON, string(bb))
		}
	}
}
