// Author: yangzq80@gmail.com
// Date: 2021-06-18
//
package algorithm

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {

	primes := []int{1, 1, 2, 2, 1}

	rs := majorityElement2(primes)
	fmt.Println(rs)
}

func majorityElement1(nums []int) int {
	vote := 0
	x := nums[0]
	for _, num := range nums {
		// 票数抵消
		if vote == 0 {
			// 假设当前数为众数
			x = num
		}
		if num == x {
			// 投正票
			vote++
		} else {
			// 投反票
			vote--
		}
	}
	return x
}

func majorityElement2(nums []int) int {
	vote := 0
	x := nums[0]
	for _, num := range nums {
		// 票数抵消
		if vote == 0 {
			// 假设当前数为众数
			x = num
		}
		if num == x {
			// 投正票
			vote++
		} else {
			// 投反票
			vote--
		}
	}
	return x
}
