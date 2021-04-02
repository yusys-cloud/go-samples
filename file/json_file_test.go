// Author: yangzq80@gmail.com
// Date: 2020-09-07
//
package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestReadJsonFile(t *testing.T) {

	file := "config.json"

	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("read config file <%s> failure. err:%+v", file, err)
	}

	cnf := &Conf{}
	err = json.Unmarshal(data, cnf)
	if err != nil {
		log.Fatalf("parse config file <%s> failure. error:%+v", file, err)
	}

	log.Println(cnf.Units)
}

type Conf struct {
	Units []*ProxyUnit `json:"units,omitempty"`
}

// ProxyUnit proxyUnit
type ProxyUnit struct {
	Src            string `json:"src,omitempty"`
	Target         string `json:"target,omitempty"`
	Desc           string `json:"desc,omitempty"`
	TimeoutConnect int    `json:"timeoutConnect,omitempty"`
	TimeoutWrite   int    `json:"timeoutWrite,omitempty"`
	//Ctl            *Ctl   `json:"ctl,omitempty"`
}
