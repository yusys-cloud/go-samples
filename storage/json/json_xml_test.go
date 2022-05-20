// Author: yangzq80@gmail.com
// Date: 2022/5/20
//
// https://pkg.go.dev/github.com/clbanning/mxj#NewMapJson
package json

import (
	"fmt"
	"github.com/clbanning/mxj/v2"
	"testing"
)

func TestConvertJsonToXML(t *testing.T) {
	jsonStr := `{
        "fruits" : {
            "a": "apple",
            "b": "banana"
        },
        "colors" : {
            "r": "red",
            "g": "green"
        }
    }`

	//var obj map[string]interface{}
	////obj := make(map[string]interface{})
	////jsonByteValue, _ := json.Marshal(sourceJson)
	//json.Unmarshal([]byte(jsonStr), &obj)
	//fmt.Println(obj)

	mapVal, merr := mxj.NewMapJson([]byte(jsonStr))
	if merr != nil {
		// handle error
	}
	xmlVal, xerr := mapVal.Xml()
	if xerr != nil {
		// handle error
	}
	fmt.Println(string(xmlVal))
}
