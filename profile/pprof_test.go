// Author: yangzq80@gmail.com
// Date: 2020-09-01
// 1.	import pprof
// 2. 	go tool pprof -inuse_space http://127.0.0.1:10001/debug/pprof/heap
// 3.	可查看pdf图,火焰图等
//
package profile

import (
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func startWeb() {
	http.ListenAndServe("0.0.0.0:10001", nil)
}

func TestGet(t *testing.T) {
	startWeb()
}
