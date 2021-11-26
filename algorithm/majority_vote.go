// Author: yangzq80@gmail.com
// Date: 2021-06-18
//
package algorithm

func majorityElement(nums []int) []int {
	// 创建返回值
	var res = make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	// 初始化两个候选人 candidate，以及他们的计数票
	cand1 := nums[0]
	count1 := 0
	cand2 := nums[0]
	count2 := 0

	//摩尔投票法
	// 配对阶段
	for _, num := range nums {
		// 投票
		if cand1 == num {
			count1++
			continue
		}
		if cand2 == num {
			count2++
			continue
		}

		if count1 == 0 {
			cand1 = num
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = num
			count2++
			continue
		}

		count1--
		count2--
	}
	// 计数阶段
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if cand1 == num {
			count1++
		} else if cand2 == num {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}
