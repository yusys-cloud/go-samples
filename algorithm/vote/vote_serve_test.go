// Author: yangzq80@gmail.com
// Date: 2021-06-28
//
package vote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestStartServe(t *testing.T) {
	startServe()
}

// mac BenchmarkQuery-12    	   13461	     88323 ns/op
func BenchmarkQuery(b *testing.B) {

	ids := []int{1, 3, 3, 2, 3, 3, 4}

	if requestVote(ids, 121).Id == 3 {
		b.Error("Not equals 3")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		requestVote(ids, 139)
	}
}

func requestVote(ids []int, sn int) RespBody {

	reqBody := &ReqBody{ids, sn}

	rb, _ := json.Marshal(reqBody)

	resp, err := http.Post("http://localhost:10000/api/vote", "application/json", bytes.NewBuffer(rb))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var respBody RespBody

	json.Unmarshal(body, &respBody)

	return respBody
}
