// Author: yangzq80@gmail.com
// Date: 2021-06-28
//
package vote

import (
	"encoding/json"
	"log"
	"net/http"
)

type ReqBody struct {
	Ids []int
	Sn  int
}
type RespBody struct {
	Sn int
	Id int
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var p ReqBody

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	resp := &RespBody{p.Sn, majorityElement(p.Ids)}

	json.NewEncoder(w).Encode(resp)
}

//摩尔投票法（Boyer–Moore majority vote algorithm）会比简单的Map等方式在时间O(n)、空间O(1)复杂度方面优势
func majorityElement(nums []int) int {
	count := 0
	restult := 0
	for _, num := range nums {
		if count == 0 {
			restult = num
		}
		if num == restult {
			count++
		} else {
			count--
		}
	}
	return restult
}

func startServe() {
	http.HandleFunc("/api/vote", handleRequest)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
