import "fmt"

// Insert Interval
// [Medium] [AC:41.9% 131.6K of 314.4K] [filetype:golang]
//
// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
// 示例 1：
//
// 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//
// 输出：[[1,5],[6,9]]
//
// 示例 2：
//
// 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//
// 输出：[[1,2],[3,10],[12,16]]
//
// 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//
// 示例 3：
//
// 输入：intervals = [], newInterval = [5,7]
//
// 输出：[[5,7]]
//
// 示例 4：
//
// 输入：intervals = [[1,5]], newInterval = [2,3]
//
// 输出：[[1,5]]
//
// 示例 5：
//
// 输入：intervals = [[1,5]], newInterval = [2,7]
//
// 输出：[[1,7]]
//
// 提示：
//
// 0 <= intervals.length <= 104
//
// intervals[i].length == 2
//
// 0 <= intervals[i][0] <= intervals[i][1] <= 105
//
// intervals 根据 intervals[i][0] 按 升序 排列
//
// newInterval.length == 2
//
// 0 <= newInterval[0] <= newInterval[1] <= 105
//
// [End of Description]
// local end
func insert(intervals [][]int, newInterval []int) [][]int {

	start := newInterval[0]
	end := newInterval[1]

	ret := make([][]int, 0)
	if len(intervals) == 0 {
		ret = append(ret, newInterval)
		return ret
	}

	incStart := start
	incEnd := end

	// find incStart
	findIncStartFlag := false
	for _, v := range intervals {
		if start < v[0] {
			incStart = start
			findIncStartFlag = true
			break
		}

		if between(v[0], v[1], incStart) {
			incStart = v[0]
			findIncStartFlag = true
			break
		}
	}
	if findIncStartFlag == false {
		incStart = start
	}

	// find incEnd
	findIncEndFlag := false
	for _, v := range intervals {
		if end < v[0] {
			incEnd = end
			findIncEndFlag = true
			break
		}

		if between(v[0], v[1], incEnd) {
			incEnd = v[1]
			findIncEndFlag = true
			break
		}
	}
	if findIncEndFlag == false {
		incEnd = end
	}

	fmt.Printf("%v,%v", incStart, incEnd)

	appendFlag := false
	for _, v := range intervals {
		if appendFlag == false {
			if v[0] > incEnd {
				ret = append(ret, []int{incStart, incEnd})
				appendFlag = true
			}
		}

		if between(incStart, incEnd, v[0]) || between(incStart, incEnd, v[1]) {
			continue
		}
		ret = append(ret, v)

	}

	if appendFlag == false {
		ret = append(ret, []int{incStart, incEnd})
	}

	return ret
}

func between(start, end, num int) bool {

	if num >= start && num <= end {
		return true
	}
	return false
}
