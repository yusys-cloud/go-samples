// Author: yangzq80@gmail.com
// Date: 2020-09-08
// net.TCPConn是允许在TCP客户端和TCP服务器之间的全双工通信的Go类型。两种主要方法是
// func (c *TCPConn) Write(b []byte) (n int, err os.Error)
// func (c *TCPConn) Read(b []byte) (n int, err os.Error)
// TCPConn被客户端和服务器用来读写消息。
//
package web

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {

	ids := []int{1, 1, 2}

	fmt.Println(majorityElement(ids))
	//time.Sleep(time.Second * 2)
}

func BenchmarkMajorityElement(b *testing.B) {
	ids := []int{1, 1, 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(ids)
	}
}

func BenchmarkMajorityElementMap(b *testing.B) {
	ids := []int{1, 1, 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElementMap(ids)
	}
}

func majorityElementMap(nums []int) int {
	m := map[int]int{}
	max_val := 0
	max_index := 0
	for _, e := range nums {
		if _, exist := m[e]; exist {
			m[e]++
		} else {
			m[e] = 1
		}
		if m[e] > max_val {
			max_val = m[e]
			max_index = e
		}
	}
	//fmt.Println(m)
	return max_index
}
