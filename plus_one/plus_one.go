// Plus One
//
// [Easy] [AC:45.4% 573.6K of 1.3M] [filetype:golang]
//
// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1：
//
// 输入：digits = [1,2,3]
//
// 输出：[1,2,4]
//
// 解释：输入数组表示数字 123。
//
// 示例 2：
//
// 输入：digits = [4,3,2,1]
//
// 输出：[4,3,2,2]
//
// 解释：输入数组表示数字 4321。
//
// 示例 3：
//
// 输入：digits = [0]
//
// 输出：[1]
//
// 提示：
//
// 1 <= digits.length <= 100
//
// 0 <= digits[i] <= 9
//
// [End of Description]

package main

// local end
func plusOne(digits []int) []int {

	stack := make([]int, 0)
	for _, v := range digits {
		stack = push(stack, v)
	}

	carryNum := 0

	leaveS := make([]int, 0)
	flag := 0
	for s, v := pop(stack); v != -1; {

		sumNum := 0
		if flag == 0 {
			sumNum = v + 1 + carryNum
			flag = 1
		} else {
			sumNum = v + carryNum
		}

		posNum := sumNum % 10
		leaveS = push(leaveS, posNum)
		carryNum = sumNum / 10

		s, v = pop(s)
	}

	if carryNum > 0 {
		leaveS = push(leaveS, carryNum)
	}

	return leaveS
}

func push(stack []int, dig int) []int {
	newStack := []int{dig}
	newStack = append(newStack, stack...)
	return newStack
}

func pop(stack []int) ([]int, int) {
	if len(stack) == 0 {
		return stack, -1
	}
	if len(stack) == 1 {
		return []int{}, stack[0]
	}

	ret := stack[0]
	s := stack[1:]
	return s, ret
}
