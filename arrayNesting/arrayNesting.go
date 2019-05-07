package leetcode

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func arrayNesting(nums []int) int {
	/**
	 * We want to mark all indicies which we've already visited and once all have been marked,
	 * we're done.
	 *
	 * We do this in 2 loops, one checking we still have unvisited positions, and one building
	 * sets
	 */

	longest := 0
	visited := make(map[int]bool, 0)
	i := 0
	for i < len(nums) {
		_, ok := visited[i]
		size := 0
		tmp := i
		for !ok {
			visited[i] = true
			size++
			tmp = nums[tmp]
			_, ok = visited[tmp]
		}
		longest = max(longest, size)
		i++
	}
	return longest
}
