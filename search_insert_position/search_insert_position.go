// Search Insert Position
//
// [Easy] [AC:45.0% 999.8K of 2.2M] [filetype:golang]
//
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例 1:
//
// 输入: nums = [1,3,5,6], target = 5
//
// 输出: 2
//
// 示例 2:
//
// 输入: nums = [1,3,5,6], target = 2
//
// 输出: 1
//
// 示例 3:
//
// 输入: nums = [1,3,5,6], target = 7
//
// 输出: 4
//
// 提示:
//
// 1 <= nums.length <= 104
//
// -104 <= nums[i] <= 104
//
// nums 为 无重复元素 的 升序 排列数组
//
// -104 <= target <= 104
//
// [End of Description]
package main

import "fmt"

func main() {
	pos := searchInsert([]int{1, 3}, 2)
	fmt.Println(pos)

}

// local end
func searchInsert(nums []int, target int) int {

	listLen := len(nums)

	i := 0
	j := listLen - 1

	if listLen == 0 {
		return 0
	}

	if target < nums[0] {
		return 0
	}

	if target > nums[listLen-1] {
		return listLen
	}

	for i <= j {

		mid := (i + j) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[mid] > target {
			j = mid - 1
		}

		if nums[mid] < target {
			i = mid + 1
		}
	}

	if nums[j] > target {
		return j
	}

	return i
}
