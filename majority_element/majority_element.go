// Majority Element
//
// [Easy] [AC:66.9% 622.7K of 931K] [filetype:golang]
//
// 给定一个大小为 n 的数组 nums
// ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
// 示例 1：
//
// 输入：nums = [3,2,3]
//
// 输出：3
//
// 示例 2：
//
// 输入：nums = [2,2,1,1,1,2,2]
//
// 输出：2
//
// 提示：
//
// n == nums.length
//
// 1 <= n <= 5 * 104
//
// -109 <= nums[i] <= 109
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// [End of Description]
// Package main provides ...

package main

// func majorityElement(nums []int) int {
// 	m := make(map[int]int, 0)

// 	for _, v := range nums {
// 		m[v] = m[v] + 1
// 	}

// 	maxNum := 0
//     maxRet :=0
// 	for k, v := range m {
// 		if v > maxNum {
// 			maxNum = v
//             maxRet = k
// 		}
// 	}

// 	return maxRet

// }

func majorityElement(nums []int) int {

	count := 0
	majorNum := 0
	for _, v := range nums {
		if count == 0 {
			majorNum = v
			count++
			continue
		}
		if v == majorNum {
			count++
		} else {
			count--
		}
	}

	return majorNum

}
